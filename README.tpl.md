# CLI version in the wild

This document describe what approach was implemented by some popular CLI to display its version.

Tested CLIs:
{{- range .CLIs }}
- [{{ .Name }}]({{ .URL }})
{{- end }}

## How to get the version
<!-- {{ add (len .ViaCmdOnly) (len .ViaFlagsOnly) (len .ViaFlagsAndCmd) }} == {{ len .CLIs }} -->

### via command (only)

{{ len .ViaCmdOnly }} of {{ len .CLIs }}

{{- range .ViaCmdOnly }}
- [{{ .Name }}]({{ .URL }})
{{- end }}

### via flag (only)

{{ len .ViaFlagsOnly }} of {{ len .CLIs }}

| CLI Name | Flags |
|----------|-------|
{{- range .ViaFlagsOnly }}
| [{{ .Name }}]({{ .URL }}) | `{{ .Analysis.Flags | sortAlpha | join "`, `" }}` |
{{- end }}


### via flag and command (both)

{{ len .ViaFlagsAndCmd }} of {{ len .CLIs }}

| CLI Name | Flags |
|----------|-------|
{{- range .ViaFlagsAndCmd }}
| [{{ .Name }}]({{ .URL }}) | `{{ .Analysis.Flags | sortAlpha | join "`, `" }}` |
{{- end }}

### Overall flag names popularity

| Name | Number of usage |
|-----|---------|
{{- range $key, $value := .FlagsNames }}
| `{{ $key }}` | {{ $value }} |
{{- end }}

## Upgrade notice

## What data is collected

## Output

{{- range $key, $val := .Output }}

### {{ $key }}

{{- range $item := $val }}
{{ $myDict := dict "json" "json" "yaml" "yaml" "plain" "text" "short" "text" }}
- `{{ $item.Cmd }}`
  ```{{ get $myDict $key }}
{{ $item.Out | indent 2}}```
{{- end }}
{{- end }}

----

### JSON

### YAML

### Go template

### Short

1 of 13

- `kubectl version --short` (deprecated, will be default in the future)
  ```
  Client Version: v1.24.0
  Kustomize Version: v4.5.4
  ```

## Frameworks built-in support

### Cobra

```go
cmd := &cobra.Command{
  Use:   "example",
  Short: "An example CLI built with github.com/spf13/cobra",

  // Version defines the version for this command. If this value is non-empty and the command does not
  // define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
  // will print content of the "Version" variable. A shorthand "v" flag will also be added if the
  // command does not define one.
  Version: "1.0.3",
}
```

###

Interesting approach to print the version of deps: https://github.com/hashicorp/terraform/blob/v1.3.0-dev/version/dependencies.go
