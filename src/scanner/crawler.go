package scanner

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/labstack/gommon/log"
)

var fetchCount int
var customParams url.Values
var customHeaders []string

//CustomCrawler of https://github.com/PuerkitoBio/gocrawl
type CustomCrawler struct {
	gocrawl.DefaultExtender // Will use the default implementation of all but Visit and Filter
}

//Fetch overrides the default implementation in order to add custom parameters and headers
func (x *CustomCrawler) Fetch(ctx *gocrawl.URLContext, userAgent string, headRequest bool) (*http.Response, error) {
	urlQueryValues := ctx.URL().Query()
	for key, values := range customParams {
		for _, val := range values {
			urlQueryValues.Add(key, val)
		}
	}

	//TODO there must be a simpler way to clone an URL and add params
	newURL := &url.URL{
		RawQuery: urlQueryValues.Encode(),
		Host:     ctx.URL().Host,
		Scheme:   ctx.URL().Scheme,
		Fragment: ctx.URL().Fragment,
		Path:     ctx.URL().Path,
		RawPath:  ctx.URL().RawPath,
	}

	fetchCount++

	x.Log(gocrawl.LogEnqueued, gocrawl.LogInfo, "Override:"+newURL.String())

	//we copied the original Fetch() so we can change the URL (it's immutable) to add the query params
	var reqType string

	// Prepare the request with the right user agent
	if headRequest {
		reqType = "HEAD"
	} else {
		reqType = "GET"
	}
	req, e := http.NewRequest(reqType, newURL.String(), nil)
	if e != nil {
		return nil, e
	}
	req.Header.Set("User-Agent", userAgent)

	for _, raw := range customHeaders {
		header := strings.Split(raw, ":")
		if len(header) != 2 {
			log.Fatalf("Request headers must be under the form \"header: value\", invalid: %s", raw)
		}
		req.Header.Set(header[0], header[1])
	}

	return gocrawl.HttpClient.Do(req)
}
