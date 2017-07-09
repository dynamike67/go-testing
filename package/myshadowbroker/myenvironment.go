package myshadowbroker

// Environment contains all runtime configuration
type Environment struct {
	Project          string
	Branch           string
	CloudConfig      *CloudConfig
	DeploymentConfig *DeploymentConfig
}

// NewEnvironment constructs a new Environment struct based on the provided data
func NewEnvironment(project string, branch string, cloudConfigData string, deploymentConfigFile string) *Environment {
	env := Environment{
		Project:          project,
		Branch:           branch,
		DeploymentConfig: parseDeploymentConfig(deploymentConfigFile),
		CloudConfig:      decodeCloudConfig(cloudConfigData),
	}

	return &env
}
