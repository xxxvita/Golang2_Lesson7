package structurecopy

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	EInStructIsNIL = iota
)

func StructureSet(in interface{}, config map[string]interface{}) (err error) {
	if in == nil {
		return errors.New("входящая структура не может иметь значение nil")
	}

	val := reflect.ValueOf(in)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("входящая структура должна быть передана по ссылке (%+v)", val.Kind())
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("входящий интерфейс должен быть ссылкой на структуру")
	}

	//typeOfSpec := val.Type()
	//fmt.Printf("Число полей струткуры: %d\n", typeOfSpec.NumField())

	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Type().Field(i)

		if !val.Field(i).CanSet() {
			continue
		}

		// Получаю в мапе одноименный ключ. Если он есть, то копирую значения по ключу
		mapValue, ok := config[fieldType.Name]
		if !ok {
			continue
		}

		mapValueType := reflect.ValueOf(mapValue)

		// Если в значении карты тоже карта конфига map[string]interface{},
		// а тип поля по ключу - это структура, то рекурсивно копирую такие значения
		if fieldType.Type.Kind() == reflect.Struct && mapValueType.Kind() == reflect.Map {
			// Проверка типов ключа и значения
			mapRangeIter := mapValueType.MapRange()
			if mapRangeIter.Next() {

				//fmt.Printf("%+v", mapRangeIter.Key().Type().Kind())

				if mapRangeIter.Key().Type().Kind() == reflect.String &&
					mapRangeIter.Value().Type().Kind() == reflect.Interface &&
					!mapRangeIter.Value().IsNil() {

					if map_child, ok := mapValue.(map[string]interface{}); ok {
						return StructureSet(val.Field(i).Addr().Interface(), map_child)
					}
				}
			}
			continue
		}

		if fieldType.Type.Kind() != mapValueType.Kind() {
			return fmt.Errorf("поддерживаются поля с одинаковыми типами или оба структуры (%s : %s) != (%s : %s)",
				fieldType.Name,
				fieldType.Type.Kind(),
				fieldType.Name,
				mapValueType.Kind())
		}

		val.Field(i).Set(mapValueType)
	}

	return
}
