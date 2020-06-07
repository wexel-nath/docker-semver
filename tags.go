package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type tag struct {
	Name string `json:"name"`
}

type tagsResponse struct {
	Tags []tag `json:"results"`
}

func getTags(imageTagsUrl string, imageTagsLimit string) (tagsResponse, error) {
	request, err := http.NewRequest(http.MethodGet, imageTagsUrl, nil)
	if err != nil {
		return tagsResponse{}, err
	}

	q := request.URL.Query()
	q.Set("page_size", imageTagsLimit)
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return tagsResponse{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tagsResponse{}, err
	}

	var tags tagsResponse
	err = json.Unmarshal(body, &tags)
	return tags, err
}

func (response tagsResponse) getMaxVersion() version {
	maxVersion := version{
		major: 1,
		minor: 0,
		patch: 0,
	}

	for _, tag := range response.Tags {
		version, err := newVersion(tag.Name)
		if err == nil && version.isHigherThan(maxVersion) {
			maxVersion = version
		}
	}

	return maxVersion
}
