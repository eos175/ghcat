package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "ghcat <github-url | owner/repo/path>",
		Short: "Fetch and print a single file from GitHub",
		Args:  cobra.ExactArgs(1),
		RunE:  run,
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	input := args[0]

	// https://gist.githubusercontent.com
	// https://raw.githubusercontent.com
	if strings.Contains(input, "githubusercontent.com") {
		return fetchRaw(input)
	}

	// https://gist.github.com/owner/id
	if strings.Contains(input, "gist.github.com") {
		if !strings.HasSuffix(input, "/raw") {
			input = strings.TrimSuffix(input, "/") + "/raw"
		}
		return fetchRaw(input)
	}

	owner, repo, branch, filePath, err := parseInput(input)
	if err != nil {
		return err
	}

	ctx := context.Background()
	client := githubClient(ctx)

	// Discover default branch if not provided
	if branch == "" {
		repository, _, err := client.Repositories.Get(ctx, owner, repo)
		if err != nil {
			return err
		}
		branch = repository.GetDefaultBranch()
	}

	file, _, _, err := client.Repositories.GetContents(
		ctx,
		owner,
		repo,
		filePath,
		&github.RepositoryContentGetOptions{Ref: branch},
	)
	if err != nil {
		return err
	}

	if file == nil || file.Content == nil {
		return errors.New("not a file")
	}

	data, err := base64.StdEncoding.DecodeString(
		strings.ReplaceAll(*file.Content, "\n", ""),
	)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(data)
	return err
}

func fetchRaw(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("raw fetch failed: %s", resp.Status)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	return err
}
