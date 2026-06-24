import { loadConfig, serverEnv } from "./load-config.mjs";
import { log } from "./log.mjs";
import { runCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const env = serverEnv(config);

await runCmd(
  "go",
  ["tool", "templ", "generate", config.templ.generate],
  { env, label: "templ generate" },
);

log("templ ready");
