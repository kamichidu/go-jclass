package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/kamichidu/go-jclass"
	"github.com/kamichidu/go-jclass/jvms"
)

var debug string

//go:generate esc -o bindata.go -ignore \.go$ -private .
var javapTemplate = template.New("javap")

func init() {
	useLocal := debug != ""
	template.Must(javapTemplate.
		Funcs(map[string]interface{}{
			"modifiers": modifiers,
			"typeKind":  typeKind,
		}).
		Parse(_escFSMustString(useLocal, "/javap.tpl")))
}

func modifiers(flags jclass.AccessFlags) string {
	mods := make([]string, 0)
	switch {
	case flags.IsPublic():
		mods = append(mods, "public")
	case flags.IsProtected():
		mods = append(mods, "protected")
	case flags.IsPrivate():
		mods = append(mods, "private")
	}
	if flags.IsStatic() {
		mods = append(mods, "static")
	}
	if flags.IsFinal() {
		mods = append(mods, "final")
	} else if flags.IsAbstract() {
		mods = append(mods, "abstract")
	}
	if len(mods) > 0 {
		mods = append(mods, "")
	}
	return strings.Join(mods, " ")
}

func typeKind(class *jclass.JavaClass) string {
	switch {
	case class.IsInterface():
		return "interface"
	case class.IsEnum():
		return "enum"
	case class.IsAnnotation():
		return "@interface"
	default:
		return "class"
	}
}

func writeFormat(w io.Writer, class *jclass.JavaClass) error {
	fmt.Fprintf(w, "Compiled from \"%s\"\n", class.SourceFile())
	fmt.Fprintf(w, "%s %s %s", modifiers(class), typeKind(class), class.CanonicalName())
	if class.SuperClass() != "" {
		fmt.Fprintf(w, " extends %s", class.SuperClass())
	}
	if len(class.Interfaces()) > 0 {
		fmt.Fprintf(w, " implements %s", strings.Join(class.Interfaces(), ", "))
	}
	fmt.Fprint(w, " {\n")
	for _, field := range class.Fields() {
		fmt.Fprintf(w, "  %s %s %s;\n", modifiers(field), field.Type(), field.Name())
	}
	for _, method := range class.Methods() {
		mod := modifiers(method)
		fmt.Fprint(w, "  ")
		if mod != "" {
			fmt.Fprintf(w, "%s ", mod)
		}
		fmt.Fprintf(w, "%s %s(%s)\n", method.ReturnType(), method.Name(), strings.Join(method.ParameterTypes(), ", "))
	}
	fmt.Fprint(w, "}\n")

	return nil
}

func run() int {
	flag.Parse()

	for _, classOrFilename := range flag.Args() {
		file, err := os.Open(classOrFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't open file: %s\n", err)
			continue
		}
		defer file.Close()

		cf, err := jvms.ParseClassFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't parse class file: %s\n", err)
			continue
		}

		class := jclass.NewJavaClass(cf)
		// if err = writeFormat(os.Stdout, class); err != nil {
		if err = javapTemplate.Execute(os.Stdout, class); err != nil {
			fmt.Fprintf(os.Stderr, "Write result error: %s\n", err)
		}
	}

	return 0
}

func main() {
	os.Exit(run())
}
