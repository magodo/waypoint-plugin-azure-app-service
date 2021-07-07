package plugin

import "github.com/hashicorp/waypoint-plugin-sdk/component"

func (r *ReleaseOutput) URL() string { return r.Url }

var _ component.Release = (*ReleaseOutput)(nil)
