package generate_token

import (
	"fmt"
	"github-app/config"

	"github.com/jferrl/go-githubauth"
	"github.com/spf13/cobra"
)

var (
	GenerateTokenCmd = &cobra.Command{
		Use:          "generate-token",
		Short:        "Generate an access token for the app",
		RunE:         runGenerateToken,
		SilenceUsage: true,
	}
)

func runGenerateToken(cmd *cobra.Command, args []string) error {
	config := config.GetConfig()

	appTokenSource, err := githubauth.NewApplicationTokenSource(config.ApplicationId, config.PrivateKey)
	if err != nil {
		return err
	}

	token, err := appTokenSource.Token()
	if err != nil {
		return err
	}
	fmt.Println(token.AccessToken)
	return nil
}

func init() {
	GenerateTokenCmd.AddCommand(installationCmd)
}
