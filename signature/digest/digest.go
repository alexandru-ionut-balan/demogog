package digest

import (
	"demogog/core"
	"net/http"
)

// Digest will compute the digest of your payload and add a new entry to
// the headers you're passing. Old values are replaced.
//
// The digest header is: "Digest": "{hashingAlgorithmName}=base64(hashingAlgorithm(payload))"
func Digest(hashingAlgorithm core.HashingAlgorithm, payload string, header *http.Header) {

}
