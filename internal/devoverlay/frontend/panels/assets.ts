import { formatAge, formatBytes } from "../api";
import { statusBadge } from "../status";
import type { AssetStatus, DevContext, DevPanel } from "../types";

function assetBadgeClass(asset: AssetStatus): string {
  if (!asset.exists) {
    return statusBadge.error;
  }
  if (asset.stale) {
    return statusBadge.warning;
  }
  return statusBadge.success;
}

function assetBadgeLabel(asset: AssetStatus, copy: DevContext["i18n"]["assets"]): string {
  if (!asset.exists) {
    return copy.missing;
  }
  if (asset.stale) {
    return copy.stale;
  }
  return copy.present;
}

function renderAsset(root: HTMLElement, asset: AssetStatus, context: DevContext): void {
  const copy = context.i18n.assets;
  const row = document.createElement("div");
  row.className = "rounded-md border border-border p-3 flex flex-col gap-1";

  const head = document.createElement("div");
  head.className = "flex items-center justify-between gap-2";
  const name = document.createElement("code");
  name.className = "text-sm";
  name.textContent = asset.id;
  const state = document.createElement("span");
  state.className = assetBadgeClass(asset);
  state.textContent = assetBadgeLabel(asset, copy);
  head.append(name, state);

  const meta = document.createElement("p");
  meta.className = "text-sm text-muted-foreground";
  if (!asset.exists) {
    meta.textContent = asset.path;
  } else {
    meta.textContent = `${asset.path}${copy.separator}${formatBytes(asset.size)}${copy.separator}${copy.age_prefix}${formatAge(asset.ageSec)}`;
  }

  row.append(head, meta);

  if (asset.hint) {
    const hint = document.createElement("p");
    hint.className = "text-sm text-foreground";
    hint.textContent = asset.hint;
    row.append(hint);
  }

  root.append(row);
}

export const assetsPanel: DevPanel = {
  id: "assets",
  title: "Assets",
  mount(root, context) {
    const copy = context.i18n.assets;
    root.replaceChildren();

    const intro = document.createElement("p");
    intro.className = "text-sm text-muted-foreground";
    intro.textContent = copy.intro;
    root.append(intro);

    const list = document.createElement("div");
    list.className = "flex flex-col gap-2";
    root.append(list);

    let cancelled = false;

    async function run(): Promise<void> {
      list.replaceChildren();
      try {
        const payload = await context.fetchStatus();
        if (cancelled) return;
        for (const asset of payload.assets) {
          renderAsset(list, asset, context);
        }
        if (payload.hints?.length) {
          const hints = document.createElement("div");
          hints.className = "rounded-md border border-border bg-muted p-3 text-sm";
          hints.textContent = payload.hints.join(" ");
          list.append(hints);
        }
      } catch (err) {
        const error = document.createElement("p");
        error.className = "text-sm text-destructive";
        error.textContent = err instanceof Error ? err.message : copy.load_failed;
        list.append(error);
      }
    }

    void run();
    return () => {
      cancelled = true;
    };
  },
};
