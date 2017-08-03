package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	flag.Parse()

	srcFile := flag.Args()
	if len(srcFile) <= 0 {
		fmt.Println("Missing original PDF file")
		os.Exit(1)
	}

	if _, err := os.Stat(srcFile[0]); os.IsNotExist(err) {
		fmt.Printf("Invalid file %s\n", srcFile[0])
		os.Exit(1)
	} else {
		fmt.Println("STart!!")
	}

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	var density float64 = 300
	var quality uint = 100
	err := mw.SetResolution(density, density)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	err = mw.ReadImage(srcFile[0])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	err = mw.SetCompressionQuality(quality)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	mw.SetIteratorIndex(0)

	err = mw.SetImageFormat("jpg")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	err = mw.WriteImage("./test.jpg")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Success!")
		os.Exit(0)
	}

}
