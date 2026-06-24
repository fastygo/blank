> **Статус:** исторический брейншторм по архитектуре сайдбаров. Действующая реализация — `layout.SidebarShell` (named shell в `internal/ui/layout/`) + `appsidebar.AppSidebar` (локальный компонент в `internal/ui/components/`). Страница композирует оба напрямую. См. [`specs/page-composes-layout.md`](specs/page-composes-layout.md) и [`specs/next-shadcn-architecture.md`](specs/next-shadcn-architecture.md).

В React/shadcn **не делают один движок на все комбинации**. Делают **малые enum на примитивах + compound parts + nested layouts + copy-paste blocks**. Для Go/templ это как раз ближе к [`templ-component-spec.mdc`](e:\_@Go\@Templ\.cursor\rules\templ-component-spec.mdc), чем к гигантскому `Variant`.

---

## Как это устроено в Next + shadcn

### 1. Nested layouts (главный механизм масштаба)

```text
app/
  layout.tsx              ← html, providers, fonts
  (marketing)/layout.tsx  ← top header + footer, без sidebar
  (app)/layout.tsx        ← SidebarProvider + Sidebar + SidebarInset
  (docs)/layout.tsx       ← header full width + toc aside справа
  (app)/dashboard/page.tsx
```

Каждый **route group** — свой ** JSX-дерево**, не один JSON-config. Комбинаций сколько угодно: вы **добавляете файл layout**, а не enum.

**Аналог в Blank:** не один `APP_LAYOUT` на всё приложение, а:
- `views/layout_marketing.templ`
- `views/layout_app.templ`
- `views/layout_docs.templ`

или route-level wrapper в `site/feature` (какой `SiteShell` вызвать).

### 2. shadcn Sidebar — compound + малые оси

Современный [shadcn Sidebar](https://ui.shadcn.com/docs/components/sidebar) — не «preset sidebars-main», а **набор частей**:

| Часть | Роль |
|--------|------|
| `SidebarProvider` | состояние open/collapsed (client + cookie) |
| `Sidebar` | `side`, `variant`, `collapsible` |
| `SidebarContent` | nav внутри |
| `SidebarInset` | main column |
| `SidebarTrigger` | burger |
| `Sheet` (внутри) | mobile offcanvas |

**API (bounded enums), не бесконечный preset:**

```tsx
<Sidebar side="left|right" variant="sidebar|floating|inset" collapsible="offcanvas|icon|none" />
```

Mobile: при `collapsible="offcanvas"` sidebar **сам** становится sheet; trigger появляется там, где вы положили `SidebarTrigger` (часто в header inset).

**Комбинации** = перемножение **нескольких осей** (3×3×3), а не тысяча имён пресетов.

### 3. Blocks — не runtime, а copy-paste

shadcn **Blocks** (`dashboard-01`, `sidebar-07`) — готовое дерево компонентов. Нужна новая геометрия → **копируете block и правите JSX**, а не `layoutEngine.setPreset()`.

У вас в templ уже тот же дух: [`blocks/dashboard`](e:\_@Go\@Templ\examples\ui\blocks\dashboard\dashboard.spec.md) — «mirrors application shell», self-contained, **не** product runtime.

---

## Сопоставление: React → templ (по spec.mdc)

| React / shadcn | templ / FastyGo (как в spec) |
|----------------|------------------------------|
| `<SidebarProvider>` | ui8kit sheet + `data-ui8kit` (серверный markup, клиент — `@ui8kit/aria`) |
| `{children}` | `{ children... }` + `templ.Component` slots |
| `layout.tsx` | `@layout.Shell(...) { @body }` в `views/layout.templ` |
| compound parts | composite `parts[]` в `layout.spec.md` |
| `side`, `variant`, `collapsible` | `api:` enums в spec, **полный контракт** |
| block copy-paste | `layout/showcase` + optional `internal/ui/blocks/*` |
| route groups | несколько shell wrappers / handlers |

Правило из spec: **читать `api.*.enum`, showcase — только примеры**. Значит layout должен быть **composite с parts**, а не один `ShellProps` на 40 полей.

---

## Предлагаемая структура (shadcn-style, масштабируется)

### Уровень 0 — примитивы (`templ/ui` + `templ/components`)

Уже есть: `Button`, `Block`, `Nav` (`Orientation: vertical|horizontal`). Не раздуваем shell.

### Уровень 1 — layout composite (`internal/ui/layout/`)

Как `components/nav` с `parts[]`:

```yaml
# layout.spec.md (concept)
parts:
  - templ: Shell
  - templ: ShellBand      # header/footer slot host
  - templ: AsideRegion    # left|right aside
  - templ: MainColumn     # SidebarInset analog
  - templ: NavList
  - templ: MobileSheet
  - templ: SheetTrigger   # burger
  - templ: PageHeader
  - templ: AppFooter

api:
  AsideSide: [left, right]
  AsideScope: [viewport, content_row]
  BandScope: [shell_full, main_column]
  BandSlot: [header, footer, subheader]
  Collapsible: [none, offcanvas, icon]   # shadcn parity
  SheetTriggerPlacement: [header_start, header_end, main_start, main_end]
```

**`AsideRegion`** — один компонент, два экземпляра (left/right):

```go
type AsideRegionProps struct {
    Side        string // api enum
    Scope       string
    Collapsible string // none | offcanvas
    Nav         NavListProps
    SheetTrigger SheetTriggerProps // если offcanvas на mobile
    // desktop classes from Scope × Side
}
```

Три ваших wireframe — **не три Variant**, а **два `AsideRegion` + набор `ShellBand`** в showcase:

| Showcase id | Aside L | Aside R | Bands |
|-------------|---------|---------|-------|
| `layout.sidebar_app` | viewport, offcanvas | — | header@main |
| `layout.sidebars_header` | content_row | content_row | header@shell_full |
| `layout.sidebars_main` | content_row | content_row | header+footer@shell_full |
| `layout.sidebars_full` | viewport | content_row | header@main |

### Уровень 2 — Shell composer (тонкий)

```templ
templ Shell(p ShellProps) {
  // document shell
  for _, sheet := range p.MobileSheets { @MobileSheet(sheet) }
  if p.ShellHeader != nil { @ShellBand(scope shell_full) { @p.ShellHeader } }
  @ShellGrid(p.Grid) {
    if p.AsideLeft.Enabled { @AsideRegion(p.AsideLeft) }
    @MainColumn(p.Main) {
      if p.MainHeader != nil { @ShellBand(scope main) { @p.MainHeader; triggers... } }
      @ShellRow {
        if p.AsideLeft.ContentRow { @AsideRegion(...) }
        @Main { { children... } }
        if p.AsideRight.Enabled { @AsideRegion(p.AsideRight) }
      }
      if p.MainFooter != nil { @p.MainFooter }
    }
  }
  if p.ShellFooter != nil { @ShellBand(shell_full) { @p.ShellFooter } }
}
```

**`ShellProps` — дерево частей**, не строки:

```go
type ShellProps struct {
    Meta LayoutMeta
    Grid GridSpec          // derived helper classes
    AsideLeft  AsideRegionProps
    AsideRight AsideRegionProps
    ShellHeader  templ.Component
    MainHeader   templ.Component
    ShellFooter  templ.Component
    MainFooter   templ.Component
    MobileSheets []MobileSheetProps // 0..2
}
```

Сборка дерева — в Go (`layout.BuildAppShell(d LayoutData)`) или в `views/layout.templ` — как в React вы пишете JSX в `layout.tsx`.

### Уровень 3 — preset / config (DX для Blank starter)

```go
// Preset = factory функция, возвращает ShellProps skeleton
func ShellFromPreset(p Preset, d LayoutData) ShellProps
```

`APP_LAYOUT=sidebars_main` → вызывает factory. **Power users** обходят preset и собирают `ShellProps` вручную (как fork shadcn block).

### Уровень 4 — route-specific layouts (Next route groups)

```go
// site/feature.go
func (f *Feature) getHome(...) {
  web.Render(..., views.MarketingShell(data, body))  // no aside
}
func (f *Feature) getDashboard(...) {
  web.Render(..., views.AppShell(data, body))         // sidebar_app preset
}
```

Это **главный** способ «бесконечных комбинаций» в production — не один global preset.

---

## Mobile sheet + burger (как shadcn)

В shadcn связка **Sidebar + collapsible=offcanvas + SidebarTrigger** декlarativная.

В templ:

```go
type AsideRegionProps struct {
    Collapsible string // none | offcanvas
    Sheet MobileSheetProps // panel id, aria, nav slot
    Trigger SheetTriggerProps // Placement, only rendered if Collapsible=offcanvas
}
```

- **Desktop:** `<aside class="hidden md:flex ...">` + `@NavList`
- **Mobile:** aside hidden; `@MobileSheet` с тем же `@NavList`; `@SheetTrigger` в band по `Placement`
- **Два aside** → два `MobileSheet` + два trigger (start/end header) — как два `SidebarTrigger` в React

`NavigationProps` из fixtures — на каждый sheet/trigger (как сейчас в Blank).

---

## Почему не один `LayoutSpec` JSON на всё приложение

| Подход | React world | Когда |
|--------|-------------|--------|
| **Nested layouts** | route groups | production, разные секции сайта |
| **Compound + enums** | Sidebar parts | переиспользуемая геометрия |
| **Blocks** | shadcn blocks | старт / showcase / CDN |
| **Preset map** | starter templates | Blank `APP_LAYOUT`, быстрый выбор |
| **Layout engine** | редко (CMS) | только если layout из БД |

shadcn **не** продаёт `layoutEngine.compute(sidebars_full)` — продаёт **компоненты и blocks**. Вам для templ разумный **hybrid**:

1. **Composite layout package** (parts + api enums + spec.md) — SoT по spec.mdc  
2. **3–5 preset factories** — wireframe + Blank starter  
3. **Несколько `views/layout_*.templ`** — как Next route groups  
4. **Blocks** для demo pages — optional  

---

## Связь с PageTitle / Footer

В shadcn это не props Sidebar, а **children areas**:

```tsx
<SidebarInset>
  <header className="..."><h1>{title}</h1></header>
  <main>{children}</main>
  <footer>...</footer>
</SidebarInset>
```

В templ:

```templ
@layout.PageHeader(p)   // slot → MainHeader или ShellHeader
@layout.AppFooter(p)    // slot → MainFooter или ShellFooter
```

**Куда** монтируется — решает `ShellBand` + preset factory, не сам `PageHeader`.

---

## Практический вывод для Blank

1. **Не раздувать один enum Variant** — сделать **composite `layout.spec.md`** с `parts` и osями (`AsideSide`, `AsideScope`, `Collapsible`, `BandScope`, `TriggerPlacement`).

2. **Три картинки** — только **`showcase: layout.sidebars_*`**, не runtime names.

3. **Blank default** — `ShellFromPreset(PresetTopNav | PresetSidebarApp, d)`.

4. **FastyGoUI** уже ближе: `HeaderExtra`, `SidebarProps.Mobile`, sections — это **slots + parts**, развивать в ту же модель.

5. **Перенос BlankSidebar** — preset `sidebar_app`, не отдельный репо.

6. **Масштаб как в Next** — когда появится docs/marketing/dashboard: **отдельный layout wrapper**, общие `AsideRegion`/`MobileSheet`.

---

Если пойдёте в Agent mode, разумный первый slice в духе shadcn: **`layout.spec.md` + parts `AsideRegion`, `MobileSheet`, `SheetTrigger`, `NavList`, `ShellBand`** + preset factories `topnav` / `sidebar_app` + один composer `Shell` — без всех wireframe preset до зелёных тестов на parts.