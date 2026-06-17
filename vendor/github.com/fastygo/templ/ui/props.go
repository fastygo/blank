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
	"github.com/fastygo/templ/ui/dialog"
	"github.com/fastygo/templ/ui/disclosure"
	"github.com/fastygo/templ/ui/form"
	"github.com/fastygo/templ/ui/grid"
	"github.com/fastygo/templ/ui/group"
	"github.com/fastygo/templ/ui/icon"
	"github.com/fastygo/templ/ui/image"
	"github.com/fastygo/templ/ui/input"
	"github.com/fastygo/templ/ui/label"
	"github.com/fastygo/templ/ui/linebreak"
	"github.com/fastygo/templ/ui/link"
	"github.com/fastygo/templ/ui/list"
	"github.com/fastygo/templ/ui/radio"
	"github.com/fastygo/templ/ui/select"
	"github.com/fastygo/templ/ui/separator"
	"github.com/fastygo/templ/ui/stack"
	"github.com/fastygo/templ/ui/switch"
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
	IconProps            = icon.IconProps
	ImageProps           = image.ImageProps
	PictureProps         = image.PictureProps
	SourceProps          = image.SourceProps
	LinkProps            = link.LinkProps
	SeparatorProps       = separator.SeparatorProps
	DisclosureProps      = disclosure.DisclosureProps
	SummaryProps         = disclosure.SummaryProps
	DialogProps          = dialog.DialogProps
	BreakProps           = linebreak.BreakProps
	TextProps            = text.TextProps
	TitleProps           = title.TitleProps
	ButtonProps          = button.ButtonProps
	BadgeProps           = badge.BadgeProps
	InputProps           = input.InputProps
	TextareaProps        = textarea.TextareaProps
	SelectProps          = selectfield.SelectProps
	OptionProps          = selectfield.OptionProps
	OptGroupProps        = selectfield.OptGroupProps
	CheckboxProps        = checkbox.CheckboxProps
	RadioProps           = radio.RadioProps
	SwitchProps          = formswitch.SwitchProps
	LabelProps           = label.LabelProps
	FormProps            = form.FormProps
	FormItemProps        = form.FormItemProps
	FormDescriptionProps = form.FormDescriptionProps
	FormMessageProps     = form.FormMessageProps
	FieldsetProps        = form.FieldsetProps
	LegendProps          = form.LegendProps
	DataListProps        = form.DataListProps
	DataOptionProps      = form.DataOptionProps
	OutputProps          = form.OutputProps
	MeterProps           = form.MeterProps
	ProgressProps        = form.ProgressProps
)

type Variants = uiutils.Variants

var (
	ButtonVariants = button.ButtonVariants
	BadgeVariants  = badge.BadgeVariants
	IconVariants   = icon.IconVariants
	ImageVariants  = image.ImageVariants
	LinkVariants   = link.LinkVariants
)

func StackClasses(p StackProps) string       { return stack.StackClasses(p) }
func GroupClasses(p GroupProps) string       { return group.GroupClasses(p) }
func GridClasses(p GridProps) string         { return grid.GridClasses(p) }
func GridColClasses(p GridColProps) string   { return grid.GridColClasses(p) }
func IconClasses(p IconProps) string         { return icon.IconClasses(p) }
func ImageClasses(p ImageProps) string       { return image.ImageClasses(p) }
func LinkClasses(p LinkProps) string         { return link.LinkClasses(p) }
func ButtonClasses(p ButtonProps) string     { return button.ButtonClasses(p) }
func BadgeClasses(p BadgeProps) string       { return badge.BadgeClasses(p) }
func InputClasses(p InputProps) string       { return input.InputClasses(p) }
func TextareaClasses(p TextareaProps) string { return textarea.TextareaClasses(p) }
func SelectClasses(p SelectProps) string     { return selectfield.SelectClasses(p) }
func CheckboxClasses(p CheckboxProps) string { return checkbox.CheckboxClasses(p) }
func RadioClasses(p RadioProps) string       { return radio.RadioClasses(p) }
func SwitchClasses(p SwitchProps) string     { return formswitch.SwitchClasses(p) }
func FieldsetClasses(p FieldsetProps) string { return form.FieldsetClasses(p) }
func LegendClasses(p LegendProps) string     { return form.LegendClasses(p) }
func MeterClasses(p MeterProps) string       { return form.MeterClasses(p) }
func ProgressClasses(p ProgressProps) string { return form.ProgressClasses(p) }
