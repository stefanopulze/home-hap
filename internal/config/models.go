package config

type HomeKitOpts struct {
	Pin     string   `yaml:"pin" env:"HOMEKIT_PIN"`
	Name    string   `yaml:"name" env:"HOMEKIT_NAME"`
	Storage string   `yaml:"storage" env:"HOMEKIT_STORAGE"`
	Ifaces  []string `yaml:"ifaces" env:"HOMEKIT_IFACES"`
}

type MetricsOpts struct {
	Enabled           bool `yaml:"enabled"`
	EnableOpenMetrics bool `yaml:"enableOpenMetrics"`
}

type LogOpts struct {
	Level string `yaml:"level"`
}

type HomeRelay struct {
	Server string `yaml:"server"`
}

type DaitemConfig struct {
	Email      string `yaml:"email" env:"DAITEM_EMAIL"`
	Password   string `yaml:"password" env:"DAITEM_PASSWORD"`
	MasterCode string `yaml:"master-code" env:"DAITEM_MASTER_CODE"`
}
