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

				fmt.Printf(amiFeed.Tabularize())

				return nil
			},
		}

		RootCmd.AddCommand(command)
	}
}
