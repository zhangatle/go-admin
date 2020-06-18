package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"go-admin/cmd/api"
	"go-admin/cmd/migrate"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "go-admin",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Printf("%v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `go-admin 1.0.0 欢迎使用，可以使用 -h 查看命令`
		log.Printf("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
