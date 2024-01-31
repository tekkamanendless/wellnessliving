package wellnessliving

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

type Client struct {
	URL               string
	AuthorizationCode string
	AuthorizationID   string
	HTTPClient        http.Client
}

type Signature struct {
	Header            http.Header
	Variable          url.Values
	Time              time.Time
	AuthorizationCode string
	CookiePersistent  string
	CookieTransient   string
	Host              string
	AuthorizationID   string
	Method            string
	Resource          string
}

func (c *Client) Raw(ctx context.Context, method string, path string) error {
	baseURL := c.URL
	if baseURL == "" {
		baseURL = "https://us.wellnessliving.com"
	}

	targetURL := path
	if !strings.Contains(targetURL, "://") {
		targetURL = strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(path, "/")
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

	authorizationCode := c.AuthorizationCode
	if authorizationCode != "" {
		authorizationCode = os.Getenv("WELLNESSLIVING_AUTHORIZATION_CODE")
	}
	authorizationID := c.AuthorizationID
	if authorizationID != "" {
		authorizationID = os.Getenv("WELLNESSLIVING_AUTHORIZATION_CODE")
	}

	signature := Signature{
		Header:            http.Header{},
		Variable:          url.Values{},
		Time:              now,
		AuthorizationCode: authorizationCode,
		CookiePersistent:  "", // TODO
		CookieTransient:   "", // TODO
		Host:              myURL.Host,
		AuthorizationID:   authorizationID,
		Method:            method,
		Resource:          "", // TODO
	}
	authorization := computeAuthorizationHash(signature)

	request.Header.Set("Date", now.Format(time.RFC1123))
	request.Header.Set("User-Agent", "WellnessLiving SDK/1.1 (WellnessLiving SDK)")
	request.Header.Set("Authorization", authorization)

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Printf("RESPONSE:\n%s\n", contents)

	return nil
}

func computeAuthorizationHash(signature Signature) string {
	return "20150518," + signature.AuthorizationID + ",," + signatureCompute(signature)
}

func signatureCompute(signature Signature) string {
	var parts []string
	parts = append(parts, "Core\\Request\\Api::20150518")

	parts = append(parts, signature.Time.Format(time.RFC1123))
	parts = append(parts, signature.AuthorizationCode)
	parts = append(parts, signature.Host)
	parts = append(parts, signature.AuthorizationID)
	parts = append(parts, signature.Method)
	parts = append(parts, signature.Resource)
	parts = append(parts, signature.CookiePersistent)
	parts = append(parts, signature.CookieTransient)

	// TODO: Maybe signature.Variable should be more abstract.
	//$a_variable=WlModelRequest::signatureArray($a_data['a_variable'])
	{
		var keys []string
		for key := range signature.Variable {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := signature.Variable.Get(key)
			// TODO: If the value is null, use "${key} is null"
			parts = append(parts, key+"="+value)
		}
	}

	{
		var keys []string
		for key := range signature.Header {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := signature.Header.Get(key)
			parts = append(parts, strings.ToLower(key)+":"+strings.TrimSpace(value))
		}
	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(parts, "\n"))))
}
