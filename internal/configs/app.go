package configs

type AppConfig struct {
	DefaultTaskPageLimit  int `yaml:"defaultTaskPageLimit"`
	SystemSummaryInterval int `yaml:"systemSummaryInterval"`
	SystemSummaryTimeout  int `yaml:"systemSummaryTimeout"`
}
