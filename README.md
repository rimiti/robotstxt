# robotstxt

A robots.txt parser and generator written in Go, based on [samclarke robotstxt package](https://github.com/samclarke/robotstxt).

It currently supports:

  * User-agent:
  * Allow:
  * Disallow:
  * Sitemap:
  * Crawl-delay:
  * Host:
  * URL encoded & UTF-8 paths
  * Paths with wildcards (*) and EOL matching ($)

## Installation

Go get:

    go get github.com/rimiti/robotstxt

## Usage

```go
import (
    "log"
    "github.com/rimiti/robotstxt"
)

func main() {
    url := "http://www.example.com/robots.txt"
    contents := `
        User-agent: *
        Disallow: /dir/
        Disallow: /test.html
        Allow: /dir/test.html
        Allow: /test.html
        Crawl-delay: 1
        Sitemap: http://example.com/sitemap.xml
        Host: example.com
    `

    robots, err := Parse(contents, url)
    if err != nil {
        log.Fatalln(err.Error())
    }

    allowed, _ := robots.IsAllowed("Bot/1.0", "http://www.example.com/test.html")
    if !allowed {
        println("Not allowed to crawl: /test.html")
    }

    allowed, _ = robots.IsAllowed("Bot/1.0", "http://www.example.com/dir/test.html")
    if allowed {
        println("Allowed to crawl: /dir/test.html")
    }

    // 1
    println("Crawl delay: " + robots.CrawlDelay("Bot/1.0"))

    // [http://example.com/sitemap.xml]
    println("Sitemaps: " + strings.Join(robots.Sitemaps(), ","))

    // example.com
    println("Preferred host: " + robots.Host())
}
```