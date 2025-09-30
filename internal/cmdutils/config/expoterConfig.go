package config

type ExpoterConfig struct {
	OtelExpoterConfig
}

type OtelExpoterConfig struct {
}

func GetExpoterConfig() *ExpoterConfig {
	return &ExpoterConfig{}
}
