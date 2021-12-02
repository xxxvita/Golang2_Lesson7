package structurecopy_test

// go test -v
import (
	"gb/Golang2_Lesson7/task1/structurecopy"
	"reflect"
	"testing"
)

// Тест на входящую структуры равную nil
func TestSimpleStruct(t *testing.T) {
	//config := Config{}
	mapConfig := map[string]interface{}{}

	err := structurecopy.StructureSet(nil, mapConfig)
	if err == nil {
		t.Error("Должна быть ошибка о nil-структуре")
	}
}

// Тест на входящий параметр(который должен быть структурой) типа не структуры
func TestWrongTypeInStruct(t *testing.T) {
	type ConfigMap map[string]interface{}

	var a int = 5
	config := &a

	mapConfig := ConfigMap{}

	err := structurecopy.StructureSet(config, mapConfig)
	if err == nil {
		t.Error("Должна быть ошибка о наличие указателя на struct")
	}
}

// Финальные структуры не равны
func TestABStruct(t *testing.T) {
	type Config struct {
		A int64
		B string
	}

	type ConfigMap map[string]interface{}

	config := Config{}

	configTest := Config{
		A: 5,
		B: "Hello",
	}

	mapConfig := ConfigMap{
		"A": interface{}(int64(5)),
		"B": "Hello",
	}

	err := structurecopy.StructureSet(&config, mapConfig)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(config, configTest) {
		t.Errorf("Финальные структуры не равны (%+v)!=(%+v)", config, configTest)
	}
}

// Попытка использовать в map-config неверный тип поля
func TestWrongTypeMapStruct(t *testing.T) {
	type Config struct {
		A int64
		B string
		C *Config
	}

	type ConfigMap map[string]interface{}

	config := Config{}

	mapConfig := ConfigMap{
		"A": interface{}(int64(5)),
		"B": "Hello",
		"C": 5, // должен быть тип *Config
	}

	err := structurecopy.StructureSet(&config, mapConfig)
	if err == nil {
		t.Error("Не произошла ошибка при присвоении полю C значения с ошибочным типом")
	}
}

// Тестируются разные типы полей
func TestMultyTypeMapStruct(t *testing.T) {
	type Config struct {
		A int64
		B string
		C *Config
		D struct{ D1 int }
		E struct {
			E1 int
			E2 int
		}
	}

	type ConfigMap map[string]interface{}

	config := Config{}
	simpleConfig := &Config{}

	mapConfig := ConfigMap{
		"A": interface{}(int64(5)),
		"B": "Hello",
		"C": simpleConfig,
		"D": struct{ D1 int }{2},
		"E": map[string]interface{}{"E1": 888, "E2": 999},
	}

	configTest := Config{
		A: 5,
		B: "Hello",
		C: simpleConfig,
		D: struct{ D1 int }{2},
		E: struct {
			E1 int
			E2 int
		}{E1: 888, E2: 999},
	}

	err := structurecopy.StructureSet(&config, mapConfig)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(config, configTest) {
		t.Errorf("Финальные структуры не равны (%+v)!=(%+v)", config, configTest)
	}
}
