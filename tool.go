package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
	"os"
	"strings"
	"text/tabwriter"
)

type TopicTool struct {
	client *github.Client
}

func NewTopicTool(token string) *TopicTool {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return &TopicTool{
		client: client,
	}
}

func (t *TopicTool) ReplaceTopics(query string, topics []string) error {
	result, err := t.searchRepositories(query)
	if err != nil {
		return err
	}

	if len(result.Repositories) == 0 {
		return fmt.Errorf("no repositories found matching query %s", query)
	}

	previewRepositories(result.Repositories)

	fmt.Println("")
	fmt.Printf("\n Replace all labels with [%s] in %d repositories? [y/n/q]:\n", topicsToStr(topics), len(result.Repositories))
	next, err := AskForBool(os.Stdin, true, false)
	if err != nil {
		return err
	}
	if !next {
		return ErrAbortInput
	}

	return t.replaceTopics(result.Repositories, topics)
}

func (t *TopicTool) replaceTopics(repositories []*github.Repository, topics []string) error {
	for _, repo := range repositories {
		return t.setRepositoryTopics(repo, topics)
	}
	return nil
}

func (t *TopicTool) AddTopics(query string, topics []string) error {
	result, err := t.searchRepositories(query)
	if err != nil {
		return err
	}

	if len(result.Repositories) == 0 {
		return fmt.Errorf("no repositories found matching query %s", query)
	}

	previewRepositories(result.Repositories)

	fmt.Println("")
	fmt.Printf("\n Add labels [%s] to %d repositories? [y/n/q]:\n", topicsToStr(topics), len(result.Repositories))
	next, err := AskForBool(os.Stdin, true, false)
	if err != nil {
		print(err)
	}
	if !next {
		return ErrAbortInput
	}

	return t.addTopics(result.Repositories, topics)
}

func (t *TopicTool) addTopics(repositories []*github.Repository, topicsToAppend []string) error {
	for _, repo := range repositories {
		repo.Topics = append(repo.Topics, topicsToAppend...)
		return t.setRepositoryTopics(repo, repo.Topics)
	}
	return nil
}

func (t *TopicTool) RemoveTopics(query string, topics []string) error {
	result, err := t.searchRepositories(query)
	if err != nil {
		return err
	}

	if len(result.Repositories) == 0 {
		return fmt.Errorf("no repositories found matching query %s", query)
	}

	previewRepositories(result.Repositories)

	fmt.Println("")
	fmt.Printf("\n Remove labels [%s] from %d repositories? [y/n/q]:\n", topicsToStr(topics), len(result.Repositories))
	next, err := AskForBool(os.Stdin, true, false)
	if err != nil {
		print(err)
	}
	if !next {
		return ErrAbortInput
	}

	return t.removeTopics(result.Repositories, topics)
}

func (t *TopicTool) removeTopics(repositories []*github.Repository, topicsToRemove []string) error {
	for _, repo := range repositories {
		repo.Topics = removeFromTopics(repo.Topics, topicsToRemove)
		return t.setRepositoryTopics(repo, repo.Topics)
	}
	return nil
}

func (t *TopicTool) setRepositoryTopics(repository *github.Repository, topics []string) error {
	topics = removeDuplicateTopics(topics)
	_, _, err := t.client.Repositories.ReplaceAllTopics(context.Background(), *repository.Owner.Login, *repository.Name, topics)
	if err != nil {
		return err
	}
	return nil
}

// searchRepositories searches repositories via various criteria.
//
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#search-repositories
func (t *TopicTool) searchRepositories(query string) (*github.RepositoriesSearchResult, error) {
	result, _, err := t.client.Search.Repositories(context.Background(), query, &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: 100,
		},
	})
	if err != nil {
		return &github.RepositoriesSearchResult{}, err
	}
	return result, err
}

func removeFromTopics(topics []string, toRemove []string) []string {
	// todo(cbrgm): we can use a set here because order does not matter
	set := make(map[string]bool)
	for _, s := range topics {
		set[s] = true
	}

	for _, s := range toRemove {
		delete(set, s)
	}

	res := []string{}
	for k, _ := range set {
		res = append(res, k)
	}
	return res
}

func removeDuplicateTopics(topics []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range topics {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func topicsToStr(topics []string) string {
	res := strings.Join(topics, ",")
	return strings.TrimRight(res, ",")
}

func previewRepositories(repos []*github.Repository) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t", "Repository Name", "Topics")
	fmt.Fprintf(w, "\n %s\t%s\t", "---", "---")

	for _, repo := range repos {
		fmt.Fprintf(w, "\n %s\t%s\t", fmt.Sprintf("%s/%s", *repo.Owner.Login, *repo.Name), topicsToStr(repo.Topics))
	}
}
