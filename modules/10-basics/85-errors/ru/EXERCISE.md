Реализуйте функцию `GetFileExtension(filename string) (string, error)`,  
которая возвращает расширение файла (текст после последней точки).  
Если в имени файла нет точки, функция должна вернуть ошибку.

Также создайте функцию `PrintFileExtension(filename string)`,  
которая вызывает `GetFileExtension` и печатает расширение или сообщение об ошибке.

**Примеры**

```go
PrintFileExtension("document.txt")
// Расширение файла: txt

PrintFileExtension("archive.tar.gz")
// Расширение файла: gz

PrintFileExtension("README")
// Ошибка: файл "README" не имеет расширения
```
