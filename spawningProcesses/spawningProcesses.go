// Sometimes our Go programs need to spawn other, non-Go
// processes. For example, the syntax highlighting on this
// site is [implemented](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)
// by spawning a [`pygmentize`](http://pygments.org/)
// process from a Go program. Let's look at a few examples
// of spawning processes from Go.

package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("java")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
