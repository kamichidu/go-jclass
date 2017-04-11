package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/kamichidu/go-jclass"
	"github.com/kamichidu/go-jclass/encoding/jvms"
)

const javapTemplateText = `Compiled from "???"
{{typeNamePrefix .}} {{.CanonicalName}}{{superclass .}}{{superinterfaces .}} {
{{range $field := .Fields}}  {{$field.Type}} {{$field.Name}};
{{end}}
{{range $method := .Methods}}  {{methodPrefix $method}}{{$method.ReturnType}} {{$method.Name}}({{join $method.ParameterTypes ", "}});
{{end}}
}
`

var javapTemplate *template.Template

func init() {
	funcs := make(template.FuncMap)
	funcs["methodPrefix"] = func(flags jclass.AccessFlags) string {
		mods := make([]string, 0)
		if flags.IsPublic() {
			mods = append(mods, "public")
		} else if flags.IsProtected() {
			mods = append(mods, "protected")
		} else if flags.IsPrivate() {
			mods = append(mods, "private")
		}
		if flags.IsStatic() {
			mods = append(mods, "static")
		}
		if flags.IsFinal() {
			mods = append(mods, "final")
		}
		if len(mods) > 0 {
			return strings.Join(mods, " ") + " "
		} else {
			return strings.Join(mods, " ")
		}
	}
	funcs["typeNamePrefix"] = func(class *jclass.JavaClass) string {
		mods := make([]string, 0)
		if class.IsPublic() {
			mods = append(mods, "public")
		} else if class.IsProtected() {
			mods = append(mods, "protected")
		} else if class.IsPrivate() {
			mods = append(mods, "private")
		}
		if class.IsStatic() {
			mods = append(mods, "static")
		}
		if class.IsFinal() {
			mods = append(mods, "final")
		}
		if class.IsInterface() {
			mods = append(mods, "interface")
		} else if class.IsEnum() {
			mods = append(mods, "enum")
		} else if class.IsAnnotation() {
			mods = append(mods, "@interface")
		} else {
			mods = append(mods, "class")
		}
		return strings.Join(mods, " ")
	}
	funcs["superclass"] = func(class *jclass.JavaClass) string {
		if class.SuperClass() == "java.lang.Object" {
			return ""
		} else {
			return " extends " + class.SuperClass()
		}
	}
	funcs["superinterfaces"] = func(class *jclass.JavaClass) string {
		if class.InterfacesCount == 0 {
			return ""
		}
		implements := make([]string, 0)
		for _, interface_ := range class.Interfaces() {
			implements = append(implements, interface_)
		}
		return " implements " + strings.Join(implements, ", ")
	}
	funcs["join"] = strings.Join
	javapTemplate = template.Must(template.New("javap").Funcs(funcs).Parse(javapTemplateText))
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

		// TODO: javap compatible output
		class := jclass.NewJavaClass(cf)
		if err = javapTemplate.Execute(os.Stdout, class); err != nil {
			fmt.Fprintf(os.Stderr, "Write result error: %s\n", err)
		}
	}

	return 0
}

func main() {
	os.Exit(run())
}
