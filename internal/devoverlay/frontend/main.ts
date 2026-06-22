import {
  disableOverlay,
  fetchStatus,
  readLocale,
  readPath,
  readRequestId,
} from "./api";
import { readPanelI18n } from "./i18n";
import { assetsPanel } from "./panels/assets";
import { healthPanel } from "./panels/health";
import { requestPanel } from "./panels/request";
import type { DevContext, DevOverlayPanelI18n, DevPanel, DevPanelID } from "./types";

const panels: DevPanel[] = [healthPanel, assetsPanel, requestPanel];

function panelById(id: DevPanelID): DevPanel | undefined {
  return panels.find((panel) => panel.id === id);
}

function setActiveTab(tabId: DevPanelID): void {
  for (const button of document.querySelectorAll<HTMLButtonElement>("[data-dev-tab]")) {
    const active = button.dataset.devTab === tabId;
    button.setAttribute("aria-selected", active ? "true" : "false");
    button.classList.toggle("bg-primary", active);
    button.classList.toggle("text-primary-foreground", active);
  }

  for (const panel of document.querySelectorAll<HTMLElement>("[data-dev-panel]")) {
    const active = panel.dataset.devPanel === tabId;
    panel.toggleAttribute("hidden", !active);
    panel.dataset.devPanelActive = active ? "true" : "false";
  }
}

function bindTabs(context: DevContext): () => void {
  const cleanups: Array<() => void> = [];

  for (const panel of panels) {
    const mount = document.getElementById(`fastygo-dev-panel-${panel.id}`);
    if (!mount) continue;
    cleanups.push(panel.mount(mount, context));
  }

  for (const button of document.querySelectorAll<HTMLButtonElement>("[data-dev-tab]")) {
    const handler = () => {
      const tabId = button.dataset.devTab as DevPanelID | undefined;
      if (!tabId || !panelById(tabId)) return;
      setActiveTab(tabId);
    };
    button.addEventListener("click", handler);
    cleanups.push(() => button.removeEventListener("click", handler));
  }

  return () => {
    for (const cleanup of cleanups) cleanup();
  };
}

function bindLauncher(): () => void {
  const launcher = document.getElementById("fastygo-dev-launcher");
  const panel = document.getElementById("fastygo-dev-panel");
  if (!launcher || !panel) {
    return () => {};
  }

  const toggle = () => {
    const open = panel.classList.toggle("hidden") === false;
    launcher.setAttribute("aria-expanded", open ? "true" : "false");
  };

  launcher.addEventListener("click", toggle);
  return () => launcher.removeEventListener("click", toggle);
}

function requirePanelI18n(): DevOverlayPanelI18n | null {
  return readPanelI18n();
}

export function mountDevOverlay(): void {
  const root = document.getElementById("fastygo-dev-overlay-root");
  const i18n = requirePanelI18n();
  if (!root || !i18n) return;

  const context: DevContext = {
    requestId: readRequestId(),
    path: readPath(),
    locale: readLocale(i18n.request.locale_missing),
    i18n,
    fetchStatus: () => fetchStatus(i18n),
    disableOverlay: () => disableOverlay(i18n),
  };

  const cleanupLauncher = bindLauncher();
  const cleanupTabs = bindTabs(context);

  const hide = document.getElementById("fastygo-dev-hide");
  const hideHandler = () => {
    void context.disableOverlay();
  };
  hide?.addEventListener("click", hideHandler);

  window.addEventListener(
    "beforeunload",
    () => {
      hide?.removeEventListener("click", hideHandler);
      cleanupLauncher();
      cleanupTabs();
    },
    { once: true },
  );
}
