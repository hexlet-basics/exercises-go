
Реализуйте функцию `ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error)`, которая выполняет джобу *MergeDictsJob* и возвращает ее. Алгоритм обработки джобы следующий:
- перебрать по порядку все словари *job.Dicts* и записать каждое ключ-значение в результирующую мапу *job.Merged*
- если в структуре *job.Dicts* меньше 2-х словарей, возвращается ошибка `errNotEnoughDicts = errors.New("at least 2 dictionaries are required")`
- если в структуре *job.Dicts* встречается словарь в виде нулевого значения `nil`, то возвращается ошибка `errNilDict = errors.New("nil dictionary")`
- независимо от успешного выполнения или ошибки в возвращаемой структуре *MergeDictsJob* поле *IsFinished* должно быть заполнено как `true`

Пример работы:

```go
ExecuteMergeDictsJob(&MergeDictsJob{}) // &MergeDictsJob{IsFinished: true}, "at least 2 dictionaries are required"
ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"},nil}}) // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b"},nil}}, "nil dictionary"
ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"},{"b": "c"}}}) // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b", "b": "c"}}}, nil
```
