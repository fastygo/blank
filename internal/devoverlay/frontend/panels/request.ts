import type { DevContext, DevPanel } from "../types";

function field(root: HTMLElement, label: string, value: string, copy = false): void {
  const row = document.createElement("div");
  row.className = "rounded-md border border-border p-3 flex flex-col gap-2";

  const title = document.createElement("p");
  title.className = "text-sm font-medium";
  title.textContent = label;

  const valueEl = document.createElement("code");
  valueEl.className = "block break-all text-sm";
  valueEl.textContent = value || "(empty)";

  row.append(title, valueEl);

  if (copy && value) {
    const button = document.createElement("button");
    button.type = "button";
    button.className =
      "inline-flex h-8 w-fit items-center justify-center rounded-md border border-border bg-background px-3 text-sm hover:bg-accent";
    button.textContent = "Copy request id";
    button.addEventListener("click", async () => {
      try {
        await navigator.clipboard.writeText(value);
        button.textContent = "Copied";
      } catch {
        button.textContent = "Copy failed";
      }
    });
    row.append(button);
  }

  root.append(row);
}

export const requestPanel: DevPanel = {
  id: "request",
  title: "Request",
  mount(root, context) {
    root.replaceChildren();

    const intro = document.createElement("p");
    intro.className = "text-sm text-muted-foreground";
    intro.textContent = "Use the request id to match browser traffic with Go structured logs.";
    root.append(intro);

    field(root, "Request ID", context.requestId, true);
    field(root, "Path", context.path);
    field(root, "Locale", context.locale);

    return () => {};
  },
};
