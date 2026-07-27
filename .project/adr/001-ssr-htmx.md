# ADR 001 — SSR instead of SPA

## Decision

Serve the app as Go SSR with templ. Progressive enhancement for interactions
(e.g. `@ui8kit/aria` Sheet); no client SPA router.

## Why

FastyGo Framework is Feature-based SSR. A 1:1 React SPA would fight the stack.
Document shells and registry composition stay server-owned.

## Consequences

Path URLs replace hash routes. Full page loads work without JS for read screens.
Client JS is limited to covered patterns (theme toggle, sheet, overlay).
