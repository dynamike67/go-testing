package command

import (
	"log"
  "fmt"

	"github.com/spf13/cobra"
  "github.com/dynamike67/go-testing/package/myshadowbroker"
)

var env *myshadowbroker.Environment

// RootCmd represents the base command when called without any subcommands
var RootCommand = &cobra.Command{
	Use:   "myshadowbroker",
	Short: "TESTING - Hayvan deployment pipeline",
}

// Execute app
func Execute() {
	if err := RootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}

func doit(a string) {
    fmt.Println("a")
}

func init() {
	RootCommand.PersistentFlags().String("project", "", "The name of the current project")
	RootCommand.PersistentFlags().String("branch", "", "The name of the current branch")
	//RootCmd.PersistentFlags().String("deployment-file", "deployment.yaml", "Deployment configuration file")
	//RootCmd.PersistentFlags().String("cloud-config", "", "Base64 encoded cloud config yaml file")
}

func environment() *myshadowbroker.Environment {
	if env == nil {
		project := RootCommand.Flag("project").Value.String()
		branch := RootCommand.Flag("branch").Value.String()
		//cloudConfig := RootCmd.Flag("cloud-config").Value.String()
		//deploymentConfigFile := RootCmd.Flag("deployment-file").Value.String()
    //env = myshadowbroker.NewEnvironment(project, branch, cloudConfig, deploymentConfigFile)
		env = myshadowbroker.NewEnvironment(project, branch)
		//errs := shadowbroker.ValidateEnvironment(env)

		//if len(errs) > 0 {
		//	log.Println("😐 configuration has errors")
		//	for _, err := range errs {
		//		log.Printf("%s", err)
		//	}
		//	log.Fatal("")
		//}
	}
	return env
}
