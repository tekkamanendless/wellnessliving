package wellnessliving

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

func computeAuthorizationHash(signature Signature) string {
	return "20150518," + signature.AuthorizationID + ",," + signatureCompute(signature)
}

func signatureCompute(signature Signature) string {
	var parts []string
	parts = append(parts, "Core\\Request\\Api::20150518")

	parts = append(parts, signature.Time.Format("2006-01-02 15:04:05"))
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
		for key := range signature.Variables {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := signature.Variables.Get(key)
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

	input := strings.Join(parts, "\n")
	logrus.Debugf("INPUT: [\n%s\n]\n", input)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}
