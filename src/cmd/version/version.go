package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appVersion               = "development"
	gitCommit                = "N/A"
	buildDate                = "N/A"
	shouldOutputVerbose bool = false

	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints version of program",
		Run:   runVersion,
	}
)

func runVersion(cmd *cobra.Command, args []string) {
	if shouldOutputVerbose {
		fmt.Printf("Version: %s Git commit: %s Build date=%s", appVersion, gitCommit, buildDate)
	} else {
		fmt.Printf("%s", appVersion)
	}
}

func init() {
	VersionCmd.Flags().BoolVarP(&shouldOutputVerbose, "verbose", "v", false, "verbose output")
}
