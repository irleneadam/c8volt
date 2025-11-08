package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/grafvonb/c8volt/c8volt"
	"github.com/grafvonb/c8volt/config"
	"github.com/grafvonb/c8volt/toolx/logging"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var ErrCmdAborted = errors.New("aborted by user")

func NewCli(cmd *cobra.Command) (c8volt.API, *slog.Logger, *config.Config, error) {
	log, _ := logging.FromContext(cmd.Context())
	svcs, err := NewFromContext(cmd.Context())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error getting services from context: %w", err)
	}
	cli, err := c8volt.New(
		c8volt.WithConfig(svcs.Config),
		c8volt.WithHTTPClient(svcs.HTTP.Client()),
		c8volt.WithLogger(log),
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error creating c8volt client: %w", err)
	}
	return cli, log, svcs.Config, nil
}

func confirmCmdOrAbort(autoConfirm bool, prompt string) error {
	if autoConfirm || !term.IsTerminal(int(os.Stdin.Fd())) {
		return nil
	}
	fmt.Printf("%s [y/N]: ", prompt)
	in := bufio.NewScanner(os.Stdin)
	if !in.Scan() {
		return ErrCmdAborted
	}
	switch strings.ToLower(strings.TrimSpace(in.Text())) {
	case "y", "yes":
		return nil
	default:
		return ErrCmdAborted
	}
}
