package git

import (
	"os/exec"
)

// Git helper functions

// GetAuthor gets the author of the latest commit
func GetAuthor() string {
	e := exec.Command("git", "log", "-1", "--pretty=format:%an")
	s, err := e.Output()
	if err != nil {
		return ""
	}
	return string(s)
}

// GetCommitMessage gets the latest commit message
func GetCommitMessage() string {
	e := exec.Command("git", "log", "-1", "--pretty=format:%B")
	s, err := e.Output()
	if err != nil {
		return ""
	}
	return string(s)
}

// GetSha1 gets the sha1 of the lastest commit
func GetSha1() string {
	e := exec.Command("git", "rev-parse", "HEAD")
	s, err := e.Output()
	if err != nil {
		return ""
	}
	return string(s)
}

// GetBranchName gets the branch name of the lastest commit
func GetBranchName() string {
	e := exec.Command("git", "branch")
	s, err := e.Output()
	if err != nil {
		return ""
	}
	return string(s)
}
