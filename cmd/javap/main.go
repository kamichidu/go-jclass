package main

import (
	"flag"
	"fmt"
	"github.com/kamichidu/go-jclass"
	"github.com/kamichidu/go-jclass/encoding/jvms"
	"os"
	"strings"
)

func run() int {
	flag.Parse()

	for _, classOrFilename := range flag.Args() {
		file, err := os.Open(classOrFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't open file: %s", err)
			continue
		}
		defer file.Close()

		cf, err := jvms.ParseClassFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't parse class file: %s", err)
			continue
		}

		// TODO: javap compatible output
		class := jclass.NewJavaClass(cf)
		fmt.Printf("Canonical Name = %s\n", class.CanonicalName())
		fmt.Printf("public(%v), protected(%v), private(%v), final(%v)\n", class.IsPublic(), class.IsProtected(), class.IsPrivate(), class.IsFinal())
		fmt.Printf("extends %s\n", class.SuperClass())
		fmt.Printf("implements %s\n", strings.Join(class.Interfaces(), ", "))
	}

	return 0
}

func main() {
	os.Exit(run())
}
