package config

type ConfigFile struct {
	Sensors struct {
		K10temp struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"k10temp"`
	} `yaml:"sensors"`
	Thresholds struct {
		Low  int `yaml:"low"`
		High int `yaml:"high"`
	} `yaml:"thresholds"`
	Logging struct {
		Level    string `yaml:"level"`
		Output   string `yaml:"output"`
		FilePath string `yaml:"file_path"`
	} `yaml:"logging"`
}
