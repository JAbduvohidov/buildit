package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	WINDOWS = "windows"
	LINUX   = "linux"
	CURRENT = "me"

	MAJOR = "major"
	MINOR = "minor"
	PATCH = "patch"

	newLine = "\n"

	filePath = "./version/version.go"
)

var (
	forF     = flag.String("for", "me", "-for linux\n-for windows\n-for me")
	nameF    = flag.String("name", "", "-name example")
	versionF = flag.String("v", "", "-v major\n-v feat\n-v fix")

	versionsTypes = map[string]string{
		MAJOR:  MAJOR,
		"feat": MINOR,
		"fix":  PATCH,
	}
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

	version := strings.TrimSpace(*versionF)

	switch version {
	case MAJOR:
		fallthrough
	case "feat":
		fallthrough
	case "fix":
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalln("unable to read data from file", err)
		}
		err = os.Remove(filePath)
		if err != nil {
			log.Fatalln("unable to remove file", err)
		}

		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalln("unable to create file", err)
		}

		lines := strings.Split(string(data), newLine)
		for index, line := range lines {
			if i := strings.Index(line, versionsTypes[version]+" = "); i != -1 {
				versionPart, err := strconv.Atoi(line[i+len(versionsTypes[version]+" = "):])
				if err != nil {
					log.Fatalln("invalid file syntax")
					return
				}
				line = strings.Replace(line, strconv.Itoa(versionPart), strconv.Itoa(versionPart+1), 1)
			}

			if index == len(lines)-1 {
				_, err := file.WriteString(line)
				if err != nil {
					log.Fatalln("unable to write,", err)
				}

				continue
			}

			_, err := file.WriteString(line + newLine)
			if err != nil {
				log.Fatalln("unable to write,", err)
			}

		}

		err = file.Close()
		if err != nil {
			log.Fatalln("unable to close file", err)
		}
	default:
		log.Fatalf("version type '%s' not found", *versionF)
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
