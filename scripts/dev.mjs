import { loadConfig, serverEnv, serverUrl } from "./load-config.mjs";
import { log } from "./log.mjs";
import { runCmd, spawnCmd, killChild } from "./run-cmd.mjs";

const config = await loadConfig();
const env = serverEnv(config);

await runCmd("node", ["scripts/templ-generate.mjs"], { env, label: "templ generate" });
await runCmd("node", ["scripts/build-css.mjs"], { env, label: "build:css" });
await runCmd("node", ["scripts/build-js.mjs"], { env, label: "build:js" });
if (env.APP_DEV_OVERLAY === "1") {
  await runCmd("node", ["scripts/build-dev-overlay.mjs"], { env, label: "build:dev-overlay" });
  log("dev overlay enabled (APP_DEV_OVERLAY=1)");
} else {
  log("dev overlay disabled (set APP_DEV_OVERLAY=1 to enable)");
}

log(`starting server ${serverUrl(config)}`);
log("after .templ edits run bun run templ; after Go edits restart dev (Ctrl+C, then bun run dev)");

const server = spawnCmd("go", ["run", "./cmd/server"], { env });

function shutdown() {
  killChild(server);
  process.exit(0);
}

for (const sig of ["SIGINT", "SIGTERM", "SIGHUP"]) {
  process.on(sig, shutdown);
}

server.on("exit", (code, signal) => {
  if (signal) {
    process.exit(1);
  }
  process.exit(code ?? 0);
});
