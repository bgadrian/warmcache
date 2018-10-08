package scanner

import (
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/mkideal/cli"
)

//CLIArguments for the scanner function
type CLIArguments struct {
	cli.Helper
	Seed           []string `cli:"*seed" usage:"The start page (seed) of the crawl, example: https://google.com"`
	MaxPages       int      `cli:"max" usage:"Max number of pages that will be scanned, for each domain" dft:"10"`
	Delay          int      `cli:"delay" usage:"Milliseconds between 2 page visits, for each domain" dft:"400"`
	RobotUserAgent string   `cli:"robot" usage:"Name of the robot, for robots.txt" dft:"Googlebot"`
	UserAgent      string   `cli:"agent" usage:"User-agent for all requests" dft:"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"`
	Debug          bool     `cli:"debug" usage:"Print all pages that are found"`
	Query          []string `cli:"query" usage:"Add custom query params to all requests"`
	Headers        []string `cli:"header" usage:"Add one or more HTTP request headers to all requests" dft:"X-hotcache:crawler"`
}

//Scan a host with specific settings. Creates and run a crawler.
func Scan(ctx *cli.Context) error {
	var err error
	argv := ctx.Argv().(*CLIArguments)
	extender := new(CustomCrawler)
	opts := gocrawl.NewOptions(extender)

	opts.RobotUserAgent = argv.RobotUserAgent
	opts.UserAgent = argv.UserAgent
	opts.CrawlDelay = time.Duration(argv.Delay) * time.Millisecond
	opts.MaxVisits = argv.MaxPages

	if argv.Debug {
		opts.LogFlags = gocrawl.LogAll
	} else {
		opts.LogFlags = gocrawl.LogError
	}

	//TODO find a better way to send the params to Fetch() (instead of global vars)
	customHeaders = argv.Headers
	for _, queries := range argv.Query {
		customParam, e := url.ParseQuery(queries)
		if e != nil {
			return e
		}
		customParams = append(customParams, customParam)
	}

	c := gocrawl.NewCrawlerWithOptions(opts)
	err = c.Run(argv.Seed)

	log.Printf("fetched: %d urls (including robots.txt)\n", fetchCount)

	if err != nil {
		return err
	}

	return nil
}
