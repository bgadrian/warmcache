#hot cache crawler

Problem: you have a lazy-init cache system (like a CDN or Prerender), the first visit for each page will be slower.
Solution: visit all your pages to warm-up the cache

This tool is aimed to more static websites, Single Page Apps with unique URLs (example React with React Router) cached by CDNs or websites that uses Prerender (or similar solutions for SEO).

## What does it do?

You provide one (or more) URLs of HTML pages and it starts crawling the websites, triggering the HTML cache (if any).

## How?

1. It starts from the seeds (starting pages) you provide
2. For each different domain found in the seeds starts a visit queue
3. Visit a page from the queue
    * Adds all the URLs found on that page in the queue
    * Waits `--delay` milliseconds
    * if the `--max` visited pages was reached, it exits
    * repeat the step 3 until the queue is empty

## Cache agnostic

The crawler has nothing to do with your cache directly, it is a simple HTTP client that visits your website.

The tool was originally written to be used with Prerender, a service for Single Page Apps that does Server Side Rendering for bots (search engines like Google and previews like WhatsApp and Facebook). So by default it adds the `_escaped_fragment_` query parameter and uses the Googlebot user-agent (to trigger the Prerender cache), but you can change the parameters values.

The crawler is caching-agnostic, because it has a simple logic and just visits the web pages it can be used as a "warm" companion for any caching solution like Squid, CDN and so on.

### When does it stop?
When it reaches the `--max` (pages) visited per domain or visits all the found URLs.

### Failsafes
* It does **not** follow the links that are on different domains than the ones found in seeds.
* Default `--max` (pages) value is 10, so by default it will stop after 10 crawled links.
* To avoid triggering a DDOS attack protection it has a default waiting time of 400ms between 2 requests on each domain.
* It follows redirects but instead of visiting them directly it adds them to the queue. As a result, each redirect will count as a different visited page.
* It adds a custom HTTP header `X-hotcache:crawler`, you can use it to make custom firewall/app rules.

## Usage
```bash
$ ./build/hotcache -help
#output:
  -h, --help                                                                               display help information
      --seed                                                                              *The start page (seed) of the crawl, example: https://google.com
      --max[=10]                                                                           Max number of pages that will be scanned, for each domain
      --delay[=400]                                                                        Milliseconds between 2 page visits, for each domain
      --robot[=Googlebot]                                                                  Name of the robot, for robots.txt
      --agent[=Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)]   User-agent for all requests
      --debug                                                                              Print all pages that are found
      --query                                                                              Add custom query params to all requests
      --header[=X-hotcache:crawler]                                                        Add one or more HTTP request headers to all requests

```
Crawl trigger for Prerender:
```bash 
$ ./build/hotcache --seed http://localhost/ --debug --query "_escaped_fragment_="```
````
Simple crawl of 2 domains, with a maximum of 400 visited pages:

```bash
 ./build/hotcache --seed http://domain1 --seed https://domain2 --max 400
```
Custom delay time and user-agent:
```bash
./build/hotcache --seed https://domain1 --delay 250 --robot "mybot" --agent "Mozilla/5.0 (compatible; MyBot/1.0)" 
```
## Test

```bash
#optional, run your own intance of http://httpbin.org, in one terminal:
$ docker run -p 80:80 kennethreitz/httpbin

#in other terminal:
$ make build
./build/hotcache --seed http://localhost/anything --debug --query "test=1" --query "_escaped_fragment_=1" --header "Accept: application/json"
```
You should see in the output the httpbin echo with all the parameters and custom headers.

## Libraries
To collect the CLI parameters: https://github.com/mkideal/cli
To crawl the websites I used: https://github.com/PuerkitoBio/gocrawl that obey to robots.txt rules (using the robotstxt.go library) and launches 1 goroutine for each domain.

## TODO
* improve the current code
* tests
* optional, trigger the cache (make requests) for all resources found on the same domain (images, CSS, JavaScript, audio, video ...)