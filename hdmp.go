package main

// ----------------------------------------------------------------------------
// hdmp.go
// Simple hex dump utility
// ----------------------------------------------------------------------------
// Copyright (c) 2024 Robert L. Snyder <rob@mooneyedkitty.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
// ----------------------------------------------------------------------------

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

func byteToASCII(b byte) string {
	if b >= 32 && b <= 126 {
		return string(b)
	}
	return "."
}

func main() {

	var fileName string
	var lineLength int
	var grouping int
	var origin int

	flag.StringVar(&fileName, "f", "", "The file name to dump")
	flag.IntVar(&lineLength, "l", 16, "Line length (default 16)")
	flag.IntVar(&grouping, "g", 4, "Byte grouping (default 4)")
	flag.IntVar(&origin, "o", 0, "Origina address (default 0)")
	flag.Parse()

	if fileName == "" {
		fmt.Println("-f (file name) must be specified.")
		flag.Usage()
		os.Exit(1)
	}

	data, err := readFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for i := 0; i < len(data); i += lineLength {
		fmt.Printf("%04x\t", origin)
		text := ""
		for n := range lineLength {
			if i+n >= len(data) {
				break
			}
			text += byteToASCII(data[i+n])
			fmt.Printf("%02x ", data[i+n])
			if (n+1)%grouping == 0 && n != (lineLength-1) {
				fmt.Print(" | ")
			}
		}
		fmt.Printf("\t|%s|", text)
		fmt.Println()
		origin += lineLength
	}

}
