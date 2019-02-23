package overleaf

import (
	"strings"
)

// parseProjectID returns projectID from given project url
func parseProjectID(url string) string {
	url = strings.TrimPrefix(url, "https://www.overleaf.com/")
	if strings.HasPrefix(url, "project/") {
		// url with form of https://www.overleaf.com/project/PROJECT_ID
		return strings.TrimPrefix(url, "project/")
	}
	if strings.HasPrefix(url, "read/") {
		// url with form of https://www.overleaf.com/read/PROJECT_ID
		return strings.TrimPrefix(url, "read/")
	}
	// url with form of https://www.overleaf.com/PROJECT_ID
	return url
}
