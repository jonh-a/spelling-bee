package cmd

import (
	"fmt"
	"spelling_bee/pkg/api"
	"spelling_bee/pkg/utils"

	"github.com/spf13/cobra"
)

var guessCmd = &cobra.Command{
	Use:   "guess",
	Short: "Submit a guess for today's puzzle",
	Long:  `Submit a guess for today's puzzle`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		guess := args[0]

		score, err := api.GuessPuzzle(date, guess)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(score)
	},
}

func init() {
	rootCmd.AddCommand(guessCmd)

	guessCmd.PersistentFlags().StringP("date", "d", utils.GetTodaysDate(), "Puzzle date (YYYY-MM-DD)")
}
