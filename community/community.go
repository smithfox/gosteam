package community

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func SetCookies(client *http.Client, sessionId, steamLogin, steamLoginSecure string) {
	if client.Jar == nil {
		client.Jar, _ = cookiejar.New(nil)
	}
	base, err := url.Parse("http://steamcommunity.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("base.host=%s,scheme=%v\n", base.Host, base.Scheme)
	cookies := []*http.Cookie{
		// It seems that, for some reason, Steam tries to URL-decode the cookie.
		&http.Cookie{
			Name: "sessionid",
			//Value:  url.QueryEscape(sessionId),
			Value:  sessionId,
			Path:   "/",
			Domain: "steamcommunity.com",
		},
		// steamLogin is already URL-encoded.
		&http.Cookie{
			Name:   "steamLogin",
			Value:  steamLogin,
			Path:   "/",
			Domain: "steamcommunity.com",
		},
	}

	if steamLoginSecure != "" {
		cookies = append(cookies, &http.Cookie{
			Name:   "steamLoginSecure",
			Value:  steamLoginSecure,
			Path:   "/",
			Domain: "steamcommunity.com",
		})
	}
	client.Jar.SetCookies(base, cookies)
}

func SetCookiesHttps(client *http.Client, sessionId, steamLogin, steamLoginSecure string) {
	if client.Jar == nil {
		client.Jar, _ = cookiejar.New(nil)
	}

	base, err := url.Parse("https://steamcommunity.com")
	if err != nil {
		panic(err)
	}

	dddd := client.Jar.Cookies(base)
	fmt.Printf("base.host=%s,scheme=%v,len(dddd)=%d\n", base.Host, base.Scheme, len(dddd))
	client.Jar.SetCookies(base, []*http.Cookie{
		// It seems that, for some reason, Steam tries to URL-decode the cookie.
		&http.Cookie{
			Name: "sessionid",
			//Value:  url.QueryEscape(sessionId),
			Value:  sessionId,
			Path:   "/",
			Domain: "steamcommunity.com",
			Secure: true,
		},
		// steamLogin is already URL-encoded.
		&http.Cookie{
			Name:   "steamLogin",
			Value:  steamLogin,
			Path:   "/",
			Domain: "steamcommunity.com",
			Secure: true,
		},
		&http.Cookie{
			Name:   "steamLoginSecure",
			Value:  steamLoginSecure,
			Path:   "/",
			Domain: "steamcommunity.com",
			Secure: true,
		},
	})
}
