import { loadConfig, serverEnv, serverUrl } from "./load-config.mjs";
import { log } from "./log.mjs";
import { runCmd, spawnCmd, killChild } from "./run-cmd.mjs";

const config = await loadConfig();
const env = serverEnv(config);

await runCmd("node", ["scripts/templ-generate.mjs"], { env });
await runCmd("node", ["scripts/build-css.mjs"], { env });
await runCmd("node", ["scripts/build-js.mjs"], { env });
if (env.APP_DEV_OVERLAY === "1") {
  await runCmd("node", ["scripts/build-dev-overlay.mjs"], { env, label: "build:dev-overlay" });
}

log(`starting server ${serverUrl(config)}`);

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
