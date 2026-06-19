import { loadConfig, resolveFromRoot, serverEnv } from "./load-config.mjs";
import { log } from "./log.mjs";
import { runCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const env = serverEnv(config);
const input = resolveFromRoot(config.css.input);
const output = resolveFromRoot(config.css.output);

await runCmd(
  "tailwindcss",
  ["-i", input, "-o", output, "--minify"],
  { env, label: "build:css" },
);

log("css ready");
