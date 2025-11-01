package cmd

import (
	"github.com/grafvonb/kamunder/embedded"
	"github.com/grafvonb/kamunder/kamunder/ferrors"
	"github.com/grafvonb/kamunder/toolx"
	"github.com/grafvonb/kamunder/toolx/logging"
	"github.com/spf13/cobra"
)

var (
	flagEmbedListDetails bool
)

var embedListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List embedded (virtual) files containing process definitions",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		log, _ := logging.FromContext(cmd.Context())

		files, err := embedded.List()
		if err != nil {
			ferrors.HandleAndExit(log, err)
		}

		for _, f := range files {
			if flagViewAsJson {
				cmd.Println(toolx.ToJSONString(f))
			} else {
				cmd.Println(f)
			}
		}
	},
}

func init() {
	embedCmd.AddCommand(embedListCmd)
	embedListCmd.Flags().BoolVar(&flagEmbedListDetails, "details", false, "show full embedded file paths")
}
