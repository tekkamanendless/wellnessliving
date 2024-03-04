package wellnessliving

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tekkamanendless/httperror"
	"golang.org/x/crypto/sha3"
)

// Client is the WellnessLiving client.
type Client struct {
	URL               string      // The base URL.  If empty, this will use the WellnessLiving production URL.
	AuthorizationCode string      // This is your authorization code.
	AuthorizationID   string      // This is your authorization ID.
	HTTPClient        http.Client // This is the HTTP client.  It's available in case you need to make tweaks.
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

// Login using the given username and password.
//
// Logging in is not required for all API requests, but it is for some.
func (c *Client) Login(ctx context.Context, username string, password string) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("could not create cookie jar: %w", err)
	}
	c.HTTPClient.Jar = jar

	var notepadResponse NotepadResponse
	err = c.Request(ctx, http.MethodGet, "/Core/Passport/Login/Enter/Notepad.json", nil, nil, &notepadResponse)
	if err != nil {
		return err
	}

	var hashedPassword string
	switch notepadResponse.Hash {
	case "sha3":
		parts := []string{
			"r",
			"4S",
			"zqX",
			"zqiOK",
			"TLVS75V",
			"Ue5aLaIIG75",
			"uODJYM2JsCX4G",
			"kt58wZfHHGQkHW4QN",
			"Lh9Fl5989crMU4E7P6E",
		}
		hashedPassword = fmt.Sprintf("%x", sha3.Sum512([]byte(notepadResponse.Notepad+fmt.Sprintf("%x", sha3.Sum512([]byte(strings.Join(parts, password)+password))))))
	}

	enterInput := url.Values{}
	enterInput.Set("s_captcha", "")
	enterInput.Set("s_login", username)
	enterInput.Set("s_notepad", notepadResponse.Notepad)
	enterInput.Set("s_password", hashedPassword)
	enterInput.Set("s_remember", "")
	var enterResponse EnterResponse
	err = c.Request(ctx, http.MethodPost, "/Core/Passport/Login/Enter/Enter.json", enterInput, nil, &enterResponse)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Request(ctx context.Context, method string, path string, variables url.Values, input interface{}, output interface{}) error {
	var bodyString string
	header := http.Header{}
	if input == nil {
		bodyString = variables.Encode()
		header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		if v, ok := input.(string); ok {
			bodyString = v
		} else if v, ok := input.(url.Values); ok {
			bodyString = v.Encode()
			header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}

	contents, err := c.Raw(ctx, method, path, variables, bodyString, header)
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

func (c *Client) Raw(ctx context.Context, method string, path string, variables url.Values, bodyString string, header http.Header) ([]byte, error) {
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

	method = strings.ToUpper(method)

	var body io.Reader
	if bodyString != "" {
		body = strings.NewReader(bodyString)
	}

	request, err := http.NewRequest(method, targetURL, body)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not create request: %w", err)
	}
	for key, values := range header {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	tz, err := time.LoadLocation("GMT")
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not load timezone: %w", err)
	}
	now := time.Now().In(tz)

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
		CookiePersistent:  "", // Default these to empty for now.
		CookieTransient:   "", // Default these to empty for now.
		Host:              myURL.Host,
		AuthorizationID:   authorizationID,
		Method:            strings.ToUpper(method),
		Resource:          strings.TrimLeft(path, "/"),
	}
	if c.HTTPClient.Jar != nil {
		cookieURL, err := url.Parse(targetURL)
		if err != nil {
			return nil, fmt.Errorf("could not parse URL for cookies: %w", err)
		}
		cookies := c.HTTPClient.Jar.Cookies(cookieURL)
		for _, cookie := range cookies {
			switch cookie.Name {
			case "p":
				signature.CookiePersistent = cookie.Value
			case "t":
				signature.CookieTransient = cookie.Value
			}
		}
	}
	authorization := computeAuthorizationHash(signature)

	request.Header.Set("Accept", "*/*")
	request.Header.Set("Date", now.Format(time.RFC1123))
	request.Header.Set("User-Agent", "WellnessLiving SDK/1.1 (WellnessLiving SDK)")
	request.Header.Set("Authorization", authorization)
	if bodyString != "" && bodyString[0] == '{' {
		request.Header.Set("Content-Type", "application/json")
	}

	{
		contents, _ := httputil.DumpRequest(request, true)
		logrus.WithContext(ctx).Debugf("REQUEST:\n%s\n", contents)
	}

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not perform request: %w", err)
	}

	logrus.WithContext(ctx).Debugf("Status code: %d", response.StatusCode)
	if response.StatusCode >= 400 {
		return nil, httperror.ErrorFromStatus(response.StatusCode)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not read response body: %w", err)
	}
	logrus.WithContext(ctx).Debugf("Response:")
	logrus.WithContext(ctx).Debugf("%s", contents)

	return contents, nil
}
