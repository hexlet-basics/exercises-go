package solution

import "errors"

// MergeDictsJob is a job to merge dictionaries into a single dictionary.
type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

// BEGIN

func (j *MergeDictsJob) setFinished() {
	j.IsFinished = true
}

// ExecuteMergeDictsJob executes the job: merges all dictionaries into a single one.
// After the job is finished the Finished flag is set to true.
func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer job.setFinished()

	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}

	job.Merged = make(map[string]string)
	for _, dict := range job.Dicts {
		if dict == nil {
			return job, errNilDict
		}

		for k, v := range dict {
			job.Merged[k] = v
		}
	}

	return job, nil
}

// END
