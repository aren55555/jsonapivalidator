package jsonapivalidator

func newStringSet(args ...string) (s map[string]interface{}) {
	s = make(map[string]interface{})
	for _, str := range args {
		s[str] = nil
	}
	return
}

func keys(s map[string]interface{}) (keys []string) {
	for k := range s {
		keys = append(keys, k)
	}
	return
}
