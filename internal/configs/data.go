package configs

type (
	DataConfig struct {
		Presets PresetsConfig
	}
	PresetsConfig struct {
		ChallengePresetsPaths []string `yaml:"challengePresetsPaths"`
	}
)
