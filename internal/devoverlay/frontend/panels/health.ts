import { probe } from "../api";
import { formatTemplate } from "../i18n";
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

function badgeText(result: ProbeResult, copy: DevContext["i18n"]["health"]): string {
  if (result.ok) {
    return formatTemplate(copy.ok_badge, result.status);
  }
  if (result.status) {
    return formatTemplate(copy.http_badge, result.status);
  }
  return copy.down_label;
}

function renderRow(root: HTMLElement, result: ProbeResult, context: DevContext): void {
  const copy = context.i18n.health;
  const row = document.createElement("div");
  row.className = "rounded-md border border-border p-3 flex flex-col gap-1";

  const title = document.createElement("div");
  title.className = "flex items-center justify-between gap-2";
  const label = document.createElement("code");
  label.className = "text-sm";
  label.textContent = result.path;
  const badge = document.createElement("span");
  badge.className = badgeClass(result.ok, result.status);
  badge.textContent = badgeText(result, copy);
  title.append(label, badge);

  const meta = document.createElement("p");
  meta.className = "text-sm text-muted-foreground";
  meta.textContent = `${result.latencyMs} ${copy.latency_unit}`;
  if (result.error) {
    meta.textContent += `${copy.separator}${result.error}`;
  }

  row.append(title, meta);
  root.append(row);
}

export const healthPanel: DevPanel = {
  id: "health",
  title: "Health",
  mount(root, context) {
    const copy = context.i18n.health;
    root.replaceChildren();

    const intro = document.createElement("p");
    intro.className = "text-sm text-muted-foreground";
    intro.textContent = copy.intro;
    root.append(intro);

    const refresh = document.createElement("button");
    refresh.type = "button";
    refresh.className =
      "inline-flex h-8 items-center justify-center rounded-md border border-border bg-background px-3 text-sm hover:bg-accent";
    refresh.textContent = copy.refresh_button;
    root.append(refresh);

    const list = document.createElement("div");
    list.className = "flex flex-col gap-2";
    root.append(list);

    let cancelled = false;

    async function run(): Promise<void> {
      list.replaceChildren();
      const paths = ["/healthz", "/readyz"];
      const results = await Promise.all(
        paths.map((path) => probe(path, context.i18n.errors.fetch_failed)),
      );
      if (cancelled) return;
      for (const result of results) {
        renderRow(list, result, context);
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
