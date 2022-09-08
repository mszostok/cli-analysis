# CLIs in the wild

> **Note**
>
> This README is automatically generated from README.tpl.md based on analyzing the data from the YAML files on the root of this directory.

This document describe what approach was implemented by some popular CLI to display its version.

Analyzed {{ len .CLIs }} CLIs:
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

| CLI Name | Flags | Command |
|----------|-------|---------|
{{- range .ViaFlagsAndCmd }}
| [{{ .Name }}]({{ .URL }}) | `{{ .Analysis.Flags | sortAlpha | join "`, `" }}` | `version` |
{{- end }}

### Overall flag names popularity

| Name | Number of usage |
|------|-----------------|
{{- range $key, $value := .FlagsNames }}
| `{{ $key }}` | {{ $value }} |
{{- end }}

## Upgrade notice

{{- range $item := .CLIs }}
{{- if $item.Analysis.UpgradeNotice.Cmd }}
- `{{ $item.Analysis.UpgradeNotice.Cmd }}`
  > **Note**
	> {{ $item.Analysis.UpgradeNotice.Note }}

  ```text
  {{- $item.Analysis.UpgradeNotice.Msg | trim | nindent 2 }}
  ```
	{{- if $item.Analysis.UpgradeNotice.Note }}
  {{- end }}
{{- end }}
{{- end }}

## What data is collected

{{ define "alternative" -}}
{{- if . -}}
  `{{- (trimSuffix "," .) -}}`
{{- end -}}
{{- end -}}

### Server

| Name | Occurrences | Alternative name |
|------|-------------|------------------|
{{- range $key, $details := .CollectedData.Server }}
|`{{ $key }}` | {{ $details.Cnt }} | {{ template "alternative" $details.Alternatives  | uniq | join "`,`" }}|
{{- end }}

### Client

| Name | Occurrences | Alternative name |
|------|-------------|------------------|
{{- range $key, $details := .CollectedData.Client }}
|`{{ $key }}` | {{ $details.Cnt }} | {{ template "alternative" $details.Alternatives  | uniq | join "," }}|
{{- end }}

### Dependencies

| Name | Occurrences | Alternative name |
|------|-------------|------------------|
{{- range $key, $details := .CollectedData.Dependencies }}
|`{{ $key }}` | {{ $details.Cnt }} | {{ template "alternative" $details.Alternatives  | uniq | join "," }}|
{{- end }}

## Output

{{- range $key, $val := .Output }}

### {{ $key | title }}

{{- range $item := $val }}
{{ $myDict := dict "json" "json" "yaml" "yaml" "plain" "text" "short" "text" }}
- `{{ $item.Cmd }}` {{ ( $item.Note ) }}
{{if $item.Out}}
  ```{{ get $myDict $key }}
{{ $item.Out | indent 2}}```
{{- end }}
{{- end }}
{{- end }}

## Frameworks' built-in support

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

## Others

- Interesting approach to print the version of deps: https://github.com/hashicorp/terraform/blob/v1.3.0-dev/version/dependencies.go
