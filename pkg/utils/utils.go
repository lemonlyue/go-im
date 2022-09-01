package utils

func ParseToStr(data map[string]string) string {
	if data == nil {
		return ""
	}
	values := ""
	for key, val := range data {
		values += "&" + key + "=" + val
	}
	values = values[1:]
	values = "?" + values
	return values
}