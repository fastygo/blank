/**
 * FastyGo Blank — dev/build tooling config only (Vite-like central config).
 *
 * This file does NOT define routes, page content, or route layout selection.
 * Runtime routing and layout adapters live in:
 *   - internal/site/router.go  (PageSpec.Layout, Title, Body, Nav)
 *   - internal/views/layout.templ  (AppShell, SidebarAppShell, …)
 *
 * @type {import('./scripts/load-config.d.ts').FastyGoConfig}
 */
export default {
  // Local server bind, static dir, and env vars passed to `go run ./cmd/server`.
  server: {
    host: "127.0.0.1",
    port: 8080,
    staticDir: "web/static",
    env: {
      APP_BIND: "127.0.0.1:8080",
      APP_STATIC_DIR: "web/static",
      APP_DEV_OVERLAY: "1",
    },
  },
  // templ generate path; `watch` is reserved for future watcher scripts (not used by `bun run dev` today).
  templ: {
    generate: "./...",
    watch: ["internal/**/*.templ"],
  },
  // Tailwind input/output; `sources` documents @source scan roots for policy/docs.
  css: {
    input: "web/static/css/input.css",
    output: "web/static/css/app.css",
    sources: [
      "internal/views/**/*.templ",
      "internal/ui/**/*.templ",
      "vendor/github.com/fastygo/templ/ui/**/*.templ",
      "vendor/github.com/fastygo/templ/components/**/*.templ",
    ],
  },
  // Client bundles (@ui8kit/aria subset) and manifest for validate:aria.
  js: {
    bundles: [
      {
        entry: "scripts/ui8kit-entry.mjs",
        output: "web/static/js/ui8kit.js",
      },
    ],
    manifest: "web/static/js/manifest.json",
  },
  // SSR dev overlay bundle (loopback only; gated by APP_DEV_OVERLAY).
  devOverlay: {
    enabledEnv: "APP_DEV_OVERLAY",
    entry: "scripts/dev-overlay-entry.ts",
    output: "internal/devoverlay/static/overlay.js",
  },
  // ui8px lint and ARIA validation paths (not runtime UI config).
  ui8px: {
    lint: [
      "internal/views",
      "internal/ui",
      "web/static/css/input.css",
      "web/static/css/latty-icons.css",
    ],
    aria: {
      paths: ["internal/views", "internal/ui"],
      manifest: "web/static/js/manifest.json",
    },
  },
};
