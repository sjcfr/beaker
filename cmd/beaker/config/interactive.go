package config

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/allenai/beaker/config"
)

func newInteractiveCmd(
	parent *kingpin.CmdClause,
	parentOpts *configOptions,
	cfg *config.Config,
) {
	cmd := parent.Command("interactive", "Interactive configuration")
	cmd.Action(func(c *kingpin.ParseContext) error {
		return config.InteractiveConfiguration()
	})
}
