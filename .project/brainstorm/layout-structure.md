Вы правы — для человека это **не та же задача**, что для LLM, который держит в голове весь граф файлов. Скриншот как раз показывает, **почему shadcn лёгкий**, а текущий Blank — тяжёлый.

## Что делает shadcn простым (на вашем скрине)

На скрине видно **одну точку входа** — `app/dashboard/page.tsx`. Открываешь файл и **сразу читаешь макет**:

```text
SidebarProvider
  ├── AppSidebar          ← локальный компонент (components/app-sidebar.tsx)
  └── SidebarInset
        └── header
              ├── SidebarTrigger
              ├── Separator
              └── Breadcrumb …
```

Плюс три очевидные папки:

| Папка | Роль | Когнитивная нагрузка |
|-------|------|----------------------|
| `app/dashboard/page.tsx` | **Здесь собран макет страницы** | Низкая — всё видно в одном файле |
| `components/app-sidebar.tsx` | Ваш sidebar (search, version switcher) | Низкая — «мой код» |
| `components/ui/sidebar` | Скопированные примитивы shadcn | Низкая — «библиотека, не трогаю» |

**Нет** скрытых слоёв «document shell vs route shell vs runtime adapter vs block organism».  
**Нет** отдельного файла «куда зарегистрировать layout».  
**Нет** пустой обёртки, которая только вызывает другую обёртку.

Copy-paste здесь буквальный: скопировал `components/app-sidebar.tsx` + `components/ui/sidebar/*` + кусок из `page.tsx` — и понятно, что куда.

---

## Что происходит в Blank для того же «sidebar dashboard»

Чтобы понять `/sample`, вам нужно пройти **цепочку**, которой на скрине **нет**:

```text
router.go  →  views.SidebarAppShell  →  blocks/sidebar_app  →  layout.Shell
                                                      ↓
                                            layout/header + footer
                                                      ↓
                                            components/navigation (sheet)
                                                      ↓
                                            views/partials (CSS, i18n)
                                                      ↓
                                            site/layout_data.go (сборка props)
```

И при этом `sidebar_app.templ` — **39 строк**, а `app_shell.templ` — **4 строки**.  
Реальный sidebar+header+sheet на скрине — **в одном `page.tsx`**, у вас — **размазан по 6+ файлам и 3 «слоям»**.

LLM это «понятно», потому что мы видим всю карту сразу.  
Вы — **открываете один файл и не видите макет**. Отсюда и «колоссальная когнитивная нагрузка».

---

## Прямое сопоставление: скрин ↔ Blank

| На скрине (Next + shadcn) | Blank сейчас | Проблема |
|---------------------------|--------------|----------|
| `page.tsx` — **видимая композиция** | `views/sample_stub.templ` — только текст по центру | Макет **не в page** |
| `AppSidebar` в `components/` | Кусок aside в `blocks/sidebar_app/` | Остальное sidebar — в `layout/` |
| `SidebarProvider`, `SidebarInset` в `ui/sidebar` | Нет аналога; sheet в `navigation/` + `layout/shell` | Примитивы **не рядом** с макетом |
| Импорты `@/components/...` | Импорты через `layout.`, `sidebarapp.`, `views.`, `site.` | **4 namespace**, не 2 |
| Layout = JSX в page | Layout = `PageSpec.Layout` + adapter + block + Shell | **3 прыжка**, чтобы понять «где sidebar» |

В shadcn вы думаете: **«открыл page — вижу дерево»**.  
В Blank: **«открыл block — пусто; иди в layout; иди в router; иди в views»**.

---

## Почему рефакторинг усугубил ощущение, а не снял

Рефакторинг решал **архитектурную** задачу (runtime vs registry, router manifest, имена как в Next).  
Скрин решает **когнитивную** задачу: **один файл = одна картина макета**.

Это разные цели. Block 05–09 дали:

- правильные **имена** (`AppShell`, `SidebarAppShell`, `router.go`);
- но **не** shadcn-опыт «всё на месте в одной папке/файле».

`blocks/dashboard/sidebar_app/` — это **не** аналог `app-sidebar.tsx` + `page.tsx`.  
Это **фрагмент** aside внутри чужого `layout.Shell`.

---

## Как должно выглядеть «как на скрине» в Go/Templ (без CLI)

Цель — **та же когнитивная модель**, не обязательно те же имена файлов.

### Вариант, близкий к вашему скрину

```text
internal/views/dashboard/
  page.templ              ← АНАЛОГ page.tsx: Provider + Sidebar + Inset + header + {children}

internal/ui/components/
  app_sidebar.templ       ← АНАЛОГ app-sidebar.tsx
  search_form.templ
  version_switcher.templ

internal/ui/components/ui/   ← или vendor templ/components
  sidebar/                ← SidebarProvider, SidebarInset, SidebarTrigger (kit)

internal/site/router.go
  Pattern: "/dashboard"
  Body:    views.DashboardPage(...)   ← только page, БЕЗ Layout: views.SidebarAppShell
```

**Ключевое правило для мозга:**

> **Макет страницы живёт в `page.templ` рядом с route, а не в `blocks/*` + `layout/*` + adapter.**

`router.go` — одна строка «какой page рендерить».  
Вся композиция sidebar — **внутри page**, как на скрине.

### Что тогда с `layout/` и `blocks/`?

| Слой | Роль (упрощённо) |
|------|------------------|
| `layout/shell.templ` | Только `<html><head><body>` + assets — **как root layout.tsx**, без header/sidebar |
| `components/ui/*` | Скопированные примитивы (sheet, sidebar parts) — **не трогаем** |
| `components/app-sidebar.templ` | **Ваш** sidebar — copy-paste между проектами |
| `views/dashboard/page.templ` | **Сборка**, как на скрине |
| `blocks/` | Только когда **целиком** переносите готовый dashboard (как Templ `home`) — self-contained папка |

`blocks/dashboard/app_shell` с 4 строками при такой модели **просто не нужен** — он только добавляет прыжок.

---

## Одна фраза-якорь (как shadcn на скрине)

> **Route page composes layout. Components are next door. Primitives are in ui/.**

Перевод на Blank:

> **`views/<route>/page.templ` собирает макет. `components/` — ваш код. Kit — `templ/components`. Router только указывает page.**

Сейчас у вас наоборот:

> Router → adapter → block → layout → components → page (контент без макета)

---

## Что делать practically, пока CLI не нужен

1. **Новые страницы с sidebar** — делайте как на скрине: один `page.templ` с явной композицией, локальный `app_sidebar.templ` в `components/`.
2. **Не создавайте** новые «blocks» из 4 строк `@layout.Shell { @body }`.
3. **`sidebar_app` и `app_shell`** — считайте ошибкой раннего рефакторинга, не образцом; либо **наполните** block до уровня Templ `home`, либо **перенесите композицию в page**.
4. **`layout/`** — сведите к document frame (html/head/body); header/sidebar **вынесите в page или components**, как на скрине.

---

## Итог

Ваша реакция **нормальная и точная**.  
Проблема не в том, что вы «не поняли архитектуру» — проблема в том, что Blank сейчас оптимизирован под **граф зависимостей для LLM/архитектора**, а shadcn на скрине — под **один файл, который человек читает сверху вниз**.

Скрин — золотой стандарт когнитивной простоты:  
**page composes, components local, ui primitives copied, всё рядом.**

Blank до этого стандарта ещё не доведён: макет разнесён, blocks пустые, page — только контент.