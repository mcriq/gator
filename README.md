---
title: Gator
tags:
  - RSS Aggregator
  - News
---
## Summary

Gator is an RSS feed aggregator

## Installation
You'll need Postgres and Go installed to run the program
To install:
1. `go install`
2. Set up config
3. Run some commands:
- `go run . register <username>`
- `go run . addfeed "TechCrunch" "https://techcrunch.com/feed/"`