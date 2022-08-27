package config

// Good fo testing proposes
type configMap struct {
	confMap map[string]interface{}
}

func (v configMap) GetInt(key string) int {
	return v.confMap[key].(int)
}

func (v configMap) Get(key string) interface{} {
	return v.confMap[key]
}

func (v configMap) GetString(key string) string {
	return v.confMap[key].(string)
}

func CreateConfigFromMap(confMap map[string]interface{}) Config {
	return configMap{confMap: confMap}
}
