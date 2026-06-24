export const shadowStyles = `
:host {
  position: fixed;
  right: 1rem;
  bottom: 1rem;
  z-index: 30;
  pointer-events: none;
  color-scheme: light dark;
  font-family:
    ui-sans-serif,
    system-ui,
    -apple-system,
    BlinkMacSystemFont,
    "Segoe UI",
    sans-serif;
  font-size: 14px;
  line-height: 1.5;
}

* {
  box-sizing: border-box;
}

button {
  font: inherit;
  cursor: pointer;
}

[hidden],
.hidden {
  display: none !important;
}

.launcher {
  pointer-events: auto;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 2.5rem;
  height: 2.5rem;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 9999px;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  box-shadow: 0 10px 30px rgb(15 23 42 / 0.16);
  transition:
    background-color 120ms ease,
    color 120ms ease,
    box-shadow 120ms ease;
}

.launcher:hover {
  background: hsl(210 40% 96%);
}

.launcher:focus-visible,
.button:focus-visible,
.tab:focus-visible {
  outline: 2px solid hsl(217 91% 60%);
  outline-offset: 2px;
}

.panel {
  pointer-events: auto;
  position: absolute;
  right: 0;
  bottom: calc(100% + 0.5rem);
  width: min(92vw, 420px);
  max-height: min(80vh, 640px);
  overflow: auto;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 0.75rem;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  padding: 1rem;
  box-shadow:
    0 20px 25px -5px rgb(15 23 42 / 0.16),
    0 8px 10px -6px rgb(15 23 42 / 0.12);
}

.mobile-reload {
  pointer-events: auto;
  position: absolute;
  right: 0;
  bottom: calc(100% + 0.5rem);
  display: flex;
  width: min(88vw, 18rem);
  flex-direction: column;
  gap: 0.75rem;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 0.75rem;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  padding: 0.875rem;
  box-shadow:
    0 20px 25px -5px rgb(15 23 42 / 0.16),
    0 8px 10px -6px rgb(15 23 42 / 0.12);
}

.stack,
.flex-col {
  display: flex;
  flex-direction: column;
}

.gap-1 {
  gap: 0.25rem;
}

.gap-2 {
  gap: 0.5rem;
}

.gap-4 {
  gap: 1rem;
}

.w-full {
  width: 100%;
}

.flex,
.inline-flex {
  display: flex;
}

.inline-flex {
  display: inline-flex;
}

.items-center {
  align-items: center;
}

.justify-between {
  justify-content: space-between;
}

.flex-wrap {
  flex-wrap: wrap;
}

.rounded-md {
  border-radius: 0.375rem;
}

.border {
  border-width: 1px;
  border-style: solid;
}

.border-border {
  border-color: hsl(214 32% 91%);
}

.border-transparent {
  border-color: transparent;
}

.bg-background {
  background: hsl(0 0% 100%);
}

.bg-muted {
  background: hsl(210 40% 96%);
}

.bg-primary {
  background: hsl(222 47% 11%);
}

.bg-emerald-500 {
  background: #10b981;
}

.bg-pink-600 {
  background: #db2777;
}

.bg-amber-500 {
  background: #f59e0b;
}

.bg-sky-600 {
  background: #0284c7;
}

.text-primary-foreground,
.text-white {
  color: #fff;
}

.text-foreground {
  color: hsl(222 47% 11%);
}

.text-muted-foreground {
  color: hsl(215 16% 47%);
}

.text-destructive {
  color: #db2777;
}

.text-sm {
  font-size: 0.875rem;
  line-height: 1.25rem;
}

.text-xs {
  font-size: 0.75rem;
  line-height: 1rem;
}

.font-medium {
  font-weight: 500;
}

.h-8 {
  height: 2rem;
}

.w-fit {
  width: fit-content;
}

.px-2 {
  padding-left: 0.5rem;
  padding-right: 0.5rem;
}

.px-3 {
  padding-left: 0.75rem;
  padding-right: 0.75rem;
}

.py-0\\.5 {
  padding-top: 0.125rem;
  padding-bottom: 0.125rem;
}

.p-3 {
  padding: 0.75rem;
}

.break-all {
  word-break: break-all;
}

.block {
  display: block;
}

.button,
.tab {
  display: inline-flex;
  min-height: 2rem;
  align-items: center;
  justify-content: center;
  border: 1px solid hsl(214 32% 91%);
  border-radius: 0.375rem;
  background: hsl(0 0% 100%);
  color: hsl(222 47% 11%);
  padding: 0.375rem 0.75rem;
}

.button:hover,
.tab:hover {
  background: hsl(210 40% 96%);
}

.separator {
  width: 100%;
  height: 1px;
  border: 0;
  background: hsl(214 32% 91%);
}

h2,
p {
  margin: 0;
}

code {
  font-family:
    ui-monospace,
    SFMono-Regular,
    Menlo,
    Monaco,
    Consolas,
    "Liberation Mono",
    monospace;
}

@media (prefers-color-scheme: dark) {
  .launcher,
  .mobile-reload,
  .panel,
  .bg-background,
  .button,
  .tab {
    border-color: hsl(217 33% 18%);
    background: hsl(222 47% 11%);
    color: hsl(210 40% 98%);
  }

  .launcher:hover,
  .button:hover,
  .tab:hover,
  .bg-muted {
    background: hsl(217 33% 18%);
  }

  .text-foreground {
    color: hsl(210 40% 98%);
  }

  .text-muted-foreground {
    color: hsl(215 20% 65%);
  }

  .border-border,
  .separator {
    border-color: hsl(217 33% 18%);
  }

  .separator {
    background: hsl(217 33% 18%);
  }
}
`;
