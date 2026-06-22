import type { DevOverlayPanelI18n } from "./types";

export function readPanelI18n(): DevOverlayPanelI18n | null {
  const script = document.querySelector<HTMLScriptElement>(
    'script[src="/__fastygo/dev/overlay.js"]',
  );
  const raw = script?.dataset.i18n?.trim();
  if (!raw) return null;
  try {
    return JSON.parse(raw) as DevOverlayPanelI18n;
  } catch {
    return null;
  }
}

export function formatTemplate(template: string, value: string | number): string {
  if (template.includes("%d")) {
    return template.replace("%d", String(value));
  }
  if (template.includes("%s")) {
    return template.replace("%s", String(value));
  }
  return template;
}
