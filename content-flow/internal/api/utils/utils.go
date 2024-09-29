package utils

import (
	"fmt"
	"reflect"
)

func UpdateStructField(obj any, fieldName string, newValue any) error {
	// 获取对象的反射值
	val := reflect.ValueOf(obj).Elem()

	// 获取字段
	field := val.FieldByName(fieldName)

	// 检查字段是否存在
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s in obj", fieldName)
	}

	// 检查字段是否可设置
	if !field.CanSet() {
		return fmt.Errorf("cannot set %s field value", fieldName)
	}

	// 确保新值的类型与字段类型相同
	newVal := reflect.ValueOf(newValue)
	if newVal.Type() != field.Type() {
		return fmt.Errorf("provided value type didn't match obj field type")
	}

	// 设置字段的新值
	field.Set(newVal)
	return nil
}
