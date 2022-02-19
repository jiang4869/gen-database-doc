package configure

import (
	"github.com/spf13/viper"
)

type Configure struct {
	DataSource DataSourceProperties `mapstructure:"datasource" json:"datasource" yaml:"datasource"`
	Gorm       GormProperties       `mapstructure:"gorm" json:"gorm" yaml:"gorm"`
}

func init() {
	DataSourcePropertiesDefault()
	ConfigurePropertiesDefault()
	GormPropertiesDefault()
}

func ConfigurePropertiesDefault() {
	viper.SetDefault("config-name", "application.yaml")
	viper.AddConfigPath(".")
}
