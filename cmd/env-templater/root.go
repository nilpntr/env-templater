package main

import (
	"github.com/nilpntr/env-templater/pkg/action"
	"github.com/nilpntr/env-templater/pkg/cli"
	"github.com/spf13/cobra"
)

var settings = cli.New()

func newRootCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "env-templater",
		Short:        "env-templater is a tool to simply create env based templated files",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return action.NewRun(settings)
		},
	}

	flags := cmd.PersistentFlags()
	settings.AddFlags(flags)

	flags.ParseErrorsWhitelist.UnknownFlags = true

	return cmd, nil
}
