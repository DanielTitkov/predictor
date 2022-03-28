package configs

type (
	AuthConfig struct {
		Secret string
		Exp    int
		Google Google
	}
	Google struct {
		Client   string
		Secret   string
		Callback string
	}
)
