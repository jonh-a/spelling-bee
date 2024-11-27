package cmd

import (
	"fmt"
	"spelling_bee/pkg/api"
	"spelling_bee/pkg/ui"
	"spelling_bee/pkg/utils"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get today's puzzle",
	Long:  `Get today's puzzle`,
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		puzzle, err := api.GetPuzzle(date)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ui.RenderGet(puzzle))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringP("date", "d", utils.GetTodaysDate(), "Puzzle date (YYYY-MM-DD)")
}
