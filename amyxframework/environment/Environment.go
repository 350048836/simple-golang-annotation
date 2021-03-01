package environment

import "strconv"

var config = make(map[string]string)

func SetProperties(key string, value string) {
	config[key] = value
}

func GetProperty(key string) string {
	return config[key]
}

func GetPropertyAsBool(key string) bool {
	r, err := strconv.ParseBool(config[key])
	if err != nil {
		return false
	}
	return r
}

func GetPropertyAsInt(key string) int {
	r, err := strconv.Atoi(config[key])
	if err != nil {
		return 0
	}
	return r
}
