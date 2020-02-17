package main

import (
	binarybytes "builder/pkg/concrete_builder/binary_bytes"
	"fmt"
	"os"

	binaryrunes "builder/pkg/concrete_builder/binary_runes"
	"builder/pkg/director"
)

func main() {
	directorService := director.NewEncoderService()
	binaryRunesEncoder := binaryrunes.NewBuilder("世")
	message, err := directorService.BuildMessage(binaryRunesEncoder)
	if err != nil {
		fmt.Println("error building message: %w", err)
		os.Exit(1)
	}
	fmt.Printf("Original: %s\n"+
		"Encoded: %s\n"+
		"Encoding type: %s\n"+
		"Symbol Type: %s\n", message.GetOriginal(), message.Get(), message.Encoding(), message.Symbol())

	binaryBytesEncoder := binarybytes.NewBuilder("世")
	message, err = directorService.BuildMessage(binaryBytesEncoder)
	if err != nil {
		fmt.Println("error building message: %w", err)
		os.Exit(1)
	}
	fmt.Printf("Original: %s\n"+
		"Encoded: %s\n"+
		"Encoding type: %s\n"+
		"Symbol Type: %s\n", message.GetOriginal(), message.Get(), message.Encoding(), message.Symbol())
}
