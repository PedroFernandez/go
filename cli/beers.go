package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

var beers = map[string]string{
	"AMAKLPF52B4U44H0UZ4L": "Voll-Damm",
	"K43EXHJNGYDPN4LUD18N": "Estrella Galicia",
	"7FYRUXSZIDCD3NY38FDR": "Heineken",
}

const idFlag = "id"

func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:                        "beers",
		Short:                      "Print data about beers",
		Run:                        runBeersfn(),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersfn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(beers[id])
		} else {
			fmt.Println(beers)
		}
	}
}