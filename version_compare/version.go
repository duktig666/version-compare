package version_compare

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	versionRegex = regexp.MustCompile(`(\d+(?:\.\d+)*)`)
)

// CompareVersions 比较两个版本号
// 如果 v1 > v2, 返回 1
// 如果 v1 < v2, 返回 -1
// 如果 v1 == v2, 返回 0
func CompareVersions(v1, v2 string) int {
	v1 = extractVersion(v1)
	v2 = extractVersion(v2)

	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		n1, _ := strconv.Atoi(parts1[i])
		n2, _ := strconv.Atoi(parts2[i])

		if n1 < n2 {
			return -1
		}
		if n1 > n2 {
			return 1
		}
	}

	if len(parts1) < len(parts2) {
		return -1
	}
	if len(parts1) > len(parts2) {
		return 1
	}

	return 0
}

func extractVersion(v string) string {
	matches := versionRegex.FindStringSubmatch(v)
	if len(matches) > 0 {
		return matches[0]
	}
	return v
}
