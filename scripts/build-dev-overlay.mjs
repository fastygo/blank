import { loadConfig, resolveFromRoot } from "./load-config.mjs";
import { log } from "./log.mjs";
import { runCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const devOverlay = config.devOverlay;
if (!devOverlay) {
  throw new Error("fastygo.config.mjs: devOverlay section is required");
}

await runCmd(
  "bun",
  [
    "build",
    resolveFromRoot(devOverlay.entry),
    "--outfile",
    resolveFromRoot(devOverlay.output),
    "--format",
    "iife",
    "--minify",
  ],
  { label: "build:dev-overlay" },
);

log("dev overlay ready");
