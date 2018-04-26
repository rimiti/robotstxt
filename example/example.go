package main

import (
  "log"
  . ".."
  //. "github.com/rimiti/robotstxt"
  "strings"
)

func main() {
  url := "https://www.clicrdv.com/robots.txt"
  contents := `
        User-Agent: *
        Disallow: /map/
        Disallow: /404
        Disallow: /422
        Disallow: /500
        Disallow: /api/
        Disallow: /images/site3/logos-clients/
        Sitemap: https://www.clicrdv.com/fr/sitemap_index.xml        
    `

  robots, err := Parse(contents, url)
  if err != nil {
    log.Fatalln(err.Error())
  }

  allowed, _ := robots.IsAllowed("Browser-Bot/1.0", "http://www.clicrdv.com/test.html")
  if !allowed {
    println("Not allowed to crawl: /test.html")
  }

  allowed, _ = robots.IsAllowed("Browser-Bot/1.0", "http://www.clicrdv.com/dir/test.html")
  if allowed {
    println("Allowed to crawl: /dir/test.html")
  }

  added, _ := robots.AddAllowed("Browser-Bot/1.0", "http://www.clicrdv.com/dir/test.html")
  if !added {
    println("/dir/test.html allowed")
  }

  // [http://clicrdv.com/sitemap.xml]
  println("Sitemaps: " + strings.Join(robots.Sitemaps(), ","))

  // clicrdv.com
  println("Preferred host: " + robots.Host())
}
