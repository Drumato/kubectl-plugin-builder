module {{ .PackageName }}

go {{ .GoVersion }}

require (
    {{- with .Requires }}
    {{- range . }}
    {{ .Name }} v{{ .Version }}
    {{- end }}
    {{- end }}
)
