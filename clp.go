/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

clp, copy the content of a file into your clipboard.
*/
package main

import (
	"fmt"
	"os"

	"github.com/Alvesafk/scolor"
	"github.com/Alvesafk/scolor/ansi"
	"github.com/atotto/clipboard"
)

// func colorError receives nothing and returns nothing, it's a side effect function to print
// the basic error message depending if the user terminal has access to the True colors or
// not.
func colorError() {
	if scolor.IsRGBSupported {
		scolor.RGB(180, 60, 60).BgPrintln("Error:")
	} else {
		ansi.ARed.BgPrintln("Error:")
	}
}

// All the logic is on the main function, is a simple program so it's not much, first get
// the arguments, check if argumenst are different from 2, if not read the file, transforms
// []byte that was returned out of the file into string and copy onto the clipboard.
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

		fmt.Printf("Was not possible to read %s file.\n", args[1])
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

/*
INDEX:
func colorError()
func main()
*/
