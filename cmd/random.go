/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"github.com/spf13/cobra"
	// "golang.org/x/text/language"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random useless fact",
	Long: `This command fetches a random useless fact from the uselessfacts.jsph.pl api`,
	Run: func(cmd *cobra.Command, args []string) {
		language, err := cmd.Flags().GetString("language")

		if err != nil {
			log.Printf("Could not parse flags - %v", err)
		}
		getRandomFact(language)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	randomCmd.Flags().StringP("language", "l", "en", "Fact Language")
}

type Fact struct {
	ID string `json:"id"`
	Text string `json:text`
	Source string `json:source`
	Source_URL string `json:source_url`
	Language string `json:language`
	Permalink string `json:permalink`
}

func getRandomFact(language string) {
	var urlBuilder strings.Builder
	url := "https://uselessfacts.jsph.pl//api/v2/facts/random?language="
	urlBuilder.WriteString(url)
	urlBuilder.WriteString(language)
	resBytes := getFactData(urlBuilder.String())
	fact := Fact{}
	if err := json.Unmarshal(resBytes, &fact); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	fmt.Println(fact.Text)
}

func getFactData(baseAPI string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("Could not request a uselessfact - %v", err)
	}

	req.Header.Add("Acept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}

	resBytes, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("Could not read response body - %v", err)
	}

	return resBytes
}