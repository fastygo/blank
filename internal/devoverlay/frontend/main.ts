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
import { shadowStyles } from "./styles";
import type { DevContext, DevOverlayPanelI18n, DevPanel, DevPanelID } from "./types";

const elementName = "fastygo-dev-overlay";
const panels: DevPanel[] = [healthPanel, assetsPanel, requestPanel];

function panelById(id: DevPanelID): DevPanel | undefined {
  return panels.find((panel) => panel.id === id);
}

function setActiveTab(root: ParentNode, tabId: DevPanelID): void {
  for (const button of root.querySelectorAll<HTMLButtonElement>("[data-dev-tab]")) {
    const active = button.dataset.devTab === tabId;
    button.setAttribute("aria-selected", active ? "true" : "false");
    button.classList.toggle("bg-primary", active);
    button.classList.toggle("text-primary-foreground", active);
  }

  for (const panel of root.querySelectorAll<HTMLElement>("[data-dev-panel]")) {
    const active = panel.dataset.devPanel === tabId;
    panel.toggleAttribute("hidden", !active);
    panel.dataset.devPanelActive = active ? "true" : "false";
  }
}

function bindTabs(root: ParentNode, context: DevContext): () => void {
  const cleanups: Array<() => void> = [];

  for (const panel of panels) {
    const mount = root.querySelector<HTMLElement>(`#fastygo-dev-panel-${panel.id}`);
    if (!mount) continue;
    cleanups.push(panel.mount(mount, context));
  }

  for (const button of root.querySelectorAll<HTMLButtonElement>("[data-dev-tab]")) {
    const handler = () => {
      const tabId = button.dataset.devTab as DevPanelID | undefined;
      if (!tabId || !panelById(tabId)) return;
      setActiveTab(root, tabId);
    };
    button.addEventListener("click", handler);
    cleanups.push(() => button.removeEventListener("click", handler));
  }

  return () => {
    for (const cleanup of cleanups) cleanup();
  };
}

function bindLauncher(root: ParentNode): () => void {
  const launcher = root.querySelector<HTMLButtonElement>("#fastygo-dev-launcher");
  const panel = root.querySelector<HTMLElement>("#fastygo-dev-panel");
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

function isMobileViewport(): boolean {
  return window.matchMedia("(max-width: 767px)").matches;
}

function escapeHTML(value: string): string {
  return value
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#39;");
}

function loadingTabText(template: string, label: string): string {
  if (template.includes("%s")) {
    return template.replace("%s", label);
  }
  return `${template} ${label}`;
}

function tabButton(id: DevPanelID, label: string, selected: boolean): string {
  const activeClasses = selected ? " bg-primary text-primary-foreground" : "";
  return `<button
    id="fastygo-dev-tab-${id}"
    class="tab${activeClasses}"
    type="button"
    role="tab"
    aria-controls="fastygo-dev-panel-${id}"
    aria-selected="${selected ? "true" : "false"}"
    data-dev-tab="${id}"
    data-dev-tab-label="${escapeHTML(label)}"
  >${escapeHTML(label)}</button>`;
}

function tabPanel(id: DevPanelID, label: string, loadingTemplate: string, selected: boolean): string {
  return `<section
    class="gap-2"
    id="fastygo-dev-panel-wrap-${id}"
    role="tabpanel"
    aria-labelledby="fastygo-dev-tab-${id}"
    data-dev-panel="${id}"
    data-dev-panel-label="${escapeHTML(label)}"
    data-dev-panel-active="${selected ? "true" : "false"}"
    ${selected ? "" : "hidden"}
  >
    <p class="text-sm text-muted-foreground">${escapeHTML(loadingTabText(loadingTemplate, label))}</p>
    <div id="fastygo-dev-panel-${id}" class="flex flex-col gap-2"></div>
  </section>`;
}

function cogIcon(): string {
  return `<svg
    width="20"
    height="20"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    aria-hidden="true"
  >
    <path d="M11 10.27 7 3.34"></path>
    <path d="m11 13.73-4 6.93"></path>
    <path d="M12 22v-2"></path>
    <path d="M12 2v2"></path>
    <path d="M14 12h8"></path>
    <path d="m17 20.66-1-1.73"></path>
    <path d="m17 3.34-1 1.73"></path>
    <path d="M2 12h2"></path>
    <path d="m20.66 17-1.73-1"></path>
    <path d="m20.66 7-1.73 1"></path>
    <path d="m3.34 17 1.73-1"></path>
    <path d="m3.34 7 1.73 1"></path>
    <circle cx="12" cy="12" r="2"></circle>
    <circle cx="12" cy="12" r="8"></circle>
  </svg>`;
}

function overlayMarkup(i18n: DevOverlayPanelI18n): string {
  return `
    <style>${shadowStyles}</style>
    <aside
      id="fastygo-dev-mobile-reload"
      class="mobile-reload hidden"
      role="status"
      aria-live="polite"
    >
      <p>${escapeHTML(i18n.mobile_reload_hint)}</p>
      <button
        id="fastygo-dev-mobile-reload-button"
        class="button w-fit"
        type="button"
      >${escapeHTML(i18n.mobile_reload_button)}</button>
    </aside>
    <button
      id="fastygo-dev-launcher"
      class="launcher"
      type="button"
      aria-controls="fastygo-dev-panel"
      aria-expanded="false"
      aria-label="${escapeHTML(i18n.launcher_aria_label)}"
    >${cogIcon()}</button>
    <section
      id="fastygo-dev-panel"
      class="panel hidden"
      role="region"
      aria-label="${escapeHTML(i18n.panel_aria_label)}"
    >
      <div class="stack gap-4 w-full">
        <h2>${escapeHTML(i18n.title)}</h2>
        <p class="text-sm text-muted-foreground">${escapeHTML(i18n.subtitle)}</p>
        <div class="flex flex-wrap gap-2" role="tablist" aria-label="${escapeHTML(i18n.tablist_aria_label)}">
          ${tabButton("health", i18n.tabs.health, true)}
          ${tabButton("assets", i18n.tabs.assets, false)}
          ${tabButton("request", i18n.tabs.request, false)}
        </div>
        ${tabPanel("health", i18n.tabs.health, i18n.loading_tab, true)}
        ${tabPanel("assets", i18n.tabs.assets, i18n.loading_tab, false)}
        ${tabPanel("request", i18n.tabs.request, i18n.loading_tab, false)}
        <hr class="separator" aria-hidden="true" />
        <button
          id="fastygo-dev-hide"
          class="button w-fit"
          type="button"
          aria-label="${escapeHTML(i18n.hide_button_aria_label)}"
        >${escapeHTML(i18n.hide_button_label)}</button>
      </div>
    </section>
  `;
}

class FastyGoDevOverlay extends HTMLElement {
  cleanup: (() => void) | null = null;

  connectedCallback(): void {
    if (this.shadowRoot) return;

    const i18n = requirePanelI18n();
    if (!i18n) return;

    const root = this.attachShadow({ mode: "open" });
    root.innerHTML = overlayMarkup(i18n);

    const context: DevContext = {
      requestId: readRequestId(),
      path: readPath(),
      locale: readLocale(i18n.request.locale_missing),
      i18n,
      fetchStatus: () => fetchStatus(i18n),
      disableOverlay: () => disableOverlay(i18n),
    };

    const cleanupLauncher = bindLauncher(root);
    const cleanupTabs = bindTabs(root, context);
    const hide = root.querySelector<HTMLButtonElement>("#fastygo-dev-hide");
    const mobileReload = root.querySelector<HTMLElement>("#fastygo-dev-mobile-reload");
    const mobileReloadButton = root.querySelector<HTMLButtonElement>(
      "#fastygo-dev-mobile-reload-button",
    );
    const mobileQuery = window.matchMedia("(max-width: 767px)");
    const hideHandler = () => {
      void context.disableOverlay();
    };
    const reloadHandler = () => {
      window.location.reload();
    };
    const mobileChangeHandler = (event: MediaQueryListEvent) => {
      mobileReload?.classList.toggle("hidden", !event.matches);
    };

    hide?.addEventListener("click", hideHandler);
    mobileReloadButton?.addEventListener("click", reloadHandler);
    mobileQuery.addEventListener("change", mobileChangeHandler);
    mobileReload?.classList.toggle("hidden", !mobileQuery.matches);

    this.cleanup = () => {
      hide?.removeEventListener("click", hideHandler);
      mobileReloadButton?.removeEventListener("click", reloadHandler);
      mobileQuery.removeEventListener("change", mobileChangeHandler);
      cleanupLauncher();
      cleanupTabs();
    };
  }

  disconnectedCallback(): void {
    this.cleanup?.();
    this.cleanup = null;
  }
}

export function mountDevOverlay(): void {
  if (isMobileViewport()) {
    return;
  }
  if (!customElements.get(elementName)) {
    customElements.define(elementName, FastyGoDevOverlay);
  }
  if (!document.querySelector(elementName)) {
    document.body.append(document.createElement(elementName));
  }
}
