package main

import (
	"fmt"
	"os"

	"asciiart/ascii"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("provide input text and banner name")
		return
	}
	artText, err := ascii.ArtStringBuilder(args[0], "./banners/"+args[1]+".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(artText)
}
