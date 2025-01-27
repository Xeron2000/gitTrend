# gitTrend - GitHub Trending Repositories Notifier (GitHub Actions Focused)

gitTrend is a tool designed to scrape trending repositories from GitHub and send daily notifications via Telegram. It is specifically designed to be used with GitHub Actions for fully automated daily updates, requiring no local setup for regular use.

## Features

- **GitHub Actions First**: Optimized for seamless integration with GitHub Actions for daily automated scraping and notifications.
- **Scrape GitHub Trending Repositories**: Automatically fetches the latest trending repositories based on your configured filters.
- **Telegram Notifications**: Sends daily updates of trending repositories directly to your Telegram chat (optional).
- **Configurable Filters**: Customize trends by spoken language, programming language, and time range via `config.json`.
- **Minimal Setup**: Primarily designed for GitHub Actions, eliminating the need for local installation for daily operation.

## GitHub Actions Setup (Recommended - No Local Setup Needed)

Get started with daily trending GitHub repository notifications in just a few steps using GitHub Actions:

1.  **Repository Secrets**:
    -   Navigate to your GitHub repository settings -> "Secrets and variables" -> "Actions".
    -   Create the following repository secrets:
        -   `TELEGRAM_BOT_TOKEN`: Your Telegram Bot API token (required for Telegram notifications).
        -   `TELEGRAM_CHAT_ID`: Your Telegram Chat ID (required for Telegram notifications).

2.  **Configuration via `config.json`**:
    -   Modify the `config.json` file in your repository to define your preferred filters and Telegram notification settings.

    ```json
    {
      "spokenLanguage": "",
      "progLanguage": "",
      "timeRange": "daily",
      "telegramBotToken": "",
      "telegramChatID": ""
    }
    ```

    -   `spokenLanguage`: Filter by spoken language (e.g., `en`, `zh`). Leave empty for all languages.
    -   `progLanguage`: Filter by programming language (e.g., `go`). Leave empty for all programming languages.
    -   `timeRange`: Set the time range for trending repositories: `daily`, `weekly`, or `monthly`. Default is `daily`.
    -   `telegramBotToken`: Your Telegram Bot API token. Required for enabling Telegram notifications.
    -   `telegramChatID`: Your Telegram Chat ID. Required if `telegramBotToken` is provided.

3.  **Enable Workflow**:
    -   Ensure the GitHub Actions workflow located at `.github/workflows/scrape-and-notify.yml` is enabled in your repository. This workflow is pre-configured to run daily at midnight UTC.

With these steps, you'll receive daily updates of GitHub trending repositories directly in your Telegram chat, completely automated through GitHub Actions!

## Optional Local Testing

For testing purposes or if you wish to run gitTrend locally, follow these steps:

### Prerequisites for Local Testing

- [Go](https://golang.org/dl/) 1.22.4+
- [Git](https://git-scm.com/downloads)

### Local Testing Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/gitTrend.git
   cd gitTrend
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the binary:

   ```bash
   go build -o gt
   ```

### Run Locally

```bash
./gt scrape
```

Before running locally, ensure your `config.json` is properly set up, especially if testing Telegram notifications.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
