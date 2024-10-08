package {{.Package}}

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"{{.API}}"
)

{{define "case"}}
	{{- $combo := .Combo}}
	{{- $prefix := .Prefix}}
	{{- $idx := 0}}

	{{- range $typ, $def := .Types}}
		{{- if gt $idx 0}} &&{{else}}{{$idx = 1}}{{end}} {{$def.VarName}} {{if contains $combo $typ}}!={{else}}=={{end}} nil
	{{- end}}:
		return &struct {
			{{.BaseType}}
{{- range $typ, $def := .Types}}
	{{- if contains $combo $typ}}
			{{$typ}}
	{{- end}}
{{- end}}
		}{
			{{.ShortBase}}: base,
{{- range $typ, $def := .Types}}
	{{- if contains $combo $typ}}
			{{$def.ShortType}}: &{{$prefix}}{{$def.ShortType}}Impl{
				{{$def.VarName}}: {{$def.VarName}},
			},
	{{- end}}
{{- end}}
		}
{{- end -}}

func {{.Function}}(base {{.BaseType}}{{range ordered}}, {{.VarName}} {{.Signature}}{{end}}) {{.ReturnType}} {
{{- $basetype := .BaseType}}
{{- $shortbase := .ShortBase}}
{{- $prefix := .Function}}
{{- $types := .Types}}
{{- $idx := 0}}
	switch {
	case {{- range $typ, $def := .Types}}
		{{- if gt $idx 0}} &&{{else}}{{$idx = 1}}{{end}} {{$def.VarName}} == nil
	{{- end}}:
		return base
{{range $combo := .Combinations}}
	case {{- template "case" dict "BaseType" $basetype "Prefix" $prefix "ShortBase" $shortbase "Types" $types "Combo" $combo}}
{{end}}	}

	return nil
}

{{range .Types -}}
type {{$prefix}}{{.ShortType}}Impl struct {
	{{.VarName}} {{.Signature}}
}

func (impl *{{$prefix}}{{.ShortType}}Impl) {{.Function}}(
	{{- range $idx, $param := .Params -}}
		{{- if gt $idx 0}}, {{end -}}
		p{{$idx}} {{ $param -}} 
	{{end}}){{ .ReturnTypes }} {
	return impl.{{.VarName}}(
	{{- range $idx, $param := .Params -}}
		{{- if gt $idx 0}}, {{end}}p{{- $idx -}}{{end}})
	}

{{end}}
