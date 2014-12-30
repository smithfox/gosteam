package community

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func SetCookies(client *http.Client, sessionId, steamLogin, steamLoginSecure string) {
	if client.Jar == nil {
		client.Jar, _ = cookiejar.New(new(cookiejar.Options))
	}
	base, err := url.Parse("steamcommunity.com")
	if err != nil {
		panic(err)
	}
	cookies := []*http.Cookie{
		// It seems that, for some reason, Steam tries to URL-decode the cookie.
		&http.Cookie{
			Name:  "sessionid",
			Value: url.QueryEscape(sessionId),
		},
		// steamLogin is already URL-encoded.
		&http.Cookie{
			Name:  "steamLogin",
			Value: url.QueryEscape(steamLogin),
		},
	}

	if steamLoginSecure != "" {
		cookies = append(cookies, &http.Cookie{
			Name:  "steamLoginSecure",
			Value: url.QueryEscape(steamLoginSecure),
		})
	}
	client.Jar.SetCookies(base, cookies)
}

func SetCookiesHttps(client *http.Client, sessionId, steamLogin, steamLoginSecure string) {
	if client.Jar == nil {
		client.Jar, _ = cookiejar.New(new(cookiejar.Options))
	}

	base, err := url.Parse("steamcommunity.com")
	if err != nil {
		panic(err)
	}

	client.Jar.SetCookies(base, []*http.Cookie{
		// It seems that, for some reason, Steam tries to URL-decode the cookie.
		&http.Cookie{
			Name:   "sessionid",
			Value:  url.QueryEscape(sessionId),
			Secure: true,
		},
		// steamLogin is already URL-encoded.
		&http.Cookie{
			Name:   "steamLogin",
			Value:  url.QueryEscape(steamLogin),
			Secure: true,
		},
		&http.Cookie{
			Name:   "steamLoginSecure",
			Value:  url.QueryEscape(steamLoginSecure),
			Secure: true,
		},
	})
}
