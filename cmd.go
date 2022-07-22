package testgorepo

import (
	"fmt"
	"os/exec"
)

func _11() {
	cmd := exec.Command("ls", "-alFh", "/etc")
	out, _ := cmd.CombinedOutput()
	fmt.Printf("combined out:\n%s\n", string(out))
}
