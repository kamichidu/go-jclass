package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	ifilename = flag.String("f", "-", "Input filename to be read")
)

func main() {
	flag.Parse()

	var in io.Reader
	if *ifilename == "-" {
		in = os.Stdin
	} else {
		f, err := os.Open(*ifilename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		in = f
	}

	// address | bytes | char
	reader := bufio.NewReader(in)
    bytes := make([]byte, 80)
	for {
		n, err := reader.Read(bytes)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%02x", bytes[i])
			if i%4 == 3 {
                fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
