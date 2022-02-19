package main

import (
	"fmt"
	"gen-database-doc/configure"
	"github.com/carmel/gooxml/color"
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/measurement"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/carmel/gooxml/schema/soo/wml"
)

var db *gorm.DB
var conf configure.Configure

func init() {
	InitConfig()

	InitDb()
}

func InitConfig() {

	configName := viper.GetString("config-name")

	fmt.Println("使用的配置文件为： ", configName)
	viper.SetConfigName("application.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&conf); err != nil {
			panic(err)
		}
	})
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
}

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", conf.DataSource.Username, conf.DataSource.Password, conf.DataSource.Address, conf.DataSource.Port, conf.DataSource.Dbname,conf.DataSource.Config)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var err error
	db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   conf.Gorm.TablePrefix,
				SingularTable: true,
			},
			PrepareStmt: true,
			Logger:      logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic("failed to connect database")
	}

}

type Result struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	ColumnType    string `gorm:"column:COLUMN_TYPE"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
	ColumnKey     string `gorm:"column:COLUMN_KEY"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}

func GetTables() []string {
	var tables []string
	db.Raw("show tables").Scan(&tables)

	return tables
}

func GetTableInfo(tableName string) []Result {
	results := make([]Result, 0)
	db.Raw("select column_name,column_type,is_nullable,column_key, column_comment from information_schema.columns where table_schema =?  and table_name = ?",conf.DataSource.Dbname, tableName).Scan(&results)
	return results
}

type Writer struct {
	doc *document.Document
}

func NewWriter() *Writer {
	doc := document.New()
	return &Writer{doc: doc}
}

func (w *Writer) WriterTable(tableName string, tableInfo []Result) {
	// 写入表名
	w.doc.AddParagraph().AddRun().AddText(tableName)

	// 添加一个表格
	table := w.doc.AddTable()
	// width of the page
	table.Properties().SetWidthPercent(100)
	// with thick borers
	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)
	row := table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("字段编号")
	row.AddCell().AddParagraph().AddRun().AddText("字段名")
	row.AddCell().AddParagraph().AddRun().AddText("字段类型")
	row.AddCell().AddParagraph().AddRun().AddText("是否为空")
	row.AddCell().AddParagraph().AddRun().AddText("字段键")
	row.AddCell().AddParagraph().AddRun().AddText("备注")

	for idx, val := range tableInfo {
		row = table.AddRow()
		row.AddCell().AddParagraph().AddRun().AddText(fmt.Sprintf("%d", idx+1))
		row.AddCell().AddParagraph().AddRun().AddText(val.ColumnName)
		row.AddCell().AddParagraph().AddRun().AddText(val.ColumnType)
		row.AddCell().AddParagraph().AddRun().AddText(val.IsNullable)
		row.AddCell().AddParagraph().AddRun().AddText(val.ColumnKey)
		row.AddCell().AddParagraph().AddRun().AddText(val.ColumnComment)
	}
	w.doc.AddParagraph()
}

func (w *Writer) Save(fileName string) error {
	return w.doc.SaveToFile(fileName)
}

func main() {
	tables := GetTables()
	writer := NewWriter()
	for _, tableName := range tables {
		results := GetTableInfo(tableName)
		writer.WriterTable(tableName, results)
	}
	writer.Save(fmt.Sprintf("%s.docx",conf.DataSource.Dbname))
}
