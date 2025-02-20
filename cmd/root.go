package cmd

import (
	"fmt"
	"os"

	"github.com/hatredholder/bookbrowse/internal"
	"github.com/hatredholder/bookbrowse/internal/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bookbrowse [book name]",
	Short: "bookbrowse allows you to browse books through hardcover.app in your terminal",
	Long: `bookbrowse - search books within your terminal

Example:
  $ bookbrowse "The Old Man and the Sea" -f
  `,
	Version: "v0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		if os.Getenv("HARDCOVER_API_KEY") == "" {
			fmt.Println("error: HARDCOVER_API_KEY environment variable must be set")
			fmt.Println("get an API key at https://hardcover.app/account/api")
			os.Exit(0)
		}

		hits := api.Query(args)
		book := internal.Chooser(hits)
		result := internal.Format(book, cmd.Flags())

		fmt.Print(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("markdown", "m", false, "format output into markdown")
	rootCmd.Flags().BoolP("fulldesc", "f", false, "display full description")
}
