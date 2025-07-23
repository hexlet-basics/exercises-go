package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUserSetting(t *testing.T) {
	settings := map[string]map[string]string{
		"alice": {
			"theme": "dark",
			"lang":  "en",
		},
	}

	SetUserSetting(settings, "alice", "lang", "ru")
	assert.Equal(t, "ru", settings["alice"]["lang"])

	SetUserSetting(settings, "bob", "theme", "light")
	assert.Equal(t, "light", settings["bob"]["theme"])
	assert.Len(t, settings, 2)
}
