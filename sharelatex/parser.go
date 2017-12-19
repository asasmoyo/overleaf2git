package sharelatex

import (
	"strings"
)

// parseProjectID returns projectID, a bool wether the url is shared project
// and another bool wether authentication needed
func parseProjectID(url string) (string, bool, bool) {
	url = strings.TrimPrefix(url, "https://www.sharelatex.com/")
	if strings.HasPrefix(url, "project/") {
		// url with form of https://www.sharelatex.com/project/PROJECT_ID
		return strings.TrimPrefix(url, "project/"), false, true
	}
	if strings.HasPrefix(url, "read/") {
		// url with form of https://www.sharelatex.com/read/PROJECT_ID
		return strings.TrimPrefix(url, "read/"), true, false
	}
	// url with form of https://www.sharelatex.com/PROJECT_ID
	return url, true, true
}
