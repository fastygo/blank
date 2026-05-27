/**
 * Package ui is the registry facade for all primitives.
 * Import as github.com/fastygo/templ/ui — use @ui.Button, @ui.Block, etc.
 * Granular imports (github.com/fastygo/templ/ui/button) remain valid.
 */

// Package ui is the registry facade. Import as:
//
//	import "github.com/fastygo/templ/ui"
//
// Use @ui.Button, @ui.Block, … Granular imports (github.com/fastygo/templ/ui/button) remain valid.
package ui

import (
	"github.com/fastygo/templ/ui/badge"
	"github.com/fastygo/templ/ui/block"
	"github.com/fastygo/templ/ui/box"
	"github.com/fastygo/templ/ui/button"
	"github.com/fastygo/templ/ui/checkbox"
	"github.com/fastygo/templ/ui/container"
	"github.com/fastygo/templ/ui/form"
	"github.com/fastygo/templ/ui/switch"
	"github.com/fastygo/templ/ui/grid"
	"github.com/fastygo/templ/ui/group"
	"github.com/fastygo/templ/ui/input"
	"github.com/fastygo/templ/ui/label"
	"github.com/fastygo/templ/ui/list"
	"github.com/fastygo/templ/ui/radio"
	"github.com/fastygo/templ/ui/select"
	"github.com/fastygo/templ/ui/stack"
	"github.com/fastygo/templ/ui/table"
	"github.com/fastygo/templ/ui/text"
	"github.com/fastygo/templ/ui/textarea"
	"github.com/fastygo/templ/ui/title"
	uiutils "github.com/fastygo/templ/utils"
)

type (
	BlockProps           = block.BlockProps
	BoxProps             = box.BoxProps
	ContainerProps       = container.ContainerProps
	StackProps           = stack.StackProps
	GroupProps           = group.GroupProps
	ListProps            = list.ListProps
	ListItemProps        = list.ListItemProps
	TableProps           = table.TableProps
	TableCaptionProps    = table.TableCaptionProps
	TableSectionProps    = table.TableSectionProps
	TableRowProps        = table.TableRowProps
	TableCellProps       = table.TableCellProps
	TableColGroupProps   = table.TableColGroupProps
	TableColProps        = table.TableColProps
	GridProps            = grid.GridProps
	GridColProps         = grid.GridColProps
	TextProps            = text.TextProps
	TitleProps           = title.TitleProps
	ButtonProps          = button.ButtonProps
	BadgeProps           = badge.BadgeProps
	InputProps           = input.InputProps
	TextareaProps        = textarea.TextareaProps
	SelectProps          = selectfield.SelectProps
	CheckboxProps        = checkbox.CheckboxProps
	RadioProps           = radio.RadioProps
	SwitchProps          = formswitch.SwitchProps
	LabelProps           = label.LabelProps
	FormProps            = form.FormProps
	FormItemProps        = form.FormItemProps
	FormDescriptionProps = form.FormDescriptionProps
	FormMessageProps     = form.FormMessageProps
)

type Variants = uiutils.Variants

var (
	ButtonVariants = button.ButtonVariants
	BadgeVariants  = badge.BadgeVariants
)

func StackClasses(p StackProps) string       { return stack.StackClasses(p) }
func GroupClasses(p GroupProps) string       { return group.GroupClasses(p) }
func GridClasses(p GridProps) string         { return grid.GridClasses(p) }
func ButtonClasses(p ButtonProps) string     { return button.ButtonClasses(p) }
func BadgeClasses(p BadgeProps) string       { return badge.BadgeClasses(p) }
func InputClasses(p InputProps) string       { return input.InputClasses(p) }
func TextareaClasses(p TextareaProps) string { return textarea.TextareaClasses(p) }
func SelectClasses(p SelectProps) string     { return selectfield.SelectClasses(p) }
func CheckboxClasses(p CheckboxProps) string { return checkbox.CheckboxClasses(p) }
func RadioClasses(p RadioProps) string       { return radio.RadioClasses(p) }
func SwitchClasses(p SwitchProps) string     { return formswitch.SwitchClasses(p) }
