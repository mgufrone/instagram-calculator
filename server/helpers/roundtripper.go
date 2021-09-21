package helpers

import (
	"net/http"
)

type RoundTripper struct {
	original http.RoundTripper
}

func NewRoundTripper(ori http.RoundTripper) *RoundTripper {
	return &RoundTripper{ori}
}

func (r *RoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	headers := map[string]string{
		"authority":                 "www.instagram.com",
		"cache-control":             "max-age=0",
		"sec-ch-ua":                 `Google Chrome";v="93", " Not; A Brand";v="99", "Chromium";v="93"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"macOS"`,
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-user":            "?1",
		"sec-fetch-dest":            "document",
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	return r.original.RoundTrip(request)
}
