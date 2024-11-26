package cmd

import (
	"fmt"
	"spelling_bee/pkg/utils"

	"github.com/spf13/cobra"
)

var guessCmd = &cobra.Command{
	Use:   "guess",
	Short: "Submit a guess for today's puzzle",
	Long:  `Submit a guess for today's puzzle`,
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		fmt.Println(date)
	},
}

func init() {
	rootCmd.AddCommand(guessCmd)

	guessCmd.PersistentFlags().StringP("date", "d", utils.GetTodaysDate(), "Puzzle date (YYYY-MM-DD)")
}
