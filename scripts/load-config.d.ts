/**
 * Tooling-only config shape for `fastygo.config.mjs`.
 * Do not add routes, layout presets, or runtime app shell selection here.
 */
export interface FastyGoConfig {
  /** Local dev server bind, static dir, and env vars for `go run ./cmd/server`. */
  server: {
    host: string;
    port: number;
    staticDir: string;
    env: Record<string, string>;
  };
  /** templ generate path; `watch` reserved for future watcher scripts. */
  templ: {
    generate: string;
    watch: string[];
  };
  /** Tailwind build input/output and documented @source scan roots. */
  css: {
    input: string;
    output: string;
    sources: string[];
  };
  /** Client JS bundles and ARIA manifest path. */
  js: {
    bundles: Array<{ entry: string; output: string }>;
    manifest: string;
  };
  /** ui8px lint paths and ARIA validation targets. */
  ui8px: {
    lint: string[];
    aria: {
      paths: string[];
      manifest: string;
    };
  };
  /** SSR dev overlay bundle (APP_DEV_OVERLAY). */
  devOverlay: {
    enabledEnv: string;
    entry: string;
    output: string;
  };
}
