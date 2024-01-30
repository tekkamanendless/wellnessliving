package wellnessliving

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	URL string
	AuthorizationCode string
	AuthorizationID string
}

type Signature struct {
		Header http.Header
		Variable url.Values
		Time time.Time
		AuthorizationCode string
		CookiePersistent string
		CookieTransient string
		Host string
		AuthorizationID string
		Method string
		Resource string  
}

func (c *Client) Raw(ctx context.Context, method string, path string) error {
	targetURL := path
	if !strings.Contains(targetURL, "://") {
		targetURL = strings.TrimRight(c.URL, "/") + "/" + strings.TrimLeft(path, "/")
	}

	myURL, err := url.Parse(targetURL)
	if err != nil {
		return err
	}

	var body io.Reader

	request, err := http.NewRequest(method, targetURL, body)
	if err != nil {
		return err
	}

	now := time.Now()

	signature := Signature {
			Header: http.Header{},
			Variable: url.Values{},
			Time: now,
			AuthorizationCode: c.AuthorizationCode,
			CookiePersistent: "",
			CookieTransient: ""
			Host: myURL.Host,
			AuthorizationID: c.AuthorizationID,
			Method: method,
			Resource: "",
	}
	authorization, err := computeAuthorizationHash(signature)
	if err != nil {
		return err
	}

	request.Header.Set("Date", now.Format(time.RFC1123))
	request.Header.Set("User-Agent", "WellnessLiving SDK/1.1 (WellnessLiving SDK)")
	request.Header.Set("Authorization", authorization)
}

func computeAuthorizationHash(signature Signature) (string, error) {
	return "20150518," + signature.AuthorizationID.',,'.WlModelRequest::signatureCompute($a_signature);
  }

}
