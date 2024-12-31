package main

import (
	"net/url"
	"strings"
)

func filterLinks(links []string, currentURL string) []string {
	current, err := url.Parse(currentURL)
	if err != nil {
		return nil
	}
	links = uniqueLinks(links)
	var filteredLinks []string
	for _, link := range links {
		linkURL, err := url.Parse(link)
		if err != nil {
			continue
		}
		if linkURL.Host == current.Host && isChildRoute(linkURL, current) {
			filteredLinks = append(filteredLinks, link)
		}
	}

	return filteredLinks
}

func isChildRoute(child, parent *url.URL) bool {
	if child.Path == parent.Path {
		return false
	}
	return strings.Contains(child.Path, parent.Path)
}

func uniqueLinks(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
