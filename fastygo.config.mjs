/** @type {import('./scripts/load-config.d.ts').FastyGoConfig} */
export default {
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
  templ: {
    generate: "./...",
    watch: ["internal/**/*.templ"],
  },
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
  js: {
    bundles: [
      {
        entry: "scripts/ui8kit-entry.mjs",
        output: "web/static/js/ui8kit.js",
      },
    ],
    manifest: "web/static/js/manifest.json",
  },
  devOverlay: {
    enabledEnv: "APP_DEV_OVERLAY",
    entry: "scripts/dev-overlay-entry.ts",
    output: "internal/devoverlay/static/overlay.js",
  },
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
