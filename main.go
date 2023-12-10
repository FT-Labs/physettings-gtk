package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	p "physettings/picom"
)

const (
	width     = 800
	height    = 600
	spacing	  = 25
	logo_path = "/usr/share/pixmaps/phyOS-logo-128x128.png"
)

func imageNew(path string) *gtk.Image {

	img, err := gtk.ImageNewFromFile(path)

	if err != nil {
		log.Fatal("Error: can not load image file", err)
	}

	return img
}

func labelNew(text string) *gtk.Label {

	label, err := gtk.LabelNew(text)

	if err != nil {
		log.Fatal("Error: Can not create label", text)
	}

	return label
}

func windowNew(title string) *gtk.Window {

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal("Error: Can not create window", err)
	}

	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetSizeRequest(width, height)
	win.SetDefaultSize(width, height)
	win.SetPosition(gtk.WIN_POS_CENTER)

	return win
}


func main() {

	gtk.Init(nil)

	nb, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal("Unable to create notebook:", err)
	}
	win := windowNew("phy")
	win.Add(nb)
		nb.SetHExpand(true)
	nb.SetVExpand(true)

	logo := imageNew(logo_path)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 15)
	logo.SetHAlign(gtk.ALIGN_CENTER)
	box.SetMarginStart(spacing)
	box.SetMarginTop(spacing)
	box.Add(logo)


	// Add a child widget and tab label to the notebook so it renders.
	nbChild, err := gtk.LabelNew("Notebook content")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	nbInfo := labelNew("INFO")
	nbOptions := labelNew("OPTIONS")
	nbAnimations := labelNew("ANIMATIONS")
	nb.SetTabPos(gtk.POS_BOTTOM)
	nb.AppendPage(box, nbInfo)
	nb.AppendPage(nbChild, nbOptions)

	animationsBox := p.SetupAnimationsTab()
	nb.AppendPage(animationsBox, nbAnimations)

	win.ShowAll()

	gtk.Main()
}
