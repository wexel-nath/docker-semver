package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	defaultTagsLimit = "100"

	labelMajor = "MAJOR"
	labelMinor = "MINOR"
	labelPatch = "PATCH"
)

func main() {
	gitCommitMessage := os.Getenv("GIT_COMMIT_MSG")
	imageTagsLimit := os.Getenv("IMAGE_TAGS_LIMIT")
	imageTagsUrl := os.Getenv("IMAGE_TAGS_URL")

	if imageTagsUrl == "" || gitCommitMessage == "" {
		fmt.Println("IMAGE_TAGS_URL and GIT_COMMIT_MSG must be defined")
		os.Exit(1)
	}

	if imageTagsLimit == "" {
		imageTagsLimit = defaultTagsLimit
	}

	// get tagsResponse from imageTagsUrl
	tags, err := getTags(imageTagsUrl, imageTagsLimit)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// get label from git commit message
	label := getLabel(gitCommitMessage)

	maxVersion := tags.getMaxVersion()
	newVersion := maxVersion.increment(label)

	fmt.Println(newVersion.Major())
	fmt.Println(newVersion.Minor())
	fmt.Println(newVersion.Patch())
}

func getLabel(commitMsg string) string {
	parts := strings.Split(commitMsg, " ")
	if len(parts) == 0 {
		return ""
	}

	return strings.Replace(parts[0], ":", "", -1)
}
