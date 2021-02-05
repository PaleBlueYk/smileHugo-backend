package config

type Config struct {
	Application string `yaml:"application"`
	Repository  string `yaml:"repository"`
	Port        string `yaml:"port"`
	RunMode     string `yaml:"runMode"`

	Logger   Logger   `yaml:"logger"`
	Database Database `yaml:"database"`
}

type Logger struct {
	Level         string `yaml:"level"`
	FilePath      string `yaml:"filePath"`
	ErrDetail     string `yaml:"errDetail"`
	ErrInResponse string `yaml:"errInResponse"`
}

type Database struct {
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	ShowSql  bool   `yaml:"showSql"`
}

func Init() error {

	return nil
}
