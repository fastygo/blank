import { mountDevOverlay } from "../internal/devoverlay/frontend/main";

function mountWhenReady(): void {
  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", mountDevOverlay, { once: true });
    return;
  }
  mountDevOverlay();
}

(
  globalThis as {
    __FASTYGO_DEV_OVERLAY_MOUNT__?: () => void;
  }
).__FASTYGO_DEV_OVERLAY_MOUNT__ = mountDevOverlay;

mountWhenReady();
window.addEventListener("pageshow", mountDevOverlay);
