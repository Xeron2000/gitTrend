package cmd

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var (
	spokenLanguage string
	progLanguage   string
	timeRange      string
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape GitHub trending repositories",
	Run: func(cmd *cobra.Command, args []string) {
		scrapeTrending(spokenLanguage, progLanguage, timeRange)
	},
}

func scrapeTrending(spokenLang, progLang, timeRange string) {
	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	c.OnHTML("article.Box-row", func(e *colly.HTMLElement) {
		projectName := strings.TrimSpace(e.ChildText("h2 a"))
		projectName = strings.Join(strings.Fields(projectName), " ")

		projectURL := strings.TrimSpace(e.ChildAttr("h2 a", "href"))
		projectDescription := strings.TrimSpace(e.ChildText("p.pr-4"))

		fmt.Printf("Name: %s\nURL: https://github.com%s\nDescription: %s\n\n", projectName, projectURL, projectDescription)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL)
	})

	baseURL := "https://github.com/trending"
	if progLang != "" {
		baseURL += fmt.Sprintf("/%s", progLang)
	}

	queryParams := []string{}
	if spokenLang != "" {
		queryParams = append(queryParams, fmt.Sprintf("spoken_language_code=%s", spokenLang))
	}
	if timeRange != "" {
		queryParams = append(queryParams, fmt.Sprintf("since=%s", timeRange))
	}

	if len(queryParams) > 0 {
		baseURL += "?" + strings.Join(queryParams, "&")
	}

	err := c.Visit(baseURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	scrapeCmd.Flags().StringVarP(&spokenLanguage, "spoken", "s", "", "Spoken language code ( en for English, zh for Chinese)")
	scrapeCmd.Flags().StringVarP(&progLanguage, "language", "l", "", "Programming language ( go for Go language)")
	scrapeCmd.Flags().StringVarP(&timeRange, "time", "t", "daily", "Time range ( daily,  weekly, monthly)")
}
