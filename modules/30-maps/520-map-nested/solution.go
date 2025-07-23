package solution

// BEGIN

func SetUserSetting(settings map[string]map[string]string, user, key, value string) {
	if settings[user] == nil {
		settings[user] = make(map[string]string)
	}
	settings[user][key] = value
}

// END
