// ------------ format struct to url form -----------------------------------
func StructToUrlValues(st interface{}) url.Values {

	s := reflect.ValueOf(st)
	typeOf := s.Type()

	// get length of struct
	length := s.NumField()

	// create map of url.Values
	m := make(map[string][]string)

	// loop over struct
	for i := 0; i < length; i++ {

		var value string
		val := s.Field(i)

		// check struct value type and convert it to string
		switch s.Field(i).Kind() {
		case reflect.Int:
			value = strconv.FormatInt(val.Int(), 10)
		case reflect.Float64:
			value = strconv.FormatFloat(val.Float(), 'f', 2, 64)
		case reflect.Float32:
			value = strconv.FormatFloat(val.Float(), 'f', 2, 32)
		case reflect.Bool:
			value = strconv.FormatBool(val.Bool())
		case reflect.String:
			value = val.String()
		}

		// get struct Tag
		field := typeOf.Field(i).Tag.Get("urlVal")

		// update map
		if field == "" {
			m[typeOf.Field(i).Name] = append(m[field], value)
		} else {
			m[field] = append(m[typeOf.Field(i).Name], value)
		}
	}
	// return struct converted to map
	return m
}
