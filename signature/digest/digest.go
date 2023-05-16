package digest

import (
	"demogog/core"
	"net/http"
	"strings"
)

// Digest will compute the digest of your payload and add a new entry to
// the headers you're passing. Old values are replaced.
//
// The digest header is: "Digest": "{hashingAlgorithmName}=base64(hashingAlgorithm(payload))"
func Digest(hashingAlgorithm core.HashingAlgorithm, payload string, header *http.Header) {
	digest := strings.ToUpper(string(hashingAlgorithm)) + "="

	switch hashingAlgorithm {
	case core.SHA256:
		digest += core.Base64(core.Sha256([]byte(payload)))
	default:
		digest += core.Base64(core.Sha512([]byte(payload)))
	}

	header.Add("Digest", digest)
}
