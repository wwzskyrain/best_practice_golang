package main

import (
	"fmt"
	"net/url"

	"code.byted.org/middleware/hertz/pkg/app"
)

func main() {
	appCtx := app.RequestContext{}

	appCtx.Request.SetRequestURI("https://www.bytedance.com/path/to/me?k=v%&k2=v2")

	originURI := appCtx.Request.URI()
	forwardDomain := "https://pangolin16.sgsnssdk.com"
	forwardUrl := fmt.Sprintf("%s%s?%s", forwardDomain, string(originURI.Path()), string(originURI.QueryString()))
	fmt.Println(forwardUrl)

	parsedURL, _ := url.Parse("https://www.bytedance.com/path/to/me?k=v%&k2=v2")
	parsedURL.Host = "apm.pangle.io"
	parsedURL.Scheme = "https"
	fmt.Println(parsedURL.String())
}
