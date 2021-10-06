package cli

import (
	"github.com/spf13/pflag"
	"os"
)

type Args struct {
	EnvPrefix    string
	TemplateFile string
	OutputFile string
}

func New() *Args {
	return &Args{
		EnvPrefix: parseStringFallback("ENV_PREFIX", ""),
		TemplateFile: parseStringFallback("TEMPLATE_FILE", ""),
		OutputFile: parseStringFallback("OUTPUT_FILE", ""),
	}
}

func parseStringFallback(envName, fallback string) string {
	val := os.Getenv(envName)
	if val == "" {
		return fallback
	}
	return val
}

func (a *Args) AddFlags(f *pflag.FlagSet) {
	f.StringVarP(&a.EnvPrefix, "env-prefix", "e", a.EnvPrefix, "the prefix of the env variables, for example KAFKA_")
	f.StringVarP(&a.TemplateFile, "template-file", "t", a.TemplateFile, "the template file, for example /opt/env-templater/template.tmpl")
	f.StringVarP(&a.OutputFile, "output-file", "o", a.OutputFile, "the output file full path, for example /opt/env-templater/server.properties")
}
