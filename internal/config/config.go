package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type App struct {
	Name string `yaml:"name"`
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Params   string `yaml:"params"`
	MaxConn  int    `yaml:"max_conn"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxLife  int    `yaml:"max_life"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type JWT struct {
	Secret      string `yaml:"secret" mapstructure:"secret"`
	ExpireHours int    `yaml:"expire_hours" mapstructure:"expire_hours"`
}

type AppConfig struct {
	App      App      `yaml:"app"`
	Database Database `yaml:"database"`
	Logger   Logger   `yaml:"logger"`
	JWT      JWT      `yaml:"jwt"`
}

func (d *Database) GormDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.Params,
	)
}

func (d *Database) MigrateURL() string {
	return fmt.Sprintf(
		"%s://%s:%s@tcp(%s:%d)/%s?%s",
		d.Driver,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.Params,
	)
}

func LoadConfig(env string) (*AppConfig, error) {
	v := viper.New()

	v.SetConfigFile("configs/config.yaml")
	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("merge config: %w", err)
	}

	v.SetConfigFile(fmt.Sprintf("configs/config.%s.yaml", env))

	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("merge config: %w", err)
	}

	v.AutomaticEnv()

	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}
