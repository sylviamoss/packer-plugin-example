package version

import (
	"github.com/hashicorp/packer/packer-plugin-sdk/version"
)

// The main version number that is being run at the moment.
const Version = "1.0.0"

// A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = "dev"

var ExamplePluginVersion *version.PluginVersion

func init() {
	ExamplePluginVersion = version.InitializePluginVersion(
		Version, VersionPrerelease)
}
