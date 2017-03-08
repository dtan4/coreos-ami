package cmd

import (
	"fmt"

	"github.com/dtan4/coreos-ami/coreos"
	"github.com/spf13/cobra"
)

var channels = []string{
	"alpha",
	"beta",
	"stable",
}

var verbose bool

func init() {
	for _, channel := range channels {
		command := &cobra.Command{
			Use:   channel,
			Short: fmt.Sprintf("Print AMIs of %s channel", channel),
			RunE: func(cmd *cobra.Command, args []string) error {
				amiFeed, err := coreos.RetrieveAMIFeed(cmd.Name())
				if err != nil {
					return err
				}

				if verbose {
					fmt.Printf("Version:      %s\n", amiFeed.ReleaseInfo["version"])
					fmt.Printf("Release date: %s\n", amiFeed.ReleaseInfo["release_date"])
					fmt.Printf("\n")

					fmt.Println("AMIs:")
				}

				fmt.Printf(amiFeed.TabularizeAMIs())

				return nil
			},
		}

		RootCmd.AddCommand(command)

		command.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print detailed information")
	}
}
