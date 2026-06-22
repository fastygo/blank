import type { DevContext, DevPanel } from "../types";

function field(
  root: HTMLElement,
  label: string,
  value: string,
  copy: DevContext["i18n"]["request"],
  copyButton = false,
): void {
  const row = document.createElement("div");
  row.className = "rounded-md border border-border p-3 flex flex-col gap-2";

  const title = document.createElement("p");
  title.className = "text-sm font-medium";
  title.textContent = label;

  const valueEl = document.createElement("code");
  valueEl.className = "block break-all text-sm";
  valueEl.textContent = value || copy.empty_value;

  row.append(title, valueEl);

  if (copyButton && value) {
    const button = document.createElement("button");
    button.type = "button";
    button.className =
      "inline-flex h-8 w-fit items-center justify-center rounded-md border border-border bg-background px-3 text-sm hover:bg-accent";
    button.textContent = copy.copy_request_id;
    button.addEventListener("click", async () => {
      try {
        await navigator.clipboard.writeText(value);
        button.textContent = copy.copied;
      } catch {
        button.textContent = copy.copy_failed;
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
    const copy = context.i18n.request;
    root.replaceChildren();

    const intro = document.createElement("p");
    intro.className = "text-sm text-muted-foreground";
    intro.textContent = copy.intro;
    root.append(intro);

    field(root, copy.request_id_label, context.requestId, copy, true);
    field(root, copy.path_label, context.path, copy);
    field(root, copy.locale_label, context.locale, copy);

    return () => {};
  },
};
