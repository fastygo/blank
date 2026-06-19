import { loadConfig } from "./load-config.mjs";
import { runCmd } from "./run-cmd.mjs";

await runCmd("node", ["scripts/templ-generate.mjs"], { label: "verify: templ" });
await runCmd("node", ["scripts/build-css.mjs"], { label: "verify: css" });
await runCmd("node", ["scripts/build-js.mjs"], { label: "verify: js" });
await runCmd("node", ["scripts/lint-ui8px.mjs"], { label: "verify: ui8px" });
await runCmd("node", ["scripts/validate-aria.mjs"], { label: "verify: aria" });
await runCmd("go", ["test", "./..."], { label: "verify: go test" });
