import { loadConfig } from "./load-config.mjs";
import { runCmd } from "./run-cmd.mjs";

const config = await loadConfig();
const { paths, manifest } = config.ui8px.aria;

await runCmd(
  "ui8px",
  ["validate", "aria", ...paths, "--manifest", manifest, "--ignore", ".fastygo"],
  { label: "validate:aria" },
);
