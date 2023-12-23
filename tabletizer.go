package tabletizer

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func DoMagic(a any, tableName string) string {
	structType := reflect.TypeOf(a)

	if structType.Kind() == reflect.Pointer {
		structType = structType.Elem()
	}

	fields := reflect.VisibleFields(structType)
	items := []string{}

	for _, v := range fields {
		fname := convertToSnakeCase(v.Name)
		stype := typeToSqlType(v)
		if stype == "" {
			continue
		}
		item := fmt.Sprintf("%s %s", fname, stype)
		if tag := v.Tag.Get("sqlarg"); tag != "" {
			item += " " + tag
		}
		items = append(items, item)
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, strings.Join(items, ", "))
}

func typeToSqlType(t reflect.StructField) string {
	if sqltype := t.Tag.Get("sqltype"); sqltype != "" {
		return sqltype
	}

	switch t.Type.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Bool:
		return "INTEGER"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	case reflect.String:
		return "TEXT"
	default:
		return ""
	}
}

func convertToSnakeCase(s string) string { // [stroke]thank you, ChatGPT[/stroke]
	var result strings.Builder
	for i, char := range s {
		if i > 0 && unicode.IsUpper(char) && !unicode.IsUpper(rune(s[i-1])) {
			result.WriteRune('_')
		}
		result.WriteRune(char)
	}
	return strings.ToLower(result.String())
}
