package libs

//import "reflect"

//todo:这里通过反射来做
func Map2Struct(maps map[interface{}]interface{}, obj interface{}) (interface{}, error) {
	/*t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	mutable := reflect.ValueOf(&obj).Elem()

	for i := 0 ; i < t.NumField(); i++ {
		if value,ok := maps[t.Field(i)]; ok {
			type := t.Field(i).Type
			mutable.FieldByName(t.Field(i)).SetString(value)
		}
	}*/
}
