package main

import (
	"fmt"

	"github.com/duktig666/version-compare/version_compare"
)

func main() {
	v1 := "v5.3.0"
	v2 := "2.2.1"
	result := version_compare.CompareVersions(v1, v2)
	fmt.Printf("Comparing %s and %s: %d\n", v1, v2, result)
}
