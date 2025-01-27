## Introduction

![gt](https://raw.githubusercontent.com/Xeron2000/gitTrend/main/gt.png)

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.22.4+
- [Git](https://git-scm.com/downloads)

### Installing from Source

1. Clone the repository

   ```
   git clone https://github.com/Xeron2000/gitTrend.git
   cd gitTrend
   ```

2. Install dependencies

```
go mod tidy
```

3. Build the binary

```
go build -o gt
cp gt /usr/local/bin/
```

### Configuration

Before running the application, configure the `config.json` file in the project root. See the [Configuration](#configuration) section for details.

4. Run the project

```
gt scrape
```

This command will scrape GitHub trending repositories based on the configurations in `config.json` and send Telegram notifications if configured.

## Configuration

The application now uses a `config.json` file to manage configurations.

```json
{
  "spokenLanguage": "",
  "progLanguage": "",
  "timeRange": "daily",
  "telegramBotToken": "",
  "telegramChatID": ""
}
```

- `spokenLanguage`:  Spoken language code for trending repositories (e.g., `en` for English, `zh` for Chinese).
- `progLanguage`: Programming language for trending repositories (e.g., `go` for Go language).
- `timeRange`: Time range for trending repositories (`daily`, `weekly`, `monthly`).
- `telegramBotToken`: Telegram Bot API token for sending notifications (optional).
- `telegramChatID`: Telegram Chat ID to send notifications to (optional).

## GitHub Actions for Daily Notifications

The project includes a GitHub Actions workflow to automatically scrape trending repositories daily and send notifications to Telegram if configured.

To set up daily notifications:

1.  **Configure `config.json`**:  Set your `telegramBotToken` and `telegramChatID` in the `config.json` file.
2.  **Set up GitHub Secrets**:
    -   Go to your GitHub repository settings -> "Secrets and variables" -> "Actions".
    -   Add two repository secrets:
        -   `TELEGRAM_BOT_TOKEN`: Your Telegram Bot API token.
        -   `TELEGRAM_CHAT_ID`: Your Telegram Chat ID.

The workflow is defined in `.github/workflows/scrape-and-notify.yml` and runs daily at midnight UTC.

## License

This project is licensed under the MIT License. For more details, see the [LICENSE](LICENSE) file.
