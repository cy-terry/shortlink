package config

type Config struct {
	Serve ServeConf `yaml:"serve"`
}

type ServeConf struct {
	Web         WebConf `yaml:"web"`
	Db          DBConf  `yaml:"database""`
	ShortPrefix string  `yaml:"short-prefix"`
}

type WebConf struct {
	Port     int    `yaml:"port""`
	RootPath string `yaml:"root-path""`
}

type DBConf struct {
	Type      string `yaml:"type""`
	DriveAddr string `yaml:"drive-addr""`
}

var Conf Config
