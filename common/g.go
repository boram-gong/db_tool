package common

type CfgDB struct {
	Host        string `yaml:"Host"`
	Port        int    `yaml:"Port"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Database    string `yaml:"Database"`
	MaxIdleConn int    `yaml:"MaxIdleConn"`
	MaxOpenConn int    `yaml:"MaxOpenConn"`
}

type Scanner interface {
	Next() bool
	MapScan(dest map[string]interface{}) error
	Close() error
}
