<<<<<<< HEAD
## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

=======
>>>>>>> a3fdbc2 (update README)
## Introduction

![gt](https://raw.githubusercontent.com/Xeron2000/gitTrend/main/gt.png)

## Installation

### Using the Pre-built Binary

See [Releases](https://github.com/Xeron2000/gitTrend/releases)

### Installing from Source

- [Go](https://golang.org/dl/) 1.22.4+
- [Git](https://git-scm.com/downloads)

1. Clone the repository

   ```
   git clone https://github.com/Xeron2000/gitTrend.git
   cd gitTrend
   ```

2. Install dependencies

```
go mod tidy
```

3.  Build the binary

```
go build -o gt
cp gt /usr/local/bin/
```
4.  Run the project

```
gt scrape
```

## Usage

```
gt scrape -s en -l go -t daily
```

Parameters explanation:

- `-s, --spoken`: Specify the spoken language code (e.g., `en` for English).
- `-l, --language`: Specify the programming language (e.g., `go` for Go language).
- `-t, --time`: Specify the time range (e.g., `daily` for daily, `weekly` for weekly, `monthly` for monthly).

## License

This project is licensed under the MIT License. For more details, see the [LICENSE](LICENSE) file.
