package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

type Config struct {
	SpokenLanguage   string `json:"spokenLanguage"`
	ProgLanguage     string `json:"progLanguage"`
	TimeRange        string `json:"timeRange"`
	TelegramBotToken string `json:"telegramBotToken"`
	TelegramChatID   string `json:"telegramChatID"`
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape GitHub trending repositories",
	Run: func(cmd *cobra.Command, args []string) {
		scrapeTrending()
	},
}

func scrapeTrending() {
	// Load config from file
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}

	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	var projects []string

	c.OnHTML("article.Box-row", func(e *colly.HTMLElement) {
		projectName := strings.TrimSpace(e.ChildText("h2 a"))
		projectName = strings.Join(strings.Fields(projectName), " ")

		projectURL := strings.TrimSpace(e.ChildAttr("h2 a", "href"))
		projectDescription := strings.TrimSpace(e.ChildText("p.pr-4"))

		projectInfo := fmt.Sprintf("Name: %s\nURL: https://github.com%s\nDescription: %s\n\n", projectName, projectURL, projectDescription)
		projects = append(projects, projectInfo)
		fmt.Print(projectInfo) // Print to console as before
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL)
	})

	baseURL := "https://github.com/trending"
	if config.ProgLanguage != "" {
		baseURL += fmt.Sprintf("/%s", config.ProgLanguage)
	}

	queryParams := []string{}
	if config.SpokenLanguage != "" {
		queryParams = append(queryParams, fmt.Sprintf("spoken_language_code=%s", config.SpokenLanguage))
	}
	if config.TimeRange != "" {
		queryParams = append(queryParams, fmt.Sprintf("since=%s", config.TimeRange))
	}

	if len(queryParams) > 0 {
		baseURL += "?" + strings.Join(queryParams, "&")
	}

	err = c.Visit(baseURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Send Telegram notification
	if config.TelegramBotToken != "" && config.TelegramChatID != "" {
		message := "GitHub Trending Repositories:\n\n" + strings.Join(projects, "")
		err = sendTelegramNotification(config.TelegramBotToken, config.TelegramChatID, message)
		if err != nil {
			fmt.Println("Error sending Telegram notification:", err)
		} else {
			fmt.Println("Telegram notification sent!")
		}
	}
}

func sendTelegramNotification(botToken, chatID, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	formData := url.Values{
		"chat_id": []string{chatID},
		"text":    []string{message},
	}

	_, err := http.PostForm(apiURL, formData)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
	// removed flags, configurations are read from config.json
}
