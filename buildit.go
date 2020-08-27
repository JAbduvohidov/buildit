package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var (
	forF  = flag.String("for", "me", "-for linux\n-for windows\n-for me")
	nameF = flag.String("name", "", "-name example")
)

const (
	WINDOWS = "windows"
	LINUX   = "linux"
	CURRENT = "me"
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

	cmd := exec.Command("powershell", fmt.Sprintf(`$env:GOOS="%s" ; go build %s`, platform, filenameString))
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}
