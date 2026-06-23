package main

import (
	"fmt"
	"os"

	"github.com/Alvesafk/scolor"
	"github.com/Alvesafk/scolor/ansi"
	"github.com/atotto/clipboard"
)

func colorError() {
	if scolor.IsRGBSupported {
		scolor.RGB(180, 60, 60).BgPrintln("Error:")
	} else {
		ansi.ARed.BgPrintln("Error:")
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		colorError()

		fmt.Printf("%s usage instructions:\n%s <File-to-clipboard>\n", args[0], args[0])
		return
	}

	fileContent, err := os.ReadFile(string(args[1]))
	if err != nil {
		colorError()

		fmt.Printf("Was not possible to read %s file.", args[1])
		return
	}

	fileContentString := string(fileContent)

	err = clipboard.WriteAll(fileContentString)
	if err != nil {
		colorError()

		fmt.Print(err, '\n')

		return
	}

	if scolor.IsRGBSupported {
		successTmpl := scolor.CreateRgbTemplate(scolor.RGB(80, 180, 80), scolor.BLACK)
		successTmpl.Println("Success, the content is in your clipboard")
	} else {
		ansi.AGreen.BgPrintln("Success, the content is in your clipboard.")
	}
}
