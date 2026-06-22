import type { StatusPayload } from "./types";

const statusURL = "/__fastygo/dev/status.json";
const disableURL = "/__fastygo/dev/disable";

export async function fetchStatus(): Promise<StatusPayload> {
  const res = await fetch(statusURL, {
    headers: { Accept: "application/json" },
    cache: "no-store",
  });
  if (!res.ok) {
    throw new Error(`status.json ${res.status}`);
  }
  return res.json() as Promise<StatusPayload>;
}

export async function disableOverlay(): Promise<void> {
  const res = await fetch(disableURL, {
    method: "POST",
    headers: { Accept: "application/json" },
    cache: "no-store",
  });
  if (!res.ok) {
    throw new Error(`disable ${res.status}`);
  }
  window.location.reload();
}

export function readRequestId(): string {
  const script = document.querySelector<HTMLScriptElement>(
    'script[src="/__fastygo/dev/overlay.js"]',
  );
  return script?.dataset.requestId?.trim() ?? "";
}

export function readLocale(): string {
  return document.documentElement.lang?.trim() || "(missing)";
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

export async function probe(path: string): Promise<{
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
      error: err instanceof Error ? err.message : "fetch failed",
    };
  }
}
