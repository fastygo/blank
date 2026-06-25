# Shell layout refactor — resolved

This draft proposed narrowing `Shell` to a document-only frame and moving app chrome into named layout layers. **Implemented** in the root layout split refactor.

## What changed

| Before | After |
|--------|-------|
| `layout.Shell` (document + header + footer + sheet) | `layout.RootLayout` (document only) + `layout.TopnavLayout` (chrome) |
| `layout.SidebarShell` | `layout.DashboardLayout` (TopnavLayout + aside) |
| `ShellProps` | `DocumentProps`, `TopnavLayoutProps`, `DashboardLayoutProps` |
| `d.ShellProps()`, `d.SidebarProps(title)` | `d.Document()`, `d.Topnav()`, `d.Dashboard(title)` |
| `views.HomePageFrom(d, f)` | `views.HomePage(d, f)` |

## Page composition (current)

```templ
// Home — topnav
templ HomePage(d layout.Data, f fixtures.Locale) {
    @layout.RootLayout(d.Document()) {
        @layout.TopnavLayout(d.Topnav()) { ... }
    }
}

// Sample — dashboard
templ SamplePage(d layout.Data, f fixtures.Locale) {
    @layout.RootLayout(d.Document()) {
        @layout.DashboardLayout(d.Dashboard(f.Sample.Title)) { ... }
    }
}
```

## Docs

- [`internal/ui/layout/README.md`](../../internal/ui/layout/README.md)
- [`docs/for-react-devs.md`](../../docs/for-react-devs.md)
