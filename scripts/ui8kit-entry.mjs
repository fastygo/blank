import { registerPattern, dialog } from "@ui8kit/aria";

const mobileSheetTriggerID = "ui8kit-mobile-sheet-trigger";
const mobileSheetPanelID = "ui8kit-mobile-sheet-panel";

registerPattern(dialog);

function bindMobileSheetTrigger() {
  const trigger = document.getElementById(mobileSheetTriggerID);
  const panel = document.getElementById(mobileSheetPanelID);
  if (!trigger || !panel) return;

  const overlay = panel.querySelector("[data-ui8kit-dialog-overlay]");
  const focusableSelector =
    'a[href], button, input, select, textarea, [tabindex]:not([tabindex="-1"])';

  function stopEvent(event) {
    event.preventDefault();
    event.stopPropagation();
    if (typeof event.stopImmediatePropagation === "function") {
      event.stopImmediatePropagation();
    }
  }

  function openPanel(event) {
    if (event) stopEvent(event);

    panel.setAttribute("data-state", "open");
    panel.removeAttribute("hidden");
    overlay?.removeAttribute("hidden");
    trigger.setAttribute("aria-expanded", "true");

    const firstFocusable = panel.querySelector(focusableSelector);
    firstFocusable?.focus();
  }

  function closePanel(event) {
    if (event) stopEvent(event);

    panel.setAttribute("data-state", "closed");
    panel.setAttribute("hidden", "hidden");
    overlay?.setAttribute("hidden", "hidden");
    trigger.setAttribute("aria-expanded", "false");
    trigger.focus();
  }

  trigger.addEventListener("pointerdown", openPanel, { capture: true });
  trigger.addEventListener("click", openPanel, { capture: true });
  trigger.addEventListener("keydown", (event) => {
    if (event.key === "Enter" || event.key === " ") {
      openPanel(event);
    }
  });

  for (const close of panel.querySelectorAll("[data-ui8kit-dialog-close]")) {
    close.addEventListener("pointerdown", closePanel, { capture: true });
    close.addEventListener("click", closePanel, { capture: true });
  }

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape") {
      closePanel(event);
    }
  });
}

if (typeof document !== "undefined") {
  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", bindMobileSheetTrigger, { once: true });
  } else {
    bindMobileSheetTrigger();
  }
}
