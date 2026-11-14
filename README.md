# Golang Mailchimp Email Sender

A small Go project for sending bulk email messages from a CSV list of recipients using a worker pool and HTML/text email template. It is designed for local testing with Mailpit, an SMTP capture server, and can be extended to a real mail provider or marketing platform later.

## Overview

This project reads recipient data from `emails.csv`, loads it into a channel, and distributes records to multiple goroutines that send emails via SMTP. The template is defined in `email.tmpl` and filled with each recipient's `Name` and `Email` values.

## Features

- Reads recipient records from a CSV file
- Uses Go channels and worker goroutines for concurrent sending
- Sends emails through an SMTP endpoint
- Renders email content from a template file
- Built for local testing with Mailpit

## Project Structure

```text
.
├── README.md
├── consumer.go        # Email sending workers
├── producer.go        # CSV loading logic
├── main.go            # Program entry point and worker startup
├── email.tmpl         # Email template
├── emails.csv         # Recipient list
├── go.mod             # Go module definition
└── .gitignore
```

## Prerequisites

- Go 1.20 or newer
- Mailpit installed and running locally
- A CSV file with at least two columns: `Name`, `Email`

## Install Mailpit

On macOS:

```bash
brew install mailpit
```

On Linux / Windows, follow the official Mailpit installation instructions for your platform.

## Start Mailpit

```bash
mailpit
```

By default, Mailpit listens on:

- SMTP: `localhost:1025`
- Web UI: `http://localhost:8025`

## CSV Format

The `emails.csv` file should have a header row like this:

```csv
Name,Email
John Doe,john@example.com
Jane Smith,jane@example.com
```

## Email Template

The email content is defined in `email.tmpl`:

```text
To:{{.Email}}
Subject: Hello, {{.Name}}


Hi, {{.Name}}

Thanks,
The CG Campaing team.
```

This template is rendered for each recipient using Go's `html/template` package.

## Run the Application

From the project root:

```bash
go run .
```

The application will:

1. Read `emails.csv`
2. Pass each recipient into a channel
3. Start 5 worker goroutines
4. Send each email through `localhost:1025`

## Configure SMTP

The project currently uses this SMTP setup in `consumer.go`:

```go
smtpHost := "localhost"
smtpPort := "1025"
```

If you want to use a real provider or another SMTP server, update the values in `consumer.go` accordingly.

## Notes

- This project is a local email testing utility rather than a full Mailchimp API integration.
- The current implementation sends to a local SMTP server for development and testing.
- It is intentionally simple and suitable for learning Go concurrency and email automation.

## Example Workflow

```bash
# 1. Start Mailpit
mailpit

# 2. Update your recipient list
# edit emails.csv

# 3. Run the sender
go run .

# 4. Open Mailpit dashboard
# http://localhost:8025
```

## License

This project is provided as-is for educational and local testing purposes.
