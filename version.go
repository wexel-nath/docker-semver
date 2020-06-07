package main

import (
	"fmt"
	"strconv"
	"strings"
)

type version struct {
	major int
	minor int
	patch int
}

func (v version) Major() string {
	return fmt.Sprintf("%d", v.major)
}

func (v version) Minor() string {
	return fmt.Sprintf("%d.%d", v.major, v.minor)
}

func (v version) Patch() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func (v version) isHigherThan(v2 version) bool {
	if v.major == v2.major {
		if v.minor == v2.minor {
			return v.patch >= v2.patch
		}
		return v.minor >= v2.minor
	}
	return v.major >= v2.major
}

func (v version) increment(label string) version {
	label = strings.ToUpper(label)
	switch label {
	case labelMajor:
		return version{
			major: v.major + 1,
			minor: 0,
			patch: 0,
		}
	case labelMinor:
		return version{
			major: v.major,
			minor: v.minor + 1,
			patch: 0,
		}
	case labelPatch:
		fallthrough
	default:
		return version{
			major: v.major,
			minor: v.minor,
			patch: v.patch + 1,
		}
	}
}

func newVersion(tag string) (version, error) {
	tagParts := strings.Split(tag, ".")
	if len(tagParts) != 3 {
		return version{}, fmt.Errorf("expected 3 parts, got %d for tag: %s", len(tagParts), tag)
	}

	major, err := strconv.Atoi(tagParts[0])
	if err != nil {
		return version{}, fmt.Errorf("could not parse major version from tag: %s", tag)
	}

	minor, err := strconv.Atoi(tagParts[1])
	if err != nil {
		return version{}, fmt.Errorf("could not parse minor version from tag: %s", tag)
	}

	patch, err := strconv.Atoi(tagParts[2])
	if err != nil {
		return version{}, fmt.Errorf("could not parse patch version from tag: %s", tag)
	}

	return version{
		major: major,
		minor: minor,
		patch: patch,
	}, nil
}
