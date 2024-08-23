package cmd

import "github.com/spf13/cobra"

// if this annotation is used the command will source the $PATH after the command is executed.
var annotationSourceTrue = map[string]string{
	"source": "true",
}

// if this annotation is used the command will not source the $PATH after the command is executed.
var annotationSourceFalse = map[string]string{
	"source": "false",
}

func shouldSource(cmd *cobra.Command) bool {
	return cmd.Annotations["source"] == "true"
}
