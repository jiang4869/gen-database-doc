# gen-database-doc

# 使用说明
修改配置文件
```yaml
datasource:
  username: root # 用户名
  password: 123456 # 密码
  config: charset=utf8mb4&parseTime=True&loc=Local #连接参数
  address: 192.168.253.100 # 数据库的ip
  port: 3306 # 端口
  dbname: blog # 数据库名

gorm:
  table_prefix: blog_ # 设置表前缀
```
安装依赖
```shell
go mod tidy
```
运行
```shell
go run main.go
```
运行完后会默认生成以数据库命名的`.docx`文件