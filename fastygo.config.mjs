/**
 * Blank tooling config (Vite-like single file for paths & policy).
 *
 * Not a bundler / HMR runtime. Go owns the server (`cmd/server`, Feature modules).
 * This file is read by Bun scripts (CSS, retag, JS) so paths and the retag allowlist
 * stay in one place — same idea as vite.config, without inventing a JS app shell.
 *
 * Runtime routes, DI, and page composition stay in Go:
 *   internal/site · internal/views · internal/ui/{layout,blocks,components}
 */
export default {
  server: {
    host: "0.0.0.0",
    port: 3000,
    staticDir: "web/static",
    /** Defaults mirrored by scripts/dev.sh; env / .env still win at runtime. */
    env: {
      APP_BIND: "0.0.0.0:3000",
      APP_STATIC_DIR: "web/static",
      APP_DEFAULT_LOCALE: "en",
      APP_AVAILABLE_LOCALES: "en,ru",
      APP_DEV_OVERLAY: "1",
    },
  },

  templ: {
    generate: "./...",
  },

  css: {
    input: "web/static/css/input.css",
    output: "web/static/css/app.css",
    /** Documented scan roots — must stay in sync with @source in input.css. */
    sources: [
      "internal/views/**/*.templ",
      "internal/ui/**/*.templ",
      "internal/kit/ui/**/*.templ",
      "internal/kit/ui/**/*.variants.json",
      "vendor/github.com/fastygo/templ/components/**/*.templ",
    ],
  },

  /**
   * Client bundles (@ui8kit/aria subset for Sheet).
   * Kept while mobile sheet uses Behavior: "ui8kit".
   */
  js: {
    bundles: [
      {
        entry: "scripts/ui8kit-entry.mjs",
        output: "web/static/js/ui8kit.js",
      },
    ],
    manifest: "web/static/js/manifest.json",
  },

  /** SSR dev overlay bundle (loopback only; gated by APP_DEV_OVERLAY). */
  devOverlay: {
    enabledEnv: "APP_DEV_OVERLAY",
    entry: "scripts/dev-overlay-entry.ts",
    output: "internal/devoverlay/static/overlay.js",
  },

  /**
   * UI registry map — where to find layers (Cursor rule: ui-registry.mdc).
   * Shapes future shadcn-like `fastygo add` / FastyGoUI extraction.
   */
  registry: {
    primitives: "internal/kit/ui",
    utils: "internal/kit/utils",
    layout: "internal/ui/layout",
    components: "internal/ui/components",
    blocks: "internal/ui/blocks",
    widgets: "internal/ui/widgets",
    variants: "internal/ui/variants",
    views: "internal/views",
    locale: "internal/fixtures/locale",
    /** Planned extract / copy targets (not Go modules in this repo yet). */
    future: {
      primitives: "github.com/fastygo/templ",
      blocks: "github.com/fastygo/blocks",
      widgets: "github.com/fastygo/widgets",
      components: "cli-copy",
      layout: "app-owned",
    },
  },

  /**
   * Retag policy — source of truth for scripts/retag-check.sh
   * and .project/retag.md.
   */
  retag: {
    /** Skip these trees (codegen atoms). */
    skip: ["internal/kit/ui"],
    /** Document shell: only layoutTags allowed as raw HTML. */
    layout: ["internal/ui/layout"],
    layoutTags: [
      "html",
      "head",
      "body",
      "meta",
      "title",
      "link",
      "script",
      "noscript",
      "main",
    ],
    /**
     * Explicit registry: raw <input type="hidden"> only in listed files.
     * Adding a new hidden field requires an entry here (comment alone is not enough).
     */
    allow: [],
    baseline: ".project/retag-baseline.txt",
  },
};
