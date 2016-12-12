package oci

func ParseENVStr(envvar string) (res map[string]string, err error) {
	var (
		key string
		val string
	)
	begin := 0
	for i := range envvar {
		if envvar[i] == "=" {
			key = envvar[begin:i]
			val = envar[i+1:]
		}
	}
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid env variable string '%s'", envvar)
	}
	res[key] = val
	return
}
