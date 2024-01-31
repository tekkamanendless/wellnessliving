package wellnessliving

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Client struct {
	URL               string
	AuthorizationCode string
	AuthorizationID   string
	HTTPClient        http.Client
}

type Signature struct {
	Header            http.Header
	Variables         url.Values
	Time              time.Time
	AuthorizationCode string
	CookiePersistent  string
	CookieTransient   string
	Host              string
	AuthorizationID   string
	Method            string
	Resource          string
}

func (c *Client) Request(ctx context.Context, method string, path string, variables url.Values, output interface{}) error {
	contents, err := c.Raw(ctx, method, path, variables)
	if err != nil {
		return err
	}

	var baseResponse BaseResponse
	err = json.Unmarshal(contents, &baseResponse)
	if err != nil {
		return fmt.Errorf("wellnessliving: could not parse response envelope: %w", err)
	}

	logrus.WithContext(ctx).Debugf("Envelope: %+v", baseResponse)
	if baseResponse.Status != "ok" {
		var errorResponse ErrorResponse
		err = json.Unmarshal(contents, &errorResponse)
		if err != nil {
			return fmt.Errorf("wellnessliving: could not parse error response: %w", err)
		}
		return &errorResponse
	}

	if output != nil {
		err = json.Unmarshal(contents, output)
		if err != nil {
			return fmt.Errorf("wellnessliving: could not parse response: %w", err)
		}
	}

	return nil
}

func (c *Client) Raw(ctx context.Context, method string, path string, variables url.Values) ([]byte, error) {
	baseURL := c.URL
	if baseURL == "" {
		baseURL = "https://us.wellnessliving.com"
	}

	targetURL := path
	if !strings.Contains(targetURL, "://") {
		targetURL = strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(path, "/")
	}
	if len(variables) > 0 {
		targetURL += "?" + variables.Encode()
	}

	myURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not parse URL: %w", err)
	}

	var body io.Reader

	request, err := http.NewRequest(method, targetURL, body)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not create request: %w", err)
	}

	tz, err := time.LoadLocation("GMT")
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not load timezone: %w", err)
	}
	now := time.Now().In(tz)
	//now, _ = time.ParseInLocation("2006-01-02 15:04:05", "2024-01-31 11:58:28", tz)

	authorizationCode := c.AuthorizationCode
	if authorizationCode == "" {
		authorizationCode = os.Getenv("WELLNESSLIVING_AUTHORIZATION_CODE")
	}
	authorizationID := c.AuthorizationID
	if authorizationID == "" {
		authorizationID = os.Getenv("WELLNESSLIVING_AUTHORIZATION_ID")
	}

	signature := Signature{
		Header:            http.Header{},
		Variables:         variables,
		Time:              now,
		AuthorizationCode: authorizationCode,
		CookiePersistent:  "", // TODO
		CookieTransient:   "", // TODO
		Host:              myURL.Host,
		AuthorizationID:   authorizationID,
		Method:            strings.ToUpper(method),
		Resource:          strings.TrimLeft(path, "/"),
	}
	authorization := computeAuthorizationHash(signature)

	request.Header.Set("Date", now.Format(time.RFC1123))
	request.Header.Set("User-Agent", "WellnessLiving SDK/1.1 (WellnessLiving SDK)")
	request.Header.Set("Authorization", authorization)

	{
		contents, _ := httputil.DumpRequest(request, true)
		logrus.WithContext(ctx).Debugf("REQUEST:\n%s\n", contents)
	}

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not perform request: %w", err)
	}

	logrus.WithContext(ctx).Debugf("Status code: %d", response.StatusCode)

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not read response body: %w", err)
	}
	logrus.WithContext(ctx).Debugf("Response:")
	logrus.WithContext(ctx).Debugf("%s", contents)

	return contents, nil
}
