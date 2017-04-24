{{$class := . -}}
Compiled from "{{$class.SourceFile}}"
{{modifiers $class}}{{typeKind $class}} {{$class.SimpleName}} {
{{- range $field := $class.Fields}}
  {{modifiers $field}}{{$field.Type}} {{$field.Name}};
{{- end -}}
{{range $ctor := $class.Constructors}}
  {{modifiers $ctor}}{{$class.CanonicalName}}({{range $i, $p := $ctor.ParameterTypes}}{{if gt $i 0}}, {{end}}{{$p}}{{end}});
{{- end -}}
{{range $method := $class.Methods}}
  {{modifiers $method}}{{$method.ReturnType}} {{$method.Name}}({{range $i, $p := $method.ParameterTypes}}{{if gt $i 0}}, {{end}}{{$p}}{{end}});
{{- end}}
}
