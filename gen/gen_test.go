package gen

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
	"testing"
)

var mysqlConfig = "root:@tcp(127.0.0.1:3306)/go?parseTime=true&loc=Local"

func TestGen(t *testing.T) {
	db, err := gorm.Open(mysql.Open(mysqlConfig))
	if err != nil {
		t.Error("连接mysql客户端失败")
		return
	}
	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		WithUnitTest:      true,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})

	g.UseDB(db)

	dataMap := map[string]func(fieldType string) string{
		"tinyint": func(fieldType string) string {
			return "int"
		},
		"smallint": func(fieldType string) string {
			return "int"
		},
		"mediumint": func(fieldType string) string {
			return "int"
		},
		"bigint": func(fieldType string) string {
			return "int64"
		},
	}

	g.WithDataTypeMap(dataMap)

	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	autoCreatedAtField := gen.FieldGORMTag("created_at", "column:createdAt;type:datetime;autoUpdateTime")
	autoUpdatedAtField := gen.FieldGORMTag("updated_at", "column:updatedAt;type:datetime;autoUpdateTime")
	softDeletedAtField := gen.FieldGORMTag("deleted_at", "soft_delete.DeletedAt")

	fieldOpts := []gen.ModelOpt{jsonField, autoCreatedAtField, autoUpdatedAtField, softDeletedAtField}

	allModel := g.GenerateAllTable(fieldOpts...)
	g.ApplyBasic(allModel...)
	g.ApplyInterface(func(method Method) {}, allModel...)
	g.Execute()
}

type Method interface{}
