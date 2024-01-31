package clearscreen

import (
	"os"
	"os/exec"
	"runtime"
)

var clearMap map[string]func() //create a map for storing clear funcs

func init() {
	clearMap = make(map[string]func()) //Initialize it
	clearMap["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearMap["darwin"] = clearMap["linux"] // MacOS uses the same 'clear' command as Linux
	clearMap["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clearMap["aix"] = clearMap["linux"]
	clearMap["android"] = clearMap["linux"]
	clearMap["dragonfly"] = clearMap["linux"]
	clearMap["freebsd"] = clearMap["linux"]
	clearMap["illumos"] = clearMap["linux"]
	clearMap["netbsd"] = clearMap["linux"]
	clearMap["openbsd"] = clearMap["linux"]
	clearMap["solaris"] = clearMap["linux"]
	clearMap["zos"] = clearMap["linux"]
}

func Exec() {
	value, ok := clearMap[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                             //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
