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
	Name      string `yaml:"name"`
	URL       string `yaml:"url"`
	Version   string `yaml:"version"`
	Framework string `yaml:"framework"`
	Analysis  struct {
		UpgradeNotice struct {
			Cmd  string `yaml:"cmd"`
			Note string `yaml:"note"`
			Msg  string `yaml:"msg"`
		} `yaml:"upgradeNotice"`
		VersionSkew bool              `yaml:"versionSkew"`
		Command     bool              `yaml:"command"`
		Flags       []string          `yaml:"flags"`
		Output      map[string]Output `yaml:"output"`
		Collected   struct {
			Dependencies []string `yaml:"dependencies"`
			Server       []string `yaml:"server"`
			Client       []string `yaml:"client"`
		} `yaml:"collected"`
	} `yaml:"analysis"`
}

type Output struct {
	Cmd  string `yaml:"cmd"`
	Out  string `yaml:"out"`
	Note string `yaml:"note"`
}

type Analysis struct {
	ViaCmdOnly   []CLI
	ViaFlagsOnly []CLI
	CLIs         []CLI

	FlagsNames     map[string]int
	ViaFlagsAndCmd []CLI

	Output map[string][]Output // output per type

	CollectedData CollectedData

	Frameworks map[string]int
}

type CollectedDataDetails struct {
	Cnt          int
	Alternatives []string
}

type CollectedData struct {
	Dependencies map[string]CollectedDataDetails
	Server       map[string]CollectedDataDetails
	Client       map[string]CollectedDataDetails
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
		CollectedData:  AggregateCollectedByType(clis),
		Frameworks:     AggregateFrameworks(clis),
	}
}

func AggregateFrameworks(clis []CLI) map[string]int {
	out := map[string]int{}
	for _, cli := range clis {
		out[cli.Framework] = out[cli.Framework] + 1
	}
	return out
}

func AggregateCollectedByType(clis []CLI) CollectedData {
	out := CollectedData{
		Dependencies: map[string]CollectedDataDetails{},
		Server:       map[string]CollectedDataDetails{},
		Client:       map[string]CollectedDataDetails{},
	}
	collect := func(in CollectedDataDetails, alt string) CollectedDataDetails {
		in.Cnt++
		if alt != "" {
			in.Alternatives = append(in.Alternatives, alt)
		}
		return in
	}
	for _, cli := range clis {
		for _, name := range cli.Analysis.Collected.Client {
			key, alternativeName, _ := strings.Cut(name, ":")
			out.Client[key] = collect(out.Client[key], alternativeName)
		}
		for _, name := range cli.Analysis.Collected.Server {
			key, alternativeName, _ := strings.Cut(name, ":")
			out.Server[key] = collect(out.Server[key], alternativeName)
		}

		for _, name := range cli.Analysis.Collected.Dependencies {
			key, alternativeName, _ := strings.Cut(name, ":")
			out.Dependencies[key] = collect(out.Dependencies[key], alternativeName)
		}
	}
	return out
}

func AggregateOutputByType(clis []CLI) map[string][]Output {
	out := map[string][]Output{}
	for _, cli := range clis {
		for name, resp := range cli.Analysis.Output {
			if strings.EqualFold(name, "json") || strings.EqualFold(name, "yaml") {
				name = strings.ToUpper(name)
			}
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
