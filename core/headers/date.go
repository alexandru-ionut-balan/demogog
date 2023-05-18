package headers

import (
	"demogog/core"
	"net/http"
	"time"
)

// Date will add a new header value to your [header] with the current timestamp formatted
// according to RFC7231
func Date(header *http.Header) {
	defer core.Log.Sync()

	location, err := time.LoadLocation("GMT")
	if err != nil {
		core.Log.Error("Cannot find time.Location GMT")
	}

	now := time.Now().In(location).Format(time.RFC850)
}
