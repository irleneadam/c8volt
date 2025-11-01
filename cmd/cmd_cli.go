package cmd

import (
	"fmt"
	"log/slog"

	"github.com/grafvonb/kamunder/config"
	"github.com/grafvonb/kamunder/kamunder"
	"github.com/grafvonb/kamunder/toolx/logging"
	"github.com/spf13/cobra"
)

func NewCli(cmd *cobra.Command) (kamunder.API, *slog.Logger, *config.Config, error) {
	log, _ := logging.FromContext(cmd.Context())
	svcs, err := NewFromContext(cmd.Context())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error getting services from context: %w", err)
	}
	cli, err := kamunder.New(
		kamunder.WithConfig(svcs.Config),
		kamunder.WithHTTPClient(svcs.HTTP.Client()),
		kamunder.WithLogger(log),
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error creating kamunder client: %w", err)
	}
	return cli, log, svcs.Config, nil
}
