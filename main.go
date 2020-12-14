package main

import (
	"fmt"
	"os"

	"packer-plugin-example/builder/null"
	"packer-plugin-example/post-processor/manifest"
	"packer-plugin-example/provisioner/comment"

	"github.com/hashicorp/packer/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer/packer-plugin-sdk/version"
)

var (
	// The main version number that is being run at the moment.
	Version = "1.0.0"

	// A pre-release marker for the version. If this is "" (empty string)
	// then it means that it is a final release. Otherwise, this is a pre-release
	// such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "dev"
)

var ExamplePluginVersion *version.PluginVersion

func init() {
	ExamplePluginVersion = version.InitializePluginVersion(
		Version, VersionPrerelease)
}

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("null", new(null.Builder))
	pps.RegisterProvisioner("comment", new(comment.Provisioner))
	pps.RegisterPostProcessor("manifest", new(manifest.PostProcessor))
	pps.SetVersion(ExamplePluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
