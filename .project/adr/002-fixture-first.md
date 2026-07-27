# ADR 002 — Fixture-first UI copy

## Decision

UI chrome and demo page strings live in `internal/fixtures/locale` (embedded JSON
per locale). Views receive resolved `fixtures.Locale`; they do not load JSON.

## Why

UX/UI debugging and i18n work without waiting on CMS. Swap to a real content
source later without rewriting templ composition.

## Consequences

Add a string → update fixtures structs + every locale JSON. Keep domain/API
data out of locale fixtures when a product grows a domain layer.
