package ebsvolume

import (
	"github.com/hashicorp/packer/template/interpolate"
	awscommon "github.com/remijouannet/packer-osc-plugins/builder/osc/common"
)

type BlockDevice struct {
	awscommon.BlockDevice `mapstructure:"-,squash"`
	Tags                  map[string]string `mapstructure:"tags"`
}

func commonBlockDevices(mappings []BlockDevice, ctx *interpolate.Context) (awscommon.BlockDevices, error) {
	result := make([]awscommon.BlockDevice, len(mappings))

	for i, mapping := range mappings {
		interpolateBlockDev, err := interpolate.RenderInterface(&mapping.BlockDevice, ctx)
		if err != nil {
			return awscommon.BlockDevices{}, err
		}
		result[i] = *interpolateBlockDev.(*awscommon.BlockDevice)
	}

	return awscommon.BlockDevices{
		LaunchBlockDevices: awscommon.LaunchBlockDevices{
			LaunchMappings: result,
		},
	}, nil
}
