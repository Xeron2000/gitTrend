# gitTrend - GitHub Trending Repositories Notifier

gitTrend is a command-line tool to scrape trending repositories from GitHub and optionally send daily notifications via Telegram. It is designed to be used with Linux cron for automated daily updates.

## Features

- **Scrape GitHub Trending Repositories**: Fetches the latest trending repositories based on configured filters.
- **Telegram Notifications**: Sends daily updates of trending repositories to your Telegram chat (optional).
- **Configurable Filters**: Filter trends by spoken language, programming language, and time range via `config.json`.
- **Cron Scheduling**: Designed for automated daily runs using Linux cron.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.22.4+
- [Git](https://git-scm.com/downloads)

### Running from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/gitTrend.git
   cd gitTrend
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

Configure the application using the `config.json` file in the project root.

```json
{
  "spokenLanguage": "",
  "progLanguage": "",
  "timeRange": "daily",
  "telegramBotToken": "",
  "telegramChatID": ""
}
```

- `spokenLanguage`: Filter trending repositories by spoken language (e.g., `en` for English, `zh` for Chinese). Leave empty for all languages.
- `progLanguage`: Filter trending repositories by programming language (e.g., `go` for Go). Leave empty for all languages.
- `timeRange`: Time range for trending repositories (`daily`, `weekly`, `monthly`). Default is `daily`.
- `telegramBotToken`: Telegram Bot API token for sending notifications. Optional, leave empty to disable notifications.
- `telegramChatID`: Telegram Chat ID to send notifications to. Optional, required if `telegramBotToken` is set.

## Usage

To scrape trending repositories and optionally send Telegram notifications, run:

```bash
go run main.go scrape
```

The tool reads configurations from `config.json`.  You can modify `config.json` to customize the scraping filters and Telegram notification settings.

## Setting up Cron for Daily Notifications (Linux)

To automate daily scraping and Telegram notifications using Linux cron:

1. **Install prerequisites and clone the repository**: Ensure you have Go and Git installed, and have cloned the gitTrend repository as described in the [Installation](#installation) section.
2. **Configure `config.json`**:  Set your `telegramBotToken` and `telegramChatID` in the `config.json` file if you want Telegram notifications.
3. **Edit crontab**: Open your crontab file to add a new cron job:

   ```bash
   crontab -e
   ```

4. **Add cron job**: Add the following line to your crontab file to run the scraper daily at midnight:

   ```cron
   0 0 * * * go run main.go scrape
   ```

   This line schedules the `go run main.go scrape` command to run every day at midnight (00:00). Adjust the time as needed. Save and exit the crontab file.

With these steps, `gitTrend` will automatically scrape GitHub trending repositories daily and send Telegram notifications (if configured) using Linux cron.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
