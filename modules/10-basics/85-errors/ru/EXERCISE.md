Реализуйте функцию `GetFileExtension(filename string) (string, error)`, которая возвращает расширение файла (текст после последней точки). Если в имени файла нет точки, функция должна вернуть ошибку.

**Примеры**

```go
ext, err := GetFileExtension("photo.jpeg")
if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Расширение файла:", ext)
}
// => Расширение файла: jpeg

ext, err = GetFileExtension("backup.zip")
if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Расширение файла:", ext)
}
// => Расширение файла: zip

ext, err = GetFileExtension("LICENSE")
if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Расширение файла:", ext)
}
// => Ошибка: файл "LICENSE" не имеет расширения

```
