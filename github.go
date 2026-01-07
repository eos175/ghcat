package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v61/github"
	"golang.org/x/oauth2"
)

func loadGitHubToken() string {
	if token := strings.TrimSpace(os.Getenv("GITHUB_TOKEN")); token != "" {
		return token
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	data, err := os.ReadFile(filepath.Join(home, ".github_token"))
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func githubClient(ctx context.Context) *github.Client {
	token := loadGitHubToken()
	if token == "" {
		return github.NewClient(nil)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	return github.NewClient(oauth2.NewClient(ctx, ts))
}
