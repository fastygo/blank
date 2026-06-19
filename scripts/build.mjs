import { loadConfig } from "./load-config.mjs";
import { log } from "./log.mjs";
import { runCmd } from "./run-cmd.mjs";

const config = await loadConfig();

await runCmd("node", ["scripts/templ-generate.mjs"], { label: "build: templ" });
await runCmd("node", ["scripts/build-css.mjs"], { label: "build: css" });
await runCmd("node", ["scripts/build-js.mjs"], { label: "build: js" });
await runCmd("go", ["build", "-o", "blank", "./cmd/server"], { label: "build: go" });

log("build finished → ./blank");
