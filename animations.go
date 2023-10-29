package main

import (
	"github.com/gotk3/gotk3/gtk"
)

// Setup the TextView, put it in a ScrolledWindow, and add both to box.
func setupTextView(box *gtk.Box) *gtk.TextView {
	sw, _ := gtk.ScrolledWindowNew(nil, nil)
	tv, _ := gtk.TextViewNew()
	sw.Add(tv)
	box.PackStart(sw, true, true, 0)
	return tv
}

// func setupPropertyCheckboxes(tv *gtk.TextView, outer *gtk.Box, props []*BoolProperty) {
// 	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
// 	for _, prop := range props {
// 		chk, _ := gtk.CheckButtonNewWithLabel(prop.Name)
// 		// initialize the checkbox with the property's current value
// 		chk.SetActive(prop.Get())
// 		p := prop // w/o this all the checkboxes will toggle the last property in props
// 		chk.Connect("toggled", func() {
// 			p.Set(chk.GetActive())
// 		})
// 		box.PackStart(chk, true, true, 0)
// 	}
// 	outer.PackStart(box, false, false, 0)
// }

func setupAnimationsTab() *gtk.Box{
	box := boxNew(gtk.ORIENTATION_VERTICAL, 0)

	// tv := setupTextView(box)

	chkGlx, _           := gtk.CheckButtonNewWithLabel("Use GLX:")
	chkVsync, _         := gtk.CheckButtonNewWithLabel("Enable VSync:")
	chkAnimations, _    := gtk.CheckButtonNewWithLabel("Enable Animations:")
	chkFading, _        := gtk.CheckButtonNewWithLabel("Enable Fading:")
	chkNextTagFading, _ := gtk.CheckButtonNewWithLabel("Next Tag Fading:")
	chkPrevTagFading, _ := gtk.CheckButtonNewWithLabel("Prev Tag Fading:")

	box.Add(chkGlx)
	box.Add(chkVsync)
	box.Add(chkAnimations)
	box.Add(chkFading)
	box.Add(chkNextTagFading)
	box.Add(chkPrevTagFading)

	return box
}
