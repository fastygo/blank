import { registerPattern, dialog } from "@ui8kit/aria";

// registerPattern schedules dialog.init on DOMContentLoaded (or immediately when the
// document is already interactive). Do not call init() synchronously here — on some
// mobile browsers defer scripts can run before the body nodes are fully wired.
registerPattern(dialog);
