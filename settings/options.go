package settings

import (
	"fmt"
	"log"
	"os/exec"
	u "github.com/FT-Labs/physettings-gtk/utils"

	"github.com/gotk3/gotk3/gtk"
)

func setupWallpaperBox() *gtk.Box {
	wallpaperLabel, _ := gtk.LabelNew("Set Wallpaper:    ")
	wallpaperOpen, _ := gtk.FileChooserButtonNew("Set Wallpaper", gtk.FILE_CHOOSER_ACTION_OPEN)

	filter, _ := gtk.FileFilterNew()
	filter.AddPattern("*.jpg")
	filter.AddPattern("*.jpeg")
	filter.AddPattern("*.png")
	filter.SetName("*.jpg, *.jpeg, *.png")
	wallpaperOpen.AddFilter(filter)


	wallpaperOpen.Connect("selection-changed", func(){
		cmd := fmt.Sprintf("pOS-setbg %s", wallpaperOpen.GetFilename())
		u.RunCommand(cmd)
	})


	wallpaperBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	wallpaperBox.Add(wallpaperLabel)
	wallpaperBox.Add(wallpaperOpen)
	return wallpaperBox
}

func SetupOptionsTab() *gtk.Box {
	u.FetchAttributes()
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box.SetMarginStart(20)
	box.SetMarginEnd(20)
	box.SetMarginTop(20)
	boxLeft, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	boxRight, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	boxLeft.SetHExpand(true)

	wbox := setupWallpaperBox()

	cbRofiColor := comboBoxNewWithLabel("Rofi Colorsheme:  ", gtk.ORIENTATION_HORIZONTAL)
	cbPowermenuType := comboBoxNewWithLabel("Powermenu Type:", gtk.ORIENTATION_HORIZONTAL)
	cbPowermenuStyle := comboBoxNewWithLabel("Powermenu Style:", gtk.ORIENTATION_HORIZONTAL)
	comboBoxAddEntries(u.RofiColors, cbRofiColor.ComboBox)
	comboBoxAddEntries(u.PowerMenuTypes, cbPowermenuType.ComboBox)
	comboBoxAddEntries(u.PowerMenuStyles, cbPowermenuStyle.ComboBox)

	cbRofiColor.ComboBox.Connect("changed", func() {
		u.SetRofiColor(cbRofiColor.ComboBox.GetActiveText())
	})

	cbPowermenuType.ComboBox.Connect("changed", func() {
		u.SetAttribute(u.POWERMENU_TYPE, cbPowermenuType.ComboBox.GetActiveText())
	})

	cbPowermenuStyle.ComboBox.Connect("changed", func() {
		u.SetAttribute(u.POWERMENU_STYLE, cbPowermenuStyle.ComboBox.GetActiveText())
	})

	chkAsk, _ := gtk.CheckButtonNewWithLabel("Confirm on Shutdown")

	if u.Attrs[u.POWERMENU_CONFIRM] == "true" {
		chkAsk.SetActive(true)
	}

	chkAsk.Connect("clicked", func() {
		if chkAsk.GetActive() {
			u.SetAttribute(u.POWERMENU_CONFIRM, "true")
		} else {
			u.SetAttribute(u.POWERMENU_CONFIRM, "false")
		}
	})

	boxLeft.Add(wbox)
	boxLeft.Add(cbRofiColor.Box)
	boxLeft.Add(cbPowermenuType.Box)
	boxLeft.Add(cbPowermenuStyle.Box)
	boxLeft.Add(chkAsk)

	pbGrub, _ := gtk.ButtonNewWithLabel("Select Grub Theme")
	pbSddm, _ := gtk.ButtonNewWithLabel("Select SDDM Theme")
	pbBar, _ := gtk.ButtonNewWithLabel("Select Statusbar Widgets")
	pbControl, _ := gtk.ButtonNewWithLabel("Open pdwm Control Center")

	pbGrub.Connect("clicked", func() {
		err := u.RunScript(u.POS_GRUB_CHOOSE_THEME)
		if err != nil {
			log.Printf("Grub script failed")
		}
	})

	pbSddm.Connect("clicked", func() {
		err := u.RunScript(u.POS_SDDM_CHOOSE_THEME)
		if err != nil {
			log.Printf("Sddm script failed")
		}
	})

	pbBar.Connect("clicked", func() {
		err := u.RunScript(u.POS_MAKE_BAR)
		if err != nil {
			log.Printf("Bar script failed")
		}
	})

	pbControl.Connect("clicked", func() {
		exec.Command("pdwmc", "-q").Run()
	})

	boxRight.Add(pbGrub)
	boxRight.Add(pbSddm)
	boxRight.Add(pbBar)
	boxRight.Add(pbControl)

	box.Add(boxLeft)
	box.Add(boxRight)

	return box
}
