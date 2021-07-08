package plugin

import "github.com/hashicorp/waypoint-plugin-sdk/component"

func (r *DeploymentOutput) URL() string { return r.Url }

var _ component.DeploymentWithUrl = (*DeploymentOutput)(nil)
