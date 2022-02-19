package configure

import "github.com/spf13/viper"

type DataSourceProperties struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
	Address      string `mapstructure:"address" json:"address" yaml:"address"`                    // 数据库地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // 数据库使用的端口
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                       // 数据库名
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

func DataSourcePropertiesDefault() {
	viper.SetDefault("datasource.username", "root")
	viper.SetDefault("datasource.password", "123456")
	viper.SetDefault("datasource.config", "charset=utf8&parseTime=True&loc=Local")
	viper.SetDefault("datasource.address", "127.0.0.1")
	viper.SetDefault("datasource.port", "3306")
	viper.SetDefault("datasource.dbname", "")
	viper.SetDefault("datasource.maxidleconns", "10")
	viper.SetDefault("datasource.maxopenconns", "100")
}

func (data *DataSourceProperties) Dsn() string {
	return data.Username + ":" + data.Password + "@tcp(" + data.Address + ")/" + data.Dbname + "?" + data.Config
}
