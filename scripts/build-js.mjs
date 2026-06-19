import { loadConfig, resolveFromRoot, serverEnv } from "./load-config.mjs";
import { log } from "./log.mjs";
import { runCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const env = serverEnv(config);
const bundle = config.js.bundles[0];

if (!bundle) {
  throw new Error("fastygo.config.mjs: js.bundles must contain at least one entry");
}

await runCmd(
  "bun",
  [
    "build",
    resolveFromRoot(bundle.entry),
    "--outfile",
    resolveFromRoot(bundle.output),
    "--minify",
  ],
  { env, label: "build:js" },
);

log("js ready");
