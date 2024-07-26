package champetre

import "reflect"

// Transform the element to a first "leaf" map where the references to other models are transformed
// into {ref: model.id(), type: model.kind()}
// it recursively iterate on the object
// TODO : set a timeout in case of recursive circle
func Serialize(model Model) map[string]any {
	res := map[string]any{}
	for name, value := range model.Parameters() {
		vof := reflect.ValueOf(value)
		switch vof.Kind() {

		case reflect.Struct :
			res[name] = SerializeStruct(value)

		case reflect.Slice, reflect.Array:
			res[name] = SerializeList(value)

		case reflect.Map:
			res[name] = SerializeMap(value)
		// deal with default values such as string, int, float, bool, etc
		default:
			res[name] = value
		}
	}
	return res
}

func SerializeList(l any) []any {
	return []any{}
}

func SerializeMap(m any) map[string]any {
	return map[string]any{}
}

func SerializeStruct(value any) map[string]any {
	switch v := value.(type) {
	case Model:
		return map[string]any{
			"ref": v.Id(),
			"kind": v.Kind(),
		}
	default:
		// TODO default stuff
		return SerializeStruct(value)
	}
}