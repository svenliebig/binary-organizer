package cmd

import "github.com/spf13/cobra"

var annotationSourceTrue = map[string]string{
	"source": "true",
}

var annotationSourceFalse = map[string]string{
	"source": "false",
}

func shouldSource(cmd *cobra.Command) bool {
	return cmd.Annotations["source"] == "true"
}
