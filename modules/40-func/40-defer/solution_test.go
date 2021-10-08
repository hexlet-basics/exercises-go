package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecuteMergeDictsJob(t *testing.T) {
	a := assert.New(t)

	job, err := ExecuteMergeDictsJob(&MergeDictsJob{})
	a.Equal(errNotEnoughDicts, err)
	a.Equal(true, job.IsFinished)

	job, err = ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
	}})
	a.Equal(errNotEnoughDicts, err)
	a.Equal(true, job.IsFinished)

	job, err = ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
		nil,
	}})
	a.Equal(errNilDict, err)
	a.Equal(true, job.IsFinished)

	job, err = ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b", "b": "c"},
		{"d": "e", "f": "g"},
		{"a": "z", "f": "g"},
	}})
	a.NoError(err)
	a.Equal(true, job.IsFinished)
	a.Equal(map[string]string{
		"b": "c",
		"d": "e",
		"a": "z",
		"f": "g",
	}, job.Merged)
}
