package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	WINDOWS = "windows"
	LINUX   = "linux"
	CURRENT = "me"
)

var (
	forF  = flag.String("for", "me", "-for linux\n-for windows\n-for me")
	nameF = flag.String("name", "", "-name example")
)

func main() {
	flag.Parse()

	platform := CURRENT
	filename := strings.TrimSpace(*nameF)
	filenameString := ""

	switch strings.ToLower(*forF) {
	case LINUX:
		platform = LINUX
	case WINDOWS:
		platform = WINDOWS
	case CURRENT:
		platform = runtime.GOOS
	}

	if filename != "" {
		filenameString = fmt.Sprintf("-o %s", filename)
	}

	build := &exec.Cmd{
		Path:   "powershell",
		Args:   []string{"clear", fmt.Sprintf(`$env:GOOS="%s" ; go build %s`, platform, filenameString)},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := build.Run(); err != nil {
		fmt.Println("error:", err)
	}

}
