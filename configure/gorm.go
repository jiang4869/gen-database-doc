package configure

import "github.com/spf13/viper"

type GormProperties struct {
	SingularTable bool   `mapstructure:"singular-table" json:"singularTable" yaml:"singular-table"`
	TablePrefix   string `mapstructure:"table-prefix" json:"tablePrefix" yaml:"table-prefix"`
	PrepareStmt   bool   `mapstructure:"prepare-stmt" json:"prepareStmt" yaml:"prepare-stmt"`
}

func GormPropertiesDefault() {
	viper.SetDefault("gorm.singular-table", true)
	viper.SetDefault("gorm.table-prefix", "")
	viper.SetDefault("gorm.prepare-stmt", false)
}
