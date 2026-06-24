import type { DevOverlayPanelI18n } from "./types";

function readOverlayPayload(): string {
  const globalPayload = (
    globalThis as {
      __FASTYGO_DEV_OVERLAY__?: { i18n?: string };
    }
  ).__FASTYGO_DEV_OVERLAY__;
  if (globalPayload?.i18n) return globalPayload.i18n.trim();

  const script = document.querySelector<HTMLScriptElement>(
    'script[data-fastygo-dev-overlay-bundle]',
  );
  return script?.dataset.i18n?.trim() ?? "";
}

export function readPanelI18n(): DevOverlayPanelI18n | null {
  const raw = readOverlayPayload();
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
