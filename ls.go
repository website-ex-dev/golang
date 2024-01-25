package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

type Color string

const (
	Reset = "\033[0m"

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BlackLight   = "\033[90m"
	RedLight     = "\033[91m"
	GreenLight   = "\033[92m"
	YellowLight  = "\033[93m"
	BlueLight    = "\033[94m"
	MagentaLight = "\033[95m"
	CyanLight    = "\033[96m"
	WhiteLight   = "\033[97m"
)

func printFiles(full bool) {
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s total: %d\n", GreenLight, len(files))

	for _, file := range files {
		var dirChar, dirColor = "-", CyanLight
		if file.IsDir() {
			dirChar = "d"
			dirColor = BlueLight
		}

		if full {

			usr, _ := user.Current()
			dateTime := strings.Fields(file.ModTime().Format("2006 Jan 01 15:04:05"))

			fmt.Printf("%s%s%s%s %s%s %s%10s %s%s %s%s %s%s %s%s ",
				dirColor, dirChar,
				YellowLight, file.Mode().Perm(),
				BlackLight, usr.Username,
				White, strconv.FormatInt(file.Size(), 10),
				WhiteLight, dateTime[0],
				GreenLight, dateTime[1],
				GreenLight, dateTime[2],
				WhiteLight, dateTime[3],
			)
		}
		fmt.Printf("%s%s\n", dirColor, file.Name())

	}
}

func only_for_windows() {
	stdout := syscall.Handle(os.Stdout.Fd())

	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= 0x0004

	syscall.MustLoadDLL("kernel32").MustFindProc("SetConsoleMode").Call(uintptr(stdout), uintptr(originalMode))
}

func main() {
	only_for_windows()

	key := flag.Bool("l", false, "use a long listing format")
	flag.Parse()

	printFiles(*key)
}
