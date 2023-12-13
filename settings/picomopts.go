package settings

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var picomConfPath string

const (
    _animations                     = "animations"
    _fading                         = "fading"
    _enable_fading_next_tag         = "enable-fading-next-tag"
    _enable_fading_prev_tag         = "enable-fading-prev-tag"
    _animation_stiffness_in_tag     = "animation-stiffness-in-tag"
    _animation_stiffness_tag_change = "animation-stiffness-tag-change"
    _animation_for_open_window      = "animation-for-open-window"
    _animation_for_unmap_window     = "animation-for-unmap-window"
    _animation_for_prev_tag         = "animation-for-prev-tag"
    _animation_for_next_tag         = "animation-for-next-tag"
    _vsync                          = "vsync"
	_backend						= "backend"
	_shadow							= "shadow"
	_shadow_radius					= "shadow-radius"
)

var picomOpts = map[string]string {
    _animations                     : "false",
    _fading                         : "false",
    _animation_stiffness_in_tag     : "0",
    _animation_stiffness_tag_change : "0",
    _animation_for_open_window      : "none",
    _animation_for_unmap_window     : "none",
    _animation_for_prev_tag         : "none",
    _animation_for_next_tag         : "none",
    _enable_fading_next_tag         : "false",
    _enable_fading_prev_tag         : "false",
    _vsync                          : "false",
	_shadow							: "false",
	_shadow_radius					: "0",
	_backend						: "glx",
}

var animInfo = map[string]string {
    "fly-in"           : "Windows fly in from random directions to the screen.",
    "maximize"         : "Windows pop from center of the screen to their respective positions.",
    "minimize"         : "Windows minimize from their position to the center of the screen.",
    "slide-in-center"  : "Windows move from upper-center of the screen to their respective positions.",
    "slide-out-center" : "Windows move to the upper-center of the screen.",
    "slide-left"       : "Windows are created from the right-most window position and slide leftwards.",
    "slide-right"      : "Windows are created from the left-most window position and slide rightwards.",
    "slide-down"       : "Windows are moved from the top of the screen and slide downward.",
    "slide-up"         : "Windows are moved from their position to top of the screen.",
    "squeeze"          : "Windows are either closed or created to/from their center y-position (the animation is similar to a blinking eye).",
    "squeeze-bottom"   : "Similar to squeeze, but the animation starts from bottom-most y-position.",
    "zoom"             : "Windows are either created or destroyed from/to their center (not the screen center).",
}

var animOpenOpts = []string{
    "fly-in",
    "slide-up",
    "slide-down",
    "slide-left",
    "slide-right",
    "squeeze",
    "squeeze-bottom",
    "zoom",
}

var animCloseOpts = []string{
    "slide-out-center",
    "squeeze",
    "squeeze-bottom",
    "zoom",
}

var animPrevOpts = []string{
    "minimize",
    "slide-out-center",
    "slide-down",
    "slide-up",
    "squeeze",
    "squeeze-bottom",
    "zoom",
}

var animNextOpts = []string {
    "fly-in",
    "maximize",
    "slide-in-center",
    "slide-down",
    "slide-up",
    "squeeze",
    "squeeze-bottom",
    "zoom",
}

func stof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func changePicomAttribute(attribute, value string, isString bool) error {
    picomOpts[attribute] = value
    var cmd string
    if isString {
        cmd = fmt.Sprintf("sed -i '/^%s/c\\%s = \"%s\";' /tmp/picom.conf", attribute + "\\ =", attribute, value)
    } else {
        cmd = fmt.Sprintf("sed -i '/^%s/c\\%s = %s;' /tmp/picom.conf", attribute + "\\ =", attribute, value)
    }
    err := exec.Command("/bin/bash", "-c", cmd).Run()

    return err
}

func readPicomOpts() {
    home, _ := os.UserHomeDir(); home += "/.config/picom/picom.conf"
    picomConfPath = home
    cmd := fmt.Sprintf("cp -f %s /tmp/picom.conf", home)
    exec.Command("/bin/bash", "-c", cmd).Start()
    for key := range picomOpts {
        var cmd string
		cmd = fmt.Sprintf("grep -Ew \"^%s[ =]+\" \"%s\" | cut -f1 -d \";\" | tr -d '\"\\n'", key, picomConfPath)
        out, err := exec.Command("/bin/bash", "-c", cmd).Output()
        if err != nil {
            fmt.Fprintf(os.Stderr, err.Error())
        } else {
            s := strings.ReplaceAll(string(out), " ", "")
            arr := strings.Split(s, "=")
            if len(arr) > 1 {
                picomOpts[key] = arr[1]
            }

            search := func(cur string, arr []string) {
                for i := 0; i < len(arr); i++ {
                    if arr[i] == cur {
                        arr[0], arr[i] = arr[i], arr[0]
                        break
                    }
                }
            }

            switch key {
            case _animation_for_open_window:
                search(arr[1], animOpenOpts)
            case _animation_for_unmap_window:
                search(arr[1], animCloseOpts)
            case _animation_for_next_tag:
                search(arr[1], animNextOpts)
            case _animation_for_prev_tag:
                search(arr[1], animPrevOpts)

            }
        }
    }
}

func savePicomOpts() error {
    cmd := fmt.Sprintf("cp -f /tmp/picom.conf %s", picomConfPath)
    err := exec.Command("/bin/bash", "-c", cmd).Run()

    if err != nil {
        return err
    }
    return nil
}
