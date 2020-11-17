# Youtuber Description Links

A tool to help Youtubers to manage theirs video links.

This project started on the 4th edition of #umaStackQueNaodomino which we learned how to use [Youtube Data API](https://developers.google.com/youtube/v3) with Go. You can check clips and vods at [Twitch.tv](https://www.twitch.tv/lucas_montano)

Previous Editions's Programming Languages

- [Typescript / Javascript with NodeJS (First Edition)](https://github.com/lucasmontano/twitch)
- [Python (Second Edition)](https://github.com/lucasmontano/magic-link)
- [Flutter / Dart Third Edition](https://github.com/lucasmontano/learn-tech)

## How to contribute

- Fork the project
- Create an ISSUE with your suggestion
- Write your code to address only the Issues scope
- Create a Pull Request and link to the Issues
- Wait for an approval
- Merge your Pull Request and close the Issues

### Important Notes

- Write your code, Issues and Pull Requests in English
- Avoid big Pull Requests and split your work properly

### Project Setup

Youtube API Developer Key: [Check Getting Started](https://developers.google.com/youtube/v3/getting-started)

- `export YOUTUBE_DEVELOPER_KEY=YOUR_KEY`

- Generate your client_secret.json:
  > [Youtube Go Quickstart](https://developers.google.com/youtube/v3/quickstart/go)

### Running application

- First, you must have the go compiler installed on your machine

  > [Official Docs from Golang website](https://golang.org/doc/install)

- Then, you must go to terminal and make:

```bash
cd youtuber-description-links
go get -v -t -d ./...
go run .
```

> **Note:** On your first running, your must have copy the link appear on the terminal, open this on web browser, authorize access and copy authentication code and paste on the terminal.
