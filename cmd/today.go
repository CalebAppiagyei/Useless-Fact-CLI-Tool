/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get the useless fact of the day",
	Long: `This command fetches the useless fact of the day from the uselessfacts.jsph.pl api`,
	Run: func(cmd *cobra.Command, args []string) {
		language, err := cmd.Flags().GetString("language")

		if err != nil {
			log.Printf("Could not parse flags - %v", err)
		}

		getTodaysFact(language)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)

	todayCmd.Flags().StringP("language", "l", "en", "Fact Language")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getTodaysFact (language string) {
	var urlBuilder strings.Builder
	url := "https://uselessfacts.jsph.pl//api/v2/facts/today?language="
	urlBuilder.WriteString(url)
	urlBuilder.WriteString(language)
	resBytes := getFactData(urlBuilder.String())
	fact := Fact{}
	if err := json.Unmarshal(resBytes, &fact); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	fmt.Println(fact.Text)
}


