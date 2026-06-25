# `appsidebar`

Local sidebar component for route pages that compose `layout.DashboardLayout`.

This is the templ analogue of shadcn's `components/app-sidebar.tsx`: a small,
props-only aside meant to be **copy-pasted between projects** and customised
in place. It is rendered inside `layout.DashboardLayout`.

## Contract

```go
appsidebar.AppSidebar(appsidebar.Props{
    Title:     "Brand",
    AriaLabel: "Sidebar navigation",
    Items:     []navigation.Item{...},
    Active:    "/sample",
    Class:     "",          // optional extra utility classes
})
```

Renders:

```text
<aside aria-label=...>
  <Stack>
    <p>{Title}</p>
    <nav>
      navigation.Nav (vertical, props.Items, aria-current on active)
    </nav>
  </Stack>
</aside>
```

The aside is `hidden` below `md` — mobile users see the same nav items via the
TopnavLayout mobile sheet.

## Usage in pages

```templ
templ SamplePage(d layout.Data, f fixtures.Locale) {
    @layout.RootLayout(d.Document()) {
        @layout.DashboardLayout(d.Dashboard(f.Sample.Title)) {
            @ui.Box(...) { ...page content... }
        }
    }
}
```

Build sidebar `Props` from `layout.Data` via `d.Dashboard(title).Sidebar`.

## Replacing or extending

`AppSidebar` is intentionally simple. Forking it for a project-specific aside
(version switcher, search field, collapsible groups) is the supported path —
copy the package, edit the templ, keep the `Props` contract or change it as
needed. There is no global sidebar engine to migrate.

## Related

- [`../navigation/`](../navigation/) — `Nav`, `MobileSheet`, `MobileSheetTrigger`
- [`../../layout/dashboard_layout.templ`](../../layout/dashboard_layout.templ) — host layout layer
- [`../../../../docs/for-react-devs.md`](../../../../docs/for-react-devs.md) — cookbook
