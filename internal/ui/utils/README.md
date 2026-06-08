# `registry:utils`

Thin app helpers on top of **`github.com/fastygo/templ/utils`**. Do not duplicate CVA, tag resolution, or ARIA — delegate to templ.

Add functions here only when they are **app-specific** (registry labels, merge helpers used across blocks).
Generic class/tag logic belongs in **`github.com/fastygo/templ/utils`**.
