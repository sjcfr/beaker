package experiment

import (
	"context"
	"encoding/json"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/beaker/client/api"
	beaker "github.com/beaker/client/client"

	"github.com/allenai/beaker/config"
)

type inspectOptions struct {
	experiments []string
}

func newInspectCmd(
	parent *kingpin.CmdClause,
	parentOpts *experimentOptions,
	config *config.Config,
) {
	o := &inspectOptions{}
	cmd := parent.Command("inspect", "Display detailed information about one or more experiments")
	cmd.Action(func(c *kingpin.ParseContext) error {
		beaker, err := beaker.NewClient(parentOpts.addr, config.UserToken)
		if err != nil {
			return err
		}
		return o.run(beaker)
	})

	cmd.Arg("experiment", "Experiment name or ID").Required().StringsVar(&o.experiments)
}

func (o *inspectOptions) run(beaker *beaker.Client) error {
	ctx := context.TODO()

	var experiments []*api.Experiment
	for _, name := range o.experiments {
		experiment, err := beaker.Experiment(ctx, name)
		if err != nil {
			return err
		}

		exp, err := experiment.Get(ctx)
		if err != nil {
			return err
		}

		experiments = append(experiments, exp)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	return encoder.Encode(experiments)
}
