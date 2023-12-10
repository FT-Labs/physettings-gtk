package picom

import (
	"log"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

const (
	spaceSize = 10
)

type ComboBoxLabel struct {
	Box		 *gtk.Box
	ComboBox *gtk.ComboBoxText
}

type SpinButtonLabel struct {
	Box        *gtk.Box
	SpinButton *gtk.SpinButton
}

// Setup the TextView, put it in a ScrolledWindow, and add both to box.
func setupTextView(box *gtk.Box) *gtk.TextView {
	sw, _ := gtk.ScrolledWindowNew(nil, nil)
	tv, _ := gtk.TextViewNew()
	sw.Add(tv)
	box.PackStart(sw, true, true, 0)
	return tv
}

func comboBoxAddEntries(s []string, cbx *gtk.ComboBoxText) {
	for _, val := range s {
		cbx.AppendText(val)
	}
	cbx.SetActive(0)
}

func comboBoxNewWithLabel(s string, o gtk.Orientation) *ComboBoxLabel {
	lbl, _ := gtk.LabelNew(s)
	box, _ := gtk.BoxNew(o, spaceSize)
	cbx, _ := gtk.ComboBoxTextNew()

	box.Add(lbl)
	box.Add(cbx)

	return &ComboBoxLabel{box, cbx}
}

func spinButtonNewWithLabel(s string, o gtk.Orientation, mn, mx, stp float64) *SpinButtonLabel {
	lbl, _ := gtk.LabelNew(s)
	box, _ := gtk.BoxNew(o, spaceSize)
	spn, _ := gtk.SpinButtonNewWithRange(mn, mx, stp)
	box.Add(lbl)
	box.Add(spn)
	return &SpinButtonLabel{box, spn}
}

func SetupAnimationsTab() *gtk.Box{
	readPicomOpts()
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	boxLeft, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	boxRight, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	boxLeft.SetHExpand(true)

	chkAnimations, _    := gtk.CheckButtonNewWithLabel(": Enable Animations")
	chkFading, _        := gtk.CheckButtonNewWithLabel(": Enable Fading")
	chkNextTagFading, _ := gtk.CheckButtonNewWithLabel(": Next Tag Fading")
	chkPrevTagFading, _ := gtk.CheckButtonNewWithLabel(": Prev Tag Fading")
	chkGlx, _           := gtk.CheckButtonNewWithLabel(": Use GLX")
	chkVsync, _         := gtk.CheckButtonNewWithLabel(": Enable VSync")
	chkShadow, _		:= gtk.CheckButtonNewWithLabel(": Enable Shadows")

	isChecked := func(chk *gtk.CheckButton, s string) {
		if picomOpts[s] == "false" {
			return
		}
		chk.SetActive(true)
	}
	isChecked(chkAnimations, _animations)
	isChecked(chkFading, _fading)
	isChecked(chkNextTagFading, _enable_fading_next_tag)
	isChecked(chkPrevTagFading, _enable_fading_prev_tag)
	isChecked(chkVsync, _vsync)
	isChecked(chkGlx, _backend)
	isChecked(chkShadow, _shadow)

	spnAnimSpeedInTag := spinButtonNewWithLabel("Anim Speed in Tag:                    ",
												gtk.ORIENTATION_HORIZONTAL,
												30, 250, 1)
	spnAnimSpeedInTag.SpinButton.SetValue(stof(picomOpts[_animation_stiffness_in_tag]))
	spnAnimSpeedOnTagChange := spinButtonNewWithLabel("Anim Speed on Tag Change: ",
												gtk.ORIENTATION_HORIZONTAL,
												30, 250, 1)
	spnAnimSpeedOnTagChange.SpinButton.SetValue(stof(picomOpts[_animation_stiffness_tag_change]))

	cmbOpenWindowAnim := comboBoxNewWithLabel("Open Window Anim: ", gtk.ORIENTATION_HORIZONTAL)
	comboBoxAddEntries(animOpenOpts, cmbOpenWindowAnim.ComboBox)
	cmbCloseWindowAnim := comboBoxNewWithLabel("Close Window Anim: ", gtk.ORIENTATION_HORIZONTAL)
	comboBoxAddEntries(animCloseOpts, cmbCloseWindowAnim.ComboBox)
	cmbPrevTag := comboBoxNewWithLabel("Anim For Prev Tag:    ", gtk.ORIENTATION_HORIZONTAL)
	comboBoxAddEntries(animPrevOpts, cmbPrevTag.ComboBox)
	cmbNextTag := comboBoxNewWithLabel("Anim For Next Tag:    ", gtk.ORIENTATION_HORIZONTAL)
	comboBoxAddEntries(animNextOpts, cmbNextTag.ComboBox)

	boxLeft.Add(chkGlx)
	boxLeft.Add(chkVsync)
	boxLeft.Add(chkAnimations)
	boxLeft.Add(chkFading)
	boxLeft.Add(chkNextTagFading)
	boxLeft.Add(chkPrevTagFading)
	boxLeft.Add(chkShadow)

	boxRight.Add(spnAnimSpeedInTag.Box)
	boxRight.Add(spnAnimSpeedOnTagChange.Box)
	boxRight.Add(cmbOpenWindowAnim.Box)
	boxRight.Add(cmbCloseWindowAnim.Box)
	boxRight.Add(cmbPrevTag.Box)
	boxRight.Add(cmbNextTag.Box)

	pb, _ := gtk.ButtonNewWithLabel("Save Changes")
	pb.SetMarginStart(20)

	pb.Connect("clicked", func(){
		if chkAnimations.GetActive() {
			changePicomAttribute(_animations, "true", false)
		} else {
			changePicomAttribute(_animations, "false", false)
		}

		if chkGlx.GetActive() {
			changePicomAttribute(_backend, "glx", true)
		} else {
			changePicomAttribute(_backend, "xrender", true)
		}

		if chkVsync.GetActive() {
			changePicomAttribute(_vsync, "true", false)
		} else {
			changePicomAttribute(_vsync, "false", false)
		}

		if chkShadow.GetActive() {
			changePicomAttribute(_shadow, "true", false)
		} else {
			changePicomAttribute(_shadow, "false", false)
		}

		if chkFading.GetActive() {
			changePicomAttribute(_fading, "true", false)
		} else {
			changePicomAttribute(_fading, "false", false)
		}

		if chkNextTagFading.GetActive() {
			changePicomAttribute(_enable_fading_next_tag, "true", false)
		} else {
			changePicomAttribute(_enable_fading_next_tag, "false", false)
		}

		if chkPrevTagFading.GetActive() {
			changePicomAttribute(_enable_fading_prev_tag, "true", false)
		} else {
			changePicomAttribute(_enable_fading_prev_tag, "false", false)
		}

		picomOpts[_animation_stiffness_in_tag] = strconv.FormatFloat(spnAnimSpeedInTag.SpinButton.GetValue(), 'f', 1, 64)
		picomOpts[_animation_stiffness_tag_change] = strconv.FormatFloat(spnAnimSpeedOnTagChange.SpinButton.GetValue(), 'f', 1, 64)

		changePicomAttribute(_animation_stiffness_in_tag, picomOpts[_animation_stiffness_in_tag], false)
		changePicomAttribute(_animation_stiffness_tag_change, picomOpts[_animation_stiffness_tag_change], false)


		err := savePicomOpts()
		if err != nil {
			log.Fatal(err.Error())
		}
	})


	box.Add(boxLeft)
	box.Add(boxRight)
	box.Add(pb)

	return box
}
