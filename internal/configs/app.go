package configs

type AppConfig struct {
	DefaultChallengePageLimit int    `yaml:"defaultChallengePageLimit"`
	HomeChallengePageLimit    int    `yaml:"homeChallengePageLimit"`
	SystemSummaryInterval     int    `yaml:"systemSummaryInterval"`
	SystemSummaryTimeout      int    `yaml:"systemSummaryTimeout"`
	DefaultTimeLayout         string `yaml:"defaultTimeLayout"`
}
