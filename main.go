package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cmdReplace = flag.NewFlagSet("replace", flag.ExitOnError)
	cmdAdd     = flag.NewFlagSet("add", flag.ExitOnError)
	cmdRemove  = flag.NewFlagSet("rm", flag.ExitOnError)
)

const GithubTokenEnvVar = "GH_ACCESS_TOKEN"

var usage = `Usage: topictool <subcommand> <search pattern> <topic labels...>

Replace, add or remove topic labels from multiple Github repositories

Subcommands:
    - replace   - replaces all existing topic labels with new ones
    - add       - adds topic labels to existing ones
    - rm        - removes topic labels from existing ones
    
Search Pattern:
    Searches repositories via various criteria.
    See Github docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#search-repositories

Topic Labels:
    A list of strings representing topic labels
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		usageAndExit("Please enter a subcommand")
		return
	}

	switch args[0] {
	case "replace":
		cmdReplace.Parse(args[1:])
		replaceCmd()
	case "add":
		cmdAdd.Parse(args[1:])
		addCmd()
	case "rm":
		cmdRemove.Parse(args[1:])
		removeCmd()
	default:
		usageAndExit("Please enter a subcommand")
	}
}

func replaceCmd() {
	token := os.Getenv(GithubTokenEnvVar)
	if token == "" {
		usageAndExit(fmt.Sprintf("please provide a personal access token via %s env var", GithubTokenEnvVar))
	}

	tool := NewTopicTool(token)

	args := cmdReplace.Args()
	if len(args) < 2 {
		usageAndExit("")
	}

	query := args[0]
	topics := args[1:]

	err := tool.ReplaceTopics(query, topics)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done!")
	os.Exit(0)
}

func addCmd() {
	token := os.Getenv(GithubTokenEnvVar)
	if token == "" {
		usageAndExit(fmt.Sprintf("please provide a personal access token via %s env var", GithubTokenEnvVar))
	}

	tool := NewTopicTool(token)

	args := cmdAdd.Args()
	if len(args) < 2 {
		usageAndExit("")
	}

	query := args[0]
	topics := args[1:]

	err := tool.AddTopics(query, topics)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done!")
	os.Exit(0)
}

func removeCmd() {
	token := os.Getenv(GithubTokenEnvVar)
	if token == "" {
		usageAndExit(fmt.Sprintf("please provide a personal access token via %s env var", GithubTokenEnvVar))
	}

	tool := NewTopicTool(token)

	args := cmdRemove.Args()
	if len(args) < 2 {
		usageAndExit("")
	}

	query := args[0]
	topics := args[1:]

	err := tool.RemoveTopics(query, topics)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done!")
	os.Exit(0)
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
