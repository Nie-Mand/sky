# Sky, the Scraper

## Description
Just a scraper for video content

## Supported Sites
  [aCloudGuru](https://acloud.guru)



## Setup
1. Install [Go](https://golang.org/)
2. Add the environment variable for token
3. Install Dependencies with `go mod tidy`
4. run `go build .`
5. run `mv sky-scraper sky`

## Status
- [x] aCloudGuru
  - ✅ Supports Video Download
  - ✅ Supports Material Download
  - [ ] Supports Lab Download
  - [ ] Supports Quiz Download


## Usage

```bash
Download a course from a learning platform.

Usage:
  sky get [course ID] [flags]

Flags:
  -g, --guru             Set the provider as a Cloud Guru
  -h, --help             help for get
  -q, --quality string   Set the Quality of the video to download (default "720p")
```