package structmerge

import (
	"fmt"
	"reflect"
)

type reflectValue map[string]reflect.Value

func (r reflectValue) DeepFields(iface any) error {
	rv := reflect.ValueOf(iface)
	if rv.Kind() != reflect.Struct || rv.IsNil() {
		return fmt.Errorf("elem parameter should be struct")
	}

	for i := 0; i < rv.NumField(); i++ {
		v := rv.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			r.DeepFields(v.Interface())
		default:
			filed := rv.Type().Field(i)
			fieldName := filed.Name
			tag := filed.Tag

			fieldValue := rv.Field(i)
			if tag.Get("merge") != "" {
				fieldName = tag.Get("merge")
			}
			r[fieldName] = fieldValue
		}
	}
	return nil
}

func (r reflectValue) Parse(v any) error {
	reflectEle := reflect.ValueOf(v)
	if reflectEle.Kind() != reflect.Struct {
		return fmt.Errorf("elem parameter should be struct")
	}

	for i := 0; i < reflectEle.NumField(); i++ {
		filed := reflectEle.Type().Field(i)

		fieldName := filed.Name
		tag := filed.Tag
		fieldValue := reflectEle.Field(i)

		if tag.Get("merge") != "" {
			fieldName = tag.Get("merge")
		}
		r[fieldName] = fieldValue
	}
	return nil
}

func (r reflectValue) GetValue(elem []any) error {
	for _, e := range elem {
		r.DeepFields(e)
	}
	return nil
}

func Merge(main any, elem ...any) error {
	r := make(reflectValue)
	if err := r.GetValue(elem); err != nil {
		return err
	}

	reflectMain := reflect.ValueOf(main)
	if reflectMain.Kind() != reflect.Ptr || reflectMain.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("main parameter should be struct or ptr")
	}
	reflectMain = reflectMain.Elem()

	for i := 0; i < reflectMain.NumField(); i++ {

		field := reflectMain.Type().Field(i)
		fieldName := field.Name
		tag := field.Tag
		filedValue := reflectMain.FieldByName(fieldName)

		if tag.Get("merge") != "" {
			fieldName = tag.Get("merge")
		}

		if v, ok := r[fieldName]; ok {
			filedValue.Set(v)
		}
	}
	return nil
}
