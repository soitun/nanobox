package commands

import (
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/processors"
	"github.com/nanobox-io/nanobox/util/display"
)

var (

	// StopCmd ...
	StopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop the Nanobox virtual machine.",
		Long: `
Stops the Nanobox virtual machine as well as any running
dev and sim platforms.
		`,
		Run:    stopFn,
	}
)

// stopFn ...
func stopFn(ccmd *cobra.Command, args []string) {
	display.CommandErr(processors.Stop())
}
