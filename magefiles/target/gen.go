package target

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

type CLI struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Version   string `json:"version"`
	Framework string `json:"framework"`
	Analysis  struct {
		VersionSkew bool              `json:"versionSkew"`
		Command     bool              `json:"command"`
		Flags       []string          `json:"flags"`
		Output      map[string]Output `json:"output"`
		Collected   struct {
			Dependencies []string `json:"dependencies"`
			Server       []string `json:"server"`
			Client       []string `json:"client"`
		} `json:"collected"`
	} `json:"analysis"`
}

type Output struct {
	Cmd  string `json:"cmd"`
	Out  string `json:"out"`
	Note string `json:"note"`
}

type Analysis struct {
	ViaCmdOnly   []CLI
	ViaFlagsOnly []CLI
	CLIs         []CLI

	FlagsNames     map[string]int
	ViaFlagsAndCmd []CLI

	Output map[string][]Output // output per type
}

func Template(tplFile, outFile string) {
	analysis := CollectAnalysis()

	tpl := lo.Must(template.New(filepath.Base(tplFile)).Funcs(sprig.FuncMap()).ParseFiles(tplFile))
	out := lo.Must(os.Create(outFile))

	lo.Must0(tpl.Execute(out, analysis))
}

func CollectAnalysis() Analysis {
	dir := lo.Must(os.Getwd())

	var clis []CLI
	lo.Must0(filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) != ".yaml" {
			return nil
		}

		var out CLI
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(raw, &out); err != nil {
			return fmt.Errorf("while unmarshaling %s: %w", path, err)
		}
		out.Name = strings.TrimSuffix(filepath.Base(path), ".yaml")
		clis = append(clis, out)
		return nil
	}))

	return Analysis{
		CLIs:           clis,
		ViaCmdOnly:     ViaCommand(clis),
		ViaFlagsAndCmd: ViaFlagsAndCommand(clis),
		ViaFlagsOnly:   ViaFlags(clis),
		FlagsNames:     FlagsNames(clis),
		Output:         AggregateOutputByType(clis),
	}

}

func AggregateOutputByType(clis []CLI) map[string][]Output {
	out := map[string][]Output{}
	for _, cli := range clis {
		for name, resp := range cli.Analysis.Output {
			out[name] = append(out[name], resp)
		}
	}
	return out
}

func ViaFlagsAndCommand(in []CLI) []CLI {
	var out []CLI
	for _, cli := range in {
		if !cli.Analysis.Command || len(cli.Analysis.Flags) == 0 {
			continue
		}
		out = append(out, cli)
	}
	return out
}

func ViaCommand(in []CLI) []CLI {
	var out []CLI
	for _, cli := range in {
		if !cli.Analysis.Command || len(cli.Analysis.Flags) > 0 {
			continue
		}
		out = append(out, cli)
	}
	return out
}

func FlagsNames(in []CLI) map[string]int {
	flags := map[string]int{}
	for _, cli := range in {
		if len(cli.Analysis.Flags) == 0 {
			continue
		}
		for _, name := range cli.Analysis.Flags {
			cnt := flags[name]
			cnt++
			flags[name] = cnt
		}
	}
	return flags
}

func ViaFlags(in []CLI) []CLI {
	var out []CLI
	for _, cli := range in {
		if cli.Analysis.Command || len(cli.Analysis.Flags) == 0 {
			continue
		}
		out = append(out, cli)
	}
	return out
}
