package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/grafvonb/c8volt/c8volt/ferrors"
	"github.com/grafvonb/c8volt/embedded"
	"github.com/spf13/cobra"
)

var (
	flagEmbedExportFileNames []string
	flagEmbedExportOut       string
	flagEmbedExportWithForce bool
)

var embedExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export embedded (virtual) resources to local files",
	Run: func(cmd *cobra.Command, args []string) {
		_, log, cfg, err := NewCli(cmd)
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, err)
		}

		all, err := embedded.List()
		if err != nil {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, err)
		}
		if len(flagEmbedExportFileNames) == 0 {
			ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("at least one --file is required"))
		}
		for _, f := range flagEmbedExportFileNames {
			if !slices.Contains(all, f) {
				ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("embedded file %q not found, run command 'embed list' to see all available embedded files, nothing exported", f))
			}
		}
		outBase := flagEmbedExportOut
		if outBase == "" {
			outBase = "."
		}
		var exported int
		for _, vpath := range flagEmbedExportFileNames {
			data, err := fs.ReadFile(embedded.FS, vpath)
			if err != nil {
				ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("read embedded %q: %w", vpath, err))
			}
			dst := filepath.Clean(filepath.Join(outBase, vpath))
			dir := filepath.Dir(dst)
			if err := os.MkdirAll(dir, 0o755); err != nil {
				ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("create dir %q: %w", dir, err))
			}
			if !flagEmbedExportWithForce {
				if info, err := os.Stat(dst); err == nil && info.Mode().IsRegular() {
					ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("destination file %q exists. use --force to overwrite", dst))
				}
			}
			if err := os.WriteFile(dst, data, 0o644); err != nil {
				ferrors.HandleAndExit(log, cfg.App.NoErrCodes, fmt.Errorf("write %q: %w", dst, err))
			}
			log.Debug(fmt.Sprintf("exported %q: %q", vpath, dst))
			exported++
		}
		log.Info(fmt.Sprintf("exported %d embedded resource(s) to %q", exported, filepath.Clean(outBase)))
	},
}

func init() {
	embedCmd.AddCommand(embedExportCmd)
	embedExportCmd.Flags().StringSliceVarP(&flagEmbedExportFileNames, "file", "f", nil, "embedded file(s) to export (repeatable)")
	embedExportCmd.Flags().StringVarP(&flagEmbedExportOut, "out", "o", ".", "output base directory")
	embedExportCmd.Flags().BoolVarP(&flagEmbedExportWithForce, "force", "", false, "overwrite if destination file exists")
	_ = embedExportCmd.MarkFlagRequired("file")
}
