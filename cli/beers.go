package cli

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
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
		f, _ := os.Open("data/beers.csv")
		reader := bufio.NewReader(f)

		var beers = make(map[int]string)

		for line := readLine(reader); line != nil; line = readLine(reader){
			values := strings.Split(string(line), ",")

			productID, _ := strconv.Atoi(values[0])
			beers[productID] = values[1]

			id, _ := cmd.Flags().GetString(idFlag)

			if id != "" {
				i, _ := strconv.Atoi(id)
				fmt.Println(beers[i])
			} else {
				fmt.Println(beers)
			}
		}
	}
}

func readLine(reader *bufio.Reader) (line[]byte) {
	line, _, _= reader.ReadLine()
	return
}