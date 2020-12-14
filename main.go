package main

import (
	"fmt"
	"os"

	"packer-plugin-example/builder/null"
	"packer-plugin-example/post-processor/manifest"
	"packer-plugin-example/provisioner/comment"

	"github.com/hashicorp/packer/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("null", new(null.Builder))
	pps.RegisterProvisioner("comment", new(comment.Provisioner))
	pps.RegisterPostProcessor("manifest", new(manifest.PostProcessor))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
