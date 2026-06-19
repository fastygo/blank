import { loadConfig, resolveFromRoot, serverEnv } from "./load-config.mjs";
import { log } from "./log.mjs";
import { spawnCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const env = serverEnv(config);
const input = resolveFromRoot(config.css.input);
const output = resolveFromRoot(config.css.output);

const child = spawnCmd(
  "tailwindcss",
  ["-i", input, "-o", output, "--watch"],
  { env },
);

log("css watching");

child.on("exit", (code, signal) => {
  if (signal) {
    process.exit(1);
  }
  process.exit(code ?? 0);
});

for (const sig of ["SIGINT", "SIGTERM", "SIGHUP"]) {
  process.on(sig, () => process.exit(0));
}
