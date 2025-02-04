package activities

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GitDefaultCommiterName  string `envconfig:"GIT_COMMIT_NAME" default:"capten-bot"`
	GitDefaultCommiterEmail string `envconfig:"GIT_COMMIT_EMAIL" default:"capten-bot@intelops.dev"`
	VaultEntityName         string `envconfig:"VAULT_ENTITY_NAME" default:"onboarding"`
	GitCLoneDir             string `envconfig:"GIT_CLONE_DIR" default:"/gitCloneDir"`
}

func GetConfig() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
