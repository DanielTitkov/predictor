package configs

type AppConfig struct {
	DefaultChallengePageLimit int `yaml:"defaultChallengePageLimit"`
	SystemSummaryInterval     int `yaml:"systemSummaryInterval"`
	SystemSummaryTimeout      int `yaml:"systemSummaryTimeout"`
}
