/**
 * List props and helpers for List and ListItem templ components.
 */

package list

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/fastygo/templ/utils"
)

// ListProps configures list tag and classes.
type ListProps struct {
	Class string
	// Tag — ul, ol, dl, or menu
	Tag   string
	Attrs templ.Attributes
}

// ListItemProps configures one list row.
type ListItemProps struct {
	Class string
	// Tag — li, dt, or dd; defaults to li
	Tag string
	// Value — optional li value for ordered contexts
	Value int
	Attrs templ.Attributes
}

func listTag(tag string) string {
	return uiutils.ResolveTag(tag, "ul", uiutils.TagGroupList)
}

func listItemAttrs(p ListItemProps) templ.Attributes {
	if p.Value <= 0 || uiutils.ResolveTag(p.Tag, "li", uiutils.TagGroupListItem) != "li" {
		return templ.Attributes{}
	}
	return templ.Attributes{"value": strconv.Itoa(p.Value)}
}

func listItemMergedAttrs(p ListItemProps) templ.Attributes {
	return uiutils.MergeAttrs(listItemAttrs(p), p.Attrs)
}

func listSpreadAttrs(p ListProps) templ.Attributes {
	return uiutils.MergeAttrs(templ.Attributes{}, p.Attrs)
}
