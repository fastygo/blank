import path from "node:path";
import { loadConfig, resolveFromRoot, serverEnv, serverUrl } from "./load-config.mjs";
import { error, log } from "./log.mjs";
import { killChild, runCmd, spawnCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const env = serverEnv(config);

/** @type {import("node:child_process").ChildProcess[]} */
const children = [];
let shuttingDown = false;
/** @type {ReturnType<typeof setTimeout> | undefined} */
let templDebounce;

async function generateTempl() {
  await runCmd(
    "go",
    ["tool", "templ", "generate", config.templ.generate],
    { env, label: "templ generate" },
  );
  log("templ ready");
}

async function initialBuild() {
  await generateTempl();
  await runCmd("node", ["scripts/build-css.mjs"], { env, label: "build:css" });
  await runCmd("node", ["scripts/build-js.mjs"], { env, label: "build:js" });
}

function startCssWatch() {
  const input = resolveFromRoot(config.css.input);
  const output = resolveFromRoot(config.css.output);
  const child = spawnCmd(
    "tailwindcss",
    ["-i", input, "-o", output, "--watch"],
    { env },
  );
  children.push(child);
  log("css watching");
}

async function startTemplWatch() {
  const watchRoot = resolveFromRoot("internal");
  let subscribe;

  try {
    ({ subscribe } = await import("@parcel/watcher"));
  } catch (err) {
    error(`templ watch unavailable: ${err.message}`);
    error("run bun install and retry");
    return;
  }

  const queueTemplGenerate = () => {
    if (shuttingDown) {
      return;
    }
    clearTimeout(templDebounce);
    templDebounce = setTimeout(() => {
      generateTempl().catch((err) => {
        error(err.message);
      });
    }, 150);
  };

  await subscribe(
    watchRoot,
    (watchErr, events) => {
      if (watchErr) {
        error(`templ watch: ${watchErr.message}`);
        return;
      }
      const hasTemplChange = events.some(
        (event) => event.type !== "delete" && event.path.endsWith(".templ"),
      );
      if (hasTemplChange) {
        queueTemplGenerate();
      }
    },
    { ignore: ["**/*_templ.go", "**/node_modules/**"] },
  );

  log("templ watching");
}

function startServer() {
  const child = spawnCmd("go", ["run", "./cmd/server"], { env });
  children.push(child);
  log(`server ${serverUrl(config)}`);
  child.on("exit", (code, signal) => {
    if (shuttingDown) {
      return;
    }
    if (signal) {
      shutdown(1);
      return;
    }
    shutdown(code ?? 0);
  });
}

function shutdown(exitCode = 0) {
  if (shuttingDown) {
    return;
  }
  shuttingDown = true;
  clearTimeout(templDebounce);
  for (const child of children) {
    killChild(child);
  }
  process.exit(exitCode);
}

for (const sig of ["SIGINT", "SIGTERM", "SIGHUP"]) {
  process.on(sig, () => shutdown(0));
}

try {
  await initialBuild();
  startCssWatch();
  await startTemplWatch();
  startServer();
} catch (err) {
  error(err.message);
  shutdown(1);
}
