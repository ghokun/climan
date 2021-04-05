package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v34/github"
)

type Climan struct {
}

func (Climan) Name() string {
	return "climan"
}

func (Climan) Type() string {
	return string(APIClient)
}

func (Climan) Latest() string {
	release, _, _ := github.NewClient(nil).Repositories.GetLatestRelease(context.Background(), "argoproj", "argo-cd")
	return *release.Name
}

func (Climan) Current() string {
	return "-"
}

func (Climan) Description() string {
	return "Command LIne tool version MANager "
}

func (Climan) List() (items []string) {
	tags, _, _ := github.NewClient(nil).Repositories.ListReleases(context.Background(), "argoproj", "argo-cd", &github.ListOptions{Page: 0, PerPage: 10000})
	for _, tag := range tags {
		if !strings.Contains(*tag.Name, "-rc") && !strings.Contains(*tag.Name, "-beta") && !strings.Contains(*tag.Name, "-alpha") {
			items = append(items, *tag.Name)
		}
	}
	return items
}

func (Climan) Install() error {
	// Download
	// Check SHA
	// Unpack
	fmt.Println("Climan install is called")
	return nil
}

func (Climan) Remove() error {
	return nil
}
