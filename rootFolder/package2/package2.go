package package2

import (
	"fmt"
	"os/exec"
)

func Printpack2() {
	fmt.Print("Hello from package 2")
	cmd_apikey := exec.Command("yq", "--version")
	apikeyOut, _ := cmd_apikey.CombinedOutput()
	fmt.Println("running apikey mask command :", string(apikeyOut))
}
