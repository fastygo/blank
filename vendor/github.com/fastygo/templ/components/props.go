/**
 * Package components is the composite registry facade.
 * Import cmp "github.com/fastygo/templ/components" for Card, Alert, Breadcrumb.
 */

// Package components is the composite registry facade. Import as:
//
//	import cmp "github.com/fastygo/templ/components"
//
// Use @cmp.Card, @cmp.Alert, … Granular imports remain valid.
package components

import (
	"github.com/fastygo/templ/components/alert"
	"github.com/fastygo/templ/components/breadcrumb"
	"github.com/fastygo/templ/components/card"
	"github.com/fastygo/templ/components/iconbadge"
	"github.com/fastygo/templ/components/nav"
	"github.com/fastygo/templ/components/sheet"
)

type (
	CardProps             = card.CardProps
	CardHeaderProps       = card.CardHeaderProps
	CardTitleProps        = card.CardTitleProps
	CardDescriptionProps  = card.CardDescriptionProps
	CardContentProps      = card.CardContentProps
	CardFooterProps       = card.CardFooterProps
	AlertProps            = alert.AlertProps
	BreadcrumbProps       = breadcrumb.BreadcrumbProps
	BreadcrumbItem        = breadcrumb.BreadcrumbItem
	SheetProps            = sheet.SheetProps
	SheetTriggerProps     = sheet.SheetTriggerProps
	SheetOverlayProps     = sheet.SheetOverlayProps
	SheetContentProps     = sheet.SheetContentProps
	SheetHeaderProps      = sheet.SheetHeaderProps
	SheetTitleProps       = sheet.SheetTitleProps
	SheetDescriptionProps = sheet.SheetDescriptionProps
	SheetCloseProps       = sheet.SheetCloseProps
	NavProps              = nav.NavProps
	NavListProps          = nav.NavListProps
	NavItemProps          = nav.NavItemProps
	NavLinkProps          = nav.NavLinkProps
	IconBadgeProps        = iconbadge.IconBadgeProps
)

func CardClasses(p CardProps) string           { return card.CardClasses(p) }
func AlertClasses(p AlertProps) string         { return alert.AlertClasses(p) }
func SheetClasses(p SheetProps) string         { return sheet.SheetClasses(p) }
func NavListClasses(p NavListProps) string     { return nav.NavListClasses(p) }
func NavLinkClasses(p NavLinkProps) string     { return nav.NavLinkClasses(p) }
func IconBadgeClasses(p IconBadgeProps) string { return iconbadge.IconBadgeClasses(p) }
