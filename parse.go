package main

import (
	"errors"
	"net/url"
	"path"
	"strings"
)

// parseInput parses either a full GitHub URL or a shorthand owner/repo/path.
// For shorthand input, branch is left empty so the caller can discover
// the repository default branch.
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

// parseURL parses supported GitHub URLs and normalizes tree paths.
//
// Supported forms:
// - https://github.com/{owner}/{repo}
// - https://github.com/{owner}/{repo}/blob/{ref}/{file}
// - https://github.com/{owner}/{repo}/tree/{ref}/{dir-or-file}
//
// For tree URLs that point to a directory, it resolves to {dir}/README.md.
func parseURL(raw string) (owner, repo, branch, filePath string, err error) {
	u, err := url.Parse(raw)
	if err != nil {
		return
	}

	if !strings.EqualFold(u.Host, "github.com") {
		err = errors.New("unsupported host")
		return
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) == 2 {
		owner = parts[0]
		repo = parts[1]
		filePath = "README.md"
		return
	}

	if len(parts) < 5 {
		err = errors.New("invalid GitHub URL")
		return
	}

	owner = parts[0]
	repo = parts[1]

	if parts[2] != "blob" && parts[2] != "tree" {
		err = errors.New("unsupported GitHub URL")
		return
	}

	branch = parts[3]
	rest := parts[4:]

	if parts[2] == "blob" {
		filePath = path.Join(rest...)
		return
	}

	if len(rest) == 0 {
		filePath = "README.md"
		return
	}

	last := rest[len(rest)-1]
	if strings.Contains(last, ".") {
		filePath = path.Join(rest...)
		return
	}

	filePath = path.Join(path.Join(rest...), "README.md")
	return
}
