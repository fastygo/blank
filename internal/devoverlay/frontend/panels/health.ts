import { probe } from "../api";
import { statusBadge } from "../status";
import type { DevContext, DevPanel, ProbeResult } from "../types";

function badgeClass(ok: boolean, status: number): string {
  if (ok) {
    return statusBadge.success;
  }
  if (status === 0) {
    return statusBadge.error;
  }
  return statusBadge.info;
}

function renderRow(root: HTMLElement, result: ProbeResult): void {
  const row = document.createElement("div");
  row.className = "rounded-md border border-border p-3 flex flex-col gap-1";

  const title = document.createElement("div");
  title.className = "flex items-center justify-between gap-2";
  const label = document.createElement("code");
  label.className = "text-sm";
  label.textContent = result.path;
  const badge = document.createElement("span");
  badge.className = badgeClass(result.ok, result.status);
  badge.textContent = result.ok ? `OK ${result.status}` : result.status ? `HTTP ${result.status}` : "Down";
  title.append(label, badge);

  const meta = document.createElement("p");
  meta.className = "text-sm text-muted-foreground";
  meta.textContent = `${result.latencyMs} ms`;
  if (result.error) {
    meta.textContent += ` · ${result.error}`;
  }

  row.append(title, meta);
  root.append(row);
}

export const healthPanel: DevPanel = {
  id: "health",
  title: "Health",
  mount(root, _context) {
    root.replaceChildren();

    const intro = document.createElement("p");
    intro.className = "text-sm text-muted-foreground";
    intro.textContent = "Server is up when both probes respond with HTTP 2xx.";
    root.append(intro);

    const refresh = document.createElement("button");
    refresh.type = "button";
    refresh.className =
      "inline-flex h-8 items-center justify-center rounded-md border border-border bg-background px-3 text-sm hover:bg-accent";
    refresh.textContent = "Refresh probes";
    root.append(refresh);

    const list = document.createElement("div");
    list.className = "flex flex-col gap-2";
    root.append(list);

    let cancelled = false;

    async function run(): Promise<void> {
      list.replaceChildren();
      const paths = ["/healthz", "/readyz"];
      const results = await Promise.all(paths.map((path) => probe(path)));
      if (cancelled) return;
      for (const result of results) {
        renderRow(list, result);
      }
    }

    refresh.addEventListener("click", () => {
      void run();
    });
    void run();

    return () => {
      cancelled = true;
    };
  },
};
