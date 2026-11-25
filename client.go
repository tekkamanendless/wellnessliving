package wellnessliving

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/sirupsen/logrus"
	"github.com/tekkamanendless/restapiclient"
	"golang.org/x/crypto/sha3"
)

// Client is the WellnessLiving client.
//
// This MUST be created using NewClient; it is not intended to be instantiated directly.
//
// At miniumum, you will need to set AuthorizationCode and AuthorizationID.
// These values must be obtained by WellnessLiving as part of signing up for their API program.
//
// If not otherwise specified, those values are loaded from the following environment variables:
// * WELLNESSLIVING_AUTHORIZATION_CODE
// * WELLNESSLIVING_AUTHORIZATION_ID
//
// If you wish to use the WellnessLiving staging API, then you will need to set URL, as well.
type Client struct {
	config     Config               // This is the configuration for the client.
	restClient restapiclient.Client // This is the REST client.  It is used to perform the actual requests.
}

// Config is the configuration for the client.
type Config struct {
	BaseURL           string                // The base URL.  If empty, this will use the WellnessLiving production URL.
	AuthorizationCode string                // This is your authorization code.  If not set, the value of WELLNESSLIVING_AUTHORIZATION_CODE will be used.
	AuthorizationID   string                // This is your authorization ID.  If not set, the value of WELLNESSLIVING_AUTHORIZATION_CODE will be used.
	HTTPClient        *http.Client          // This is the HTTP client.  It's available in case you need to make tweaks.
	BackOffOptions    []backoff.RetryOption // This is the set of backoff options.  If not set, the default will be used.
}

// Option is a function that can be used to configure the client.
type Option func(*Config)

// WithAuthorizationCode sets the authorization code for the client.
func WithAuthorizationCode(authorizationCode string) Option {
	return func(config *Config) {
		config.AuthorizationCode = authorizationCode
	}
}

// WithAuthorizationID sets the authorization ID for the client.
func WithAuthorizationID(authorizationID string) Option {
	return func(config *Config) {
		config.AuthorizationID = authorizationID
	}
}

// WithBackOffOptions sets the backoff options for the client.
func WithBackOffOptions(backOffOptions ...backoff.RetryOption) Option {
	return func(config *Config) {
		config.BackOffOptions = append(config.BackOffOptions, backOffOptions...)
	}
}

// WithBaseURL sets the base URL for the client.
func WithBaseURL(baseURL string) Option {
	return func(config *Config) {
		config.BaseURL = baseURL
	}
}

// WithHTTPClient sets the HTTP client for the client.
func WithHTTPClient(httpClient *http.Client) Option {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return func(config *Config) {
		config.HTTPClient = httpClient
	}
}

// NewClient creates a new WellenssLiving client.
func NewClient(options ...Option) *Client {
	config := Config{
		BaseURL:           "https://us.wellnessliving.com",
		AuthorizationCode: os.Getenv("WELLNESSLIVING_AUTHORIZATION_CODE"),
		AuthorizationID:   os.Getenv("WELLNESSLIVING_AUTHORIZATION_ID"),
		HTTPClient:        &http.Client{},
		BackOffOptions:    nil,
	}
	for _, option := range options {
		option(&config)
	}
	return &Client{
		config: config,
		restClient: restapiclient.New(
			restapiclient.OptionBaseURL(config.BaseURL),
			restapiclient.OptionHTTPClient(config.HTTPClient),
			restapiclient.OptionBackOffIsPermanentErrorFunction(restapiclient.DefaultBackOffIsPermanentErrorFunction),
			restapiclient.OptionBackOffOptions(config.BackOffOptions...),
		),
	}
}

// Signature contains all of the pieces of information needed to compute the signature verification
// that is needed for every API request.
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
//
// After logging in, the client will be authenticated for all future requests using the client's
// cookie jar as part of HTTPClient.
func (c *Client) Login(ctx context.Context, username string, password string) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("could not create cookie jar: %w", err)
	}
	c.config.HTTPClient.Jar = jar

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

// Request performs and API request.
//
// variables is what WellnessLiving refers to as such.  These will be used as query parameters.
// For requests that normally expect a body (such as POST), they will be converted to form values
// and the encoding will be set appropriately.
func (c *Client) Request(ctx context.Context, method string, path string, variables url.Values, input any, output any) error {
	var bodyString string
	header := http.Header{}
	if input == nil {
		switch strings.ToUpper(method) {
		case http.MethodGet, http.MethodDelete:
			// Do not attempt to construct a body.
		default:
			// Construct a body for the request.
			bodyString = variables.Encode()
			header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	} else {
		if v, ok := input.(string); ok {
			// The input is a string; use it as-is.
			bodyString = v
		} else if v, ok := input.(url.Values); ok {
			// The input is a collection of values; encode it as a form.
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

// Raw performs a raw request and returns any response content.
//
// If the response is an error response, then the appropriate `httperror` response will be returned.
//
// variables will be used as query parameters.
// bodyString, if not empty, will be used as the body.  Please ensure that the "Content-Type" header is set appropriately.
// header is the set of HTTP headers to send.
//
// In addition to any headers specified, the following headers will be set:
// * Accept
// * Date
// * User-Agent
// * Authorization
func (c *Client) Raw(ctx context.Context, method string, path string, variables url.Values, bodyString string, header http.Header) ([]byte, error) {
	baseURL := c.config.BaseURL

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

	tz, err := time.LoadLocation("GMT")
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not load timezone: %w", err)
	}
	now := time.Now().In(tz)

	signature := Signature{
		Header:            http.Header{},
		Variables:         variables,
		Time:              now,
		AuthorizationCode: c.config.AuthorizationCode,
		CookiePersistent:  "", // Default these to empty for now.
		CookieTransient:   "", // Default these to empty for now.
		Host:              myURL.Host,
		AuthorizationID:   c.config.AuthorizationID,
		Method:            strings.ToUpper(method),
		Resource:          strings.TrimLeft(path, "/"),
	}
	if c.config.HTTPClient.Jar != nil {
		cookieURL, err := url.Parse(targetURL)
		if err != nil {
			return nil, fmt.Errorf("could not parse URL for cookies: %w", err)
		}
		cookies := c.config.HTTPClient.Jar.Cookies(cookieURL)
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

	header.Set("Accept", "*/*")
	header.Set("Date", now.Format(time.RFC1123))
	header.Set("User-Agent", "WellnessLiving SDK/1.1 (WellnessLiving SDK)")
	header.Set("Authorization", authorization)
	if bodyString != "" && bodyString[0] == '{' {
		header.Set("Content-Type", "application/json")
	}

	var requestOptions []restapiclient.Option
	for key, values := range header {
		for _, value := range values {
			requestOptions = append(requestOptions, restapiclient.OptionHeader(key, value))
		}
	}

	var outputBytes restapiclient.RawBytes
	err = c.restClient.Do(ctx, method, targetURL, bodyString, &outputBytes, requestOptions...)
	if err != nil {
		return nil, fmt.Errorf("wellnessliving: could not perform request: %w", err)
	}

	logrus.WithContext(ctx).Debugf("Response:")
	logrus.WithContext(ctx).Debugf("%s", outputBytes)

	return outputBytes, nil
}
