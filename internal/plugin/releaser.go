package plugin

import (
	"context"
	"fmt"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type ReleaseConfig struct {
	Active bool `hcl:"directory,optional"`
}

type Releaser struct {
	config ReleaseConfig
}

func (rm *Releaser) Config() (interface{}, error) {
	return &rm.config, nil
}

func (rm *Releaser) ConfigSet(config interface{}) error {
	_, ok := config.(*ReleaseConfig)
	if !ok {
		return fmt.Errorf("Expected *ReleaseConfig as parameter")
	}

	return nil
}

func (rm *Releaser) ReleaseFunc() interface{} {
	return rm.release
}

func (rm *Releaser) release(ctx context.Context, ui terminal.UI, deployment *DeploymentOutput) (*ReleaseOutput, error) {
	u := ui.Status()
	defer u.Close()
	ui.Output("Release application")

	u.Step(terminal.StatusOK, "Created release")

	return &ReleaseOutput{}, nil
}
