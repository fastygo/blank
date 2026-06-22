import {
  disableOverlay,
  fetchStatus,
  readLocale,
  readPath,
  readRequestId,
} from "./api";
import { assetsPanel } from "./panels/assets";
import { healthPanel } from "./panels/health";
import { requestPanel } from "./panels/request";
import type { DevContext, DevPanel, DevPanelID } from "./types";

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

export function mountDevOverlay(): void {
  const root = document.getElementById("fastygo-dev-overlay-root");
  if (!root) return;

  const context: DevContext = {
    requestId: readRequestId(),
    path: readPath(),
    locale: readLocale(),
    fetchStatus,
    disableOverlay,
  };

  const cleanupLauncher = bindLauncher();
  const cleanupTabs = bindTabs(context);

  const hide = document.getElementById("fastygo-dev-hide");
  const hideHandler = () => {
    void disableOverlay();
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
