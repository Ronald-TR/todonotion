package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display version",
		Long:  "Display todonotion's version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("v0.2.0\n")
		},
	}
)

func init() {

}
