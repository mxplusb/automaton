package initialiser

import (
	"github.com/sirupsen/logrus"
	"github.com/libgit2/git2go"
	"github.com/spf13/viper"
	"os"
	"path"
)

type WorkSpace struct {
	config        Config
	logger        *logrus.Logger
	viper         *viper.Viper
	key           string
	internalsPath string
	ManifestURL   string

	// Parent is the top level node in the workspace that determines
	// the layout and manages the children in said layout.
	Parent *Parent
}

func (ws *WorkSpace) loadViper() {
	ws.viper = viper.GetViper()
}

func (ws *WorkSpace) Clone() error {
	if err := os.Mkdir(".automaton", 0666); err != nil {
		return err
	}

	ws.internalsPath = path.Join(ws.Parent.Location, ".automaton")

	repo, err := git.Clone(ws.ManifestURL, ws.internalsPath + ws.key, &git.CloneOptions{
		Bare: false,
		CheckoutBranch: "master",
	})

	if err != nil {
		return nil
	}

	ws.Parent.Repository = repo

	return nil
}

type Parent struct {
	ws       WorkSpace
	Repository *git.Repository
	Location string
	Update   bool
	Children []*Child
}

type Child struct {
	parent Parent

	// Relative to the top level of the workspace.
	// Keeping it relative makes it relocatable.
	RelativePath string

	// native git ftw.
	Repository *git.Repository
}

type Config struct {
	User     string
	Email    string
	UserHome string
}
