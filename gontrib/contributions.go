package gontrib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jubalh/gontributions/util"
	"github.com/jubalh/gontributions/vcs/git"
	"github.com/jubalh/gontributions/vcs/mediawiki"
)

// MediaWiki holds the base URL of the wiki page to which later the
// API call will get appended and the username to the wiki.
type MediaWiki struct {
	BaseUrl string
	User    string
}

// Project hold all important information
// about a project.
type Project struct {
	Name        string
	Description string
	URL         string
	Gitrepos    []string
	MediaWikis  []MediaWiki
}

// Configuration holds the users E-Mail adresses
// and Projects he contributed to.
type Configuration struct {
	Emails   []string
	Projects []Project
}

// Contribution hols the Projects the user
// contributed to and a the ammounts of contributions
type Contribution struct {
	Project Project
	Count   int
}

// ScanContributions takes a Configuration containing a list of emails
// and a list of projects and returns a list of Contributions
// which contain of a project, how often a user contributed to it and
// description of the project.
func ScanContributions(configuration Configuration) []Contribution {
	contributions := []Contribution{}

	os.Mkdir("repos", 0755)

	for _, project := range configuration.Projects {
		var sumCount int
		for _, repo := range project.Gitrepos {
			util.PrintInfo("Working on "+repo, util.PI_TASK)
			git.GetLatestGitRepo(repo)
			for _, email := range configuration.Emails {
				path := filepath.Join("repos", util.LocalRepoName(repo))
				count, err := git.CountCommits(path, email)
				if err != nil {
					fmt.Println(err)
				}

				util.PrintInfoF("%s: %d commits", util.PI_RESULT, email, count)

				sumCount += count
			}
		}

		for _, wiki := range project.MediaWikis {
			util.PrintInfoF("Working on MediaWiki %s as %s", util.PI_TASK, wiki.BaseUrl, wiki.User)

			wikiCount, err := mediawiki.GetUserEdits(wiki.BaseUrl, wiki.User)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			util.PrintInfoF("%d edits", util.PI_RESULT, wikiCount)

			sumCount += wikiCount
		}

		if sumCount > 0 {
			c := Contribution{project, sumCount}
			contributions = append(contributions, c)
		}
	}
	return contributions
}
