package cmd

import (
	"log"

	"github.com/spf13/cobra"
  "github.com/dynamike67/go-testing/package/myshadowbroker"
)

var env *shadowbroker.Environment

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "shadowbroker",
	Short: "Hayvan deployment pipeline",
}

// Execute app
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.PersistentFlags().String("project", "", "The name of the current project")
	RootCmd.PersistentFlags().String("branch", "", "The name of the current branch")
	RootCmd.PersistentFlags().String("deployment-file", "deployment.yaml", "Deployment configuration file")
	RootCmd.PersistentFlags().String("cloud-config", "", "Base64 encoded cloud config yaml file")
}

func environment() *shadowbroker.Environment {
	if env == nil {
		project := RootCmd.Flag("project").Value.String()
		branch := RootCmd.Flag("branch").Value.String()
		cloudConfig := RootCmd.Flag("cloud-config").Value.String()
		deploymentConfigFile := RootCmd.Flag("deployment-file").Value.String()
		env = shadowbroker.NewEnvironment(project, branch, cloudConfig, deploymentConfigFile)
		errs := shadowbroker.ValidateEnvironment(env)

		if len(errs) > 0 {
			log.Println("😐 configuration has errors")
			for _, err := range errs {
				log.Printf("%s", err)
			}
			log.Fatal("")
		}
	}
	return env
}
