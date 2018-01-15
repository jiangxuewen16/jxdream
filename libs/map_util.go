package libs

import (
	"reflect"
	"fmt"
)

//import "reflect"

//todo:这里通过反射来做
func Map2Struct(maps map[interface{}]interface{}, obj interface{}) {
	t := reflect.TypeOf(obj)
	mutable := reflect.ValueOf(&obj).Elem()
	fmt.Println(mutable)
	mutable.FieldByName(t.Field(1).Name).CanSet()

	for i := 0 ; i < t.NumField(); i++ {
		if value,ok := maps[t.Field(i).Name]; ok {
			if mutable.FieldByName(t.Field(i).Name).CanSet() {
				switch t.Field(i).Type.String() {
				case "int":
					mutable.FieldByName(t.Field(i).Name).SetInt(value.(int64))
				case "string":
					mutable.FieldByName(t.Field(i).Name).SetString(value.(string))
				default:

				}
			}
		}
	}
}
