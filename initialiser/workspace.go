package initialiser

import (
	"github.com/sirupsen/logrus"
	"github.com/libgit2/git2go"
)

type WorkSpace struct {
	config      Config
	logger      *logrus.Logger
	ManifestURL string
	Parent      *Parent
}

type Parent struct {
	Location string
	Update bool
	Root string
}

type Child struct {
	RelativeLocation string
}

type Config struct {
	User string
	Email string
	UserHome string
}

type Project struct {
	Repository git.Repository
	// Path relative to the parent/workspace root.
	RelativePath string
}