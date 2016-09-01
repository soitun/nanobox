package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/processors"
	"github.com/nanobox-io/nanobox/util/display"
	"github.com/nanobox-io/nanobox/commands/steps"
)

var (

	// TunnelCmd ...
	TunnelCmd = &cobra.Command{
		Use:   "tunnel",
		Short: "Creates a secure tunnel between your local machine & a production component.",
		Long: `
Creates a secure tunnel between your local machine & a
production component. The tunnel allows you to manage
production data using your local client of choice.
		`,
		PreRun: steps.Run("login"),
		Run:    tunnelFn,
	}

	// tunnelCmdFlags ...
	tunnelCmdFlags = struct {
		app  string
		port string
	}{}
)

//
func init() {
	TunnelCmd.Flags().StringVarP(&tunnelCmdFlags.app, "app", "a", "", "name or alias of a production app")
	TunnelCmd.Flags().StringVarP(&tunnelCmdFlags.port, "port", "p", "", "local port to start listening on")
}

// tunnelFn ...
func tunnelFn(ccmd *cobra.Command, args []string) {

	// validate we have args required to set the meta we'll need; if we don't have
	// the required args this will return with instructions
	if len(args) != 1 {
		fmt.Printf(`
Wrong number of arguments (expecting 1 got %v). Run the command again with the
name of the container you would like to tunnel into:

ex: nanobox tunnel <container>

`, len(args))

		return
	}

	// set the meta arguments to be used in the processor and run the processor
	tunnelConfig := processors.TunnelConfig{
		App:       tunnelCmdFlags.app,
		Port:      tunnelCmdFlags.port,
		Container: args[0],
	}

	display.CommandErr(processors.Tunnel(tunnelConfig))
}
