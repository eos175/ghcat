package main

import (
	"errors"
	"net/url"
	"path"
	"strings"
)

func parseInput(input string) (owner, repo, branch, filePath string, err error) {
	if strings.HasPrefix(input, "http") {
		return parseURL(input)
	}

	parts := strings.SplitN(input, "/", 3)
	if len(parts) < 3 {
		return "", "", "", "", errors.New("invalid format, expected owner/repo/path")
	}

	return parts[0], parts[1], "", parts[2], nil
}

func parseURL(raw string) (owner, repo, branch, filePath string, err error) {
	u, err := url.Parse(raw)
	if err != nil {
		return
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 5 {
		err = errors.New("invalid GitHub URL")
		return
	}

	owner = parts[0]
	repo = parts[1]

	if parts[2] != "blob" {
		err = errors.New("unsupported GitHub URL")
		return
	}

	branch = parts[3]
	filePath = path.Join(parts[4:]...)
	return
}
