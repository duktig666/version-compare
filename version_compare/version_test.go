package version_compare

import (
	"testing"
)

func TestCompareVersions(t *testing.T) {
	testCases := []struct {
		v1       string
		v2       string
		expected int
	}{
		// 新增的测试用例
		{"v5.3", "v5.3.1", -1},
		{"v5.3.0", "v5.2.9", 1},
		{"2.2.1", "2.2.2", -1},
		{"2.2.1", "2.2.0", 1},
		{"GreatVoyage-v4.7.6", "GreatVoyage-v4.7.7", -1},
		{"GreatVoyage-v4.7.6", "GreatVoyage-v4.7.5", 1},
		{"v1.1", "v1.2", -1},
		{"v1.1", "v1.0", 1},
		{"0.12.0.1", "0.12.0.2", -1},
		{"0.12.0.1", "0.12.0.0", 1},
		{"v2024.09", "v2024.10", -1},
		{"v2024.09", "v2024.08", 1},
		{"v1.1.3-rc.5", "v1.1.4", -1},
		{"v1.1.3-rc.5", "v1.1.2", 1},
		{"Haha-v4.7.8", "GreatVoyage-v4.7.7", 1},
		{"v2.2.1", "2.2.2", -1},
	}

	for _, tc := range testCases {
		result := CompareVersions(tc.v1, tc.v2)
		if result != tc.expected {
			t.Errorf("CompareVersions(%s, %s) = %d; want %d", tc.v1, tc.v2, result, tc.expected)
		}
	}
}
