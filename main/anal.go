package main

import (
	"fmt"
	"github.com/kamichidu/go-jclass"
	"os"
	"time"
)

func main() {
	filename := os.Args[1]
	fmt.Printf("file: %s\n", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v\n", err)
		os.Exit(1)
	}

    startTime := time.Now()
	jclass, err := jclass.NewJClass(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v\n", err)
		os.Exit(1)
	}
    endTime := time.Now()

    fmt.Printf("Success: %d [ns]\n", endTime.Sub(startTime).Nanoseconds())
	fmt.Printf("Name: %s\n", jclass.GetThisClass())
	fmt.Printf("Superclass: %s\n", jclass.GetSuperclass())
	fmt.Println("Implemented interfaces:")
	for _, name := range jclass.GetInterfaces() {
		fmt.Printf("\t- %s\n", name)
	}
	fmt.Println("Fields:")
	for _, jfield := range jclass.GetFields() {
		fmt.Printf("\t- %s / %s\n", jfield.GetName(), jfield.GetDescriptor())
		if len(jfield.GetAttributes()) > 0 {
			fmt.Print("\t\t")
			for _, jattribute := range jfield.GetAttributes() {
				fmt.Printf("%s, ", jattribute.GetName())
			}
			fmt.Println()
		}
	}
	fmt.Println("Methods:")
	for _, jmethod := range jclass.GetMethods() {
		fmt.Printf("\t- %s / %s\n", jmethod.GetName(), jmethod.GetDescriptor())
		if len(jmethod.GetAttributes()) > 0 {
			fmt.Print("\t\t")
			for _, jattribute := range jmethod.GetAttributes() {
				fmt.Printf("%s, ", jattribute.GetName())
			}
			fmt.Println()
		}
	}
	fmt.Println("Attributes:")
	for _, jattribute := range jclass.GetAttributes() {
		fmt.Printf("\t- %s\n", jattribute.GetName())
	}
	fmt.Println()
}
