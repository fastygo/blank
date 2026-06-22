import type { DevOverlayPanelI18n, StatusPayload } from "./types";
import { formatTemplate } from "./i18n";

const statusURL = "/__fastygo/dev/status.json";
const disableURL = "/__fastygo/dev/disable";

export async function fetchStatus(i18n: DevOverlayPanelI18n): Promise<StatusPayload> {
  const res = await fetch(statusURL, {
    headers: { Accept: "application/json" },
    cache: "no-store",
  });
  if (!res.ok) {
    throw new Error(formatTemplate(i18n.errors.status_json_failed, res.status));
  }
  return res.json() as Promise<StatusPayload>;
}

export async function disableOverlay(i18n: DevOverlayPanelI18n): Promise<void> {
  const res = await fetch(disableURL, {
    method: "POST",
    headers: { Accept: "application/json" },
    cache: "no-store",
  });
  if (!res.ok) {
    throw new Error(formatTemplate(i18n.errors.disable_failed, res.status));
  }
  window.location.reload();
}

export function readRequestId(): string {
  const script = document.querySelector<HTMLScriptElement>(
    'script[src="/__fastygo/dev/overlay.js"]',
  );
  return script?.dataset.requestId?.trim() ?? "";
}

export function readLocale(missingLabel: string): string {
  return document.documentElement.lang?.trim() || missingLabel;
}

export function readPath(): string {
  return `${window.location.pathname}${window.location.search}`;
}

export function formatBytes(size: number): string {
  if (size < 1024) return `${size} B`;
  if (size < 1024 * 1024) return `${(size / 1024).toFixed(1)} KB`;
  return `${(size / (1024 * 1024)).toFixed(1)} MB`;
}

export function formatAge(seconds: number): string {
  if (seconds < 60) return `${seconds}s`;
  if (seconds < 3600) return `${Math.floor(seconds / 60)}m`;
  return `${Math.floor(seconds / 3600)}h`;
}

export async function probe(
  path: string,
  fetchFailedLabel: string,
): Promise<{
  path: string;
  ok: boolean;
  status: number;
  latencyMs: number;
  error?: string;
}> {
  const start = performance.now();
  try {
    const res = await fetch(path, { cache: "no-store" });
    return {
      path,
      ok: res.ok,
      status: res.status,
      latencyMs: Math.round(performance.now() - start),
    };
  } catch (err) {
    return {
      path,
      ok: false,
      status: 0,
      latencyMs: Math.round(performance.now() - start),
      error: err instanceof Error ? err.message : fetchFailedLabel,
    };
  }
}
