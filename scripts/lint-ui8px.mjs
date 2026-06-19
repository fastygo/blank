import { loadConfig } from "./load-config.mjs";
import { runCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const lintPaths = config.ui8px.lint;

await runCmd("ui8px", ["lint", ...lintPaths, "--ignore", ".fastygo"], {
  label: "lint:ui8px",
});
