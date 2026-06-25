Да — **можно и, с точки зрения shadcn/Next, это даже логичнее**. Сейчас `Shell` — не «document shell», а **готовый app chrome** (header + main + footer + sheet).

## Что есть сейчас

```8:47:e:\_@Go\@Blank\internal\ui\layout\shell.templ
templ Shell(props ShellProps) {
	<!DOCTYPE html>
	<html ...>
		<head>...</head>
		<body>
			Box (flex col)
				Header          ← app chrome
				main { children }
				Footer
			MobileSheet         ← app chrome
		</body>
	</html>
}
```

Страница кладёт контент **внутрь уже готового `<main>`**, а header/footer/sheet зашиты в Shell. Отсюда и расхождение со скрином: sidebar grid живёт **под** глобальным header, а не внутри «inset».

## Если ограничить Shell только document frame

Идея:

```templ
templ Document(props DocumentProps) {
	<!DOCTYPE html>
	<html lang={...}>
		<head>
			<title>...</title>
			@ShellHead(...)
		</head>
		<body class={...}>
			{ children... }
		</body>
	</html>
}
```

Всё остальное — **снаружи Shell, в page или в layout-organism**:

```text
Document
  └── TopnavApp { Header, main, Footer, MobileSheet }     ← home
  └── DashboardApp { Sidebar, Inset{ PageHeader, content } } ← sample
```

Это ближе к:

- Next: `app/layout.tsx` = только `html/body`, route layout = всё внутри  
- shadcn: `SidebarProvider` + `SidebarInset` — **в page**, не в root

## Почему так не сделали изначально

1. **Blank как demo app** — один и тот же header/footer на все route «из коробки»; меньше кода в каждой page.
2. **Один `layout.Data` → один `ShellProps`** — nav, theme, i18n, assets собираются раз и прокидываются в Header одним махом.
3. **Wireframe-фаза** — быстрее показать working shell, чем заставлять каждую page собирать chrome.
4. **Mobile sheet** привязан к global header (trigger в `header.templ`) — удобно, пока один layout на всё.

То есть это **осознанный компромисс «стартер с батарейками»**, не ошибка.

## Плюсы узкого Shell

| Плюс | Для кого |
|------|----------|
| Page = полная картина макета (как `page.tsx`) | React/shadcn |
| Нет «скрытого» header над sidebar | Dashboard routes |
| Copy-paste blocks могут включать свой header/sidebar | Registry |
| `SidebarShell` не оборачивает Shell с лишним Header | Архитектура |
| Два layout-варианта не борются в одном компоненте | Поддержка |

## Минусы / цена

| Минус | Что делать |
|-------|------------|
| Home и Sample дублируют сборку chrome | `layout/topnav_app.templ`, `layout/dashboard_app.templ` |
| Каждая page должна выбрать organism | Явно в `home.templ` / `sample.templ` — это плюс для читаемости |
| `ShellProps` раздуется или распадётся | `DocumentProps` (title, lang, assets) + `ChromeProps` (nav, theme) |
| Mobile sheet + trigger — куда? | В тот organism, где header (topnav или page header) |

## Нужно ли это вам?

**Да, если** Blank — onboarding для React devs и `/sample` = эталон dashboard.

**Можно отложить, если** приоритет — минимум файлов на route и один chrome на всё приложение.

Компромисс без большого рефактора:

- **`Document`** — только `html/head/body` + assets  
- **`TopnavChrome`** — Header + main slot + Footer + MobileSheet (текущий home)  
- **`DashboardChrome`** — sidebar + inset column без global header (shadcn-like sample)

Home: `@Document { @TopnavChrome { @hero.Hero } }`  
Sample: `@Document { @DashboardChrome { @PageHeader; content } }`

## Ответ на «почему бы просто не…»

**Можно.** Ограничение Shell только document frame — **правильное направление** для shadcn-пarity и снимает пункт 8 из аудита.

Сейчас Shell шире `<body>{children}</body>`, потому что Blank задумывался как **app with fixed chrome**, а не как **minimal document + composable layouts**.

Если идти этим путём — не удалять header/footer, а **вынести их из Shell** в named organisms; Shell переименовать по смыслу в `Document` (или оставить имя, но сузить контракт).

Стоит ли набросать конкретный slice: `Document` + `TopnavChrome` + правка `home.templ`/`sample.templ`? Или еще остались подводные камни? 