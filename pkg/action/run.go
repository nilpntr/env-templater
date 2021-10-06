package action

import (
	"github.com/nilpntr/env-templater/pkg/cli"
	"os"
	"strings"
	"text/template"
)

type TemplateData struct {
	Args      map[string]string
	ExtraArgs map[string]string
}

func NewRun(settings *cli.Args) error {
	envArgs := os.Environ()
	args := make(map[string]string)
	extraArgs := make(map[string]string)

	for _, envArg := range envArgs {
		if strings.HasPrefix(envArg, settings.EnvPrefix) {
			split := strings.Split(envArg, "=")
			args[strings.ReplaceAll(split[0], settings.EnvPrefix, "")] = split[1]
		}
	}
	for _, extraArg := range settings.ExtraArgs {
		split := strings.Split(extraArg, "=")
		if len(split) == 2 {
			extraArgs[split[0]] = split[1]
		}
	}

	data := TemplateData{Args: args, ExtraArgs: extraArgs}

	f, err := os.ReadFile(settings.TemplateFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"to_property": func(val string) string {
			return strings.ReplaceAll(strings.ToLower(val), "_", ".")
		},
	}).Parse(string(f))
	if err != nil {
		return err
	}

	o, err := os.Create(settings.OutputFile)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(o, data); err != nil {
		return err
	}

	return nil
}
