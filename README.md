# Version Compare

Version Compare is a Go library designed to compare version numbers in various formats. It provides a simple and efficient way to determine the relationship between two version strings, regardless of their prefixes or suffixes.

## Features

- Compares version numbers in different formats (e.g., "v1.2.3", "2.2.1", "GreatVoyage-v4.7.6")
- Ignores version prefixes (such as "v" or "GreatVoyage-v")
- Compares only the numeric parts of the version strings
- Handles versions with different number of components (e.g., "1.2" vs "1.2.3")
- Supports pre-release versions (e.g., "v1.1.3-rc.5")

## Installation

To use this library in your Go project, you can install it using `go get`:

