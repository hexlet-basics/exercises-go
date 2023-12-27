Реализуйте функцию `CopyParent(p *Parent) Parent`, которая создает копию структуры Parent и возвращает ее:

```go
type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

cp := CopyParent(nil) // Parent{}

p := &Parent{
   Name: "Harry",
   Children: []Child{
       {
           Name: "Andy",
           Age:  18,
       },
   },
}
cp := CopyParent(p)

// при мутациях в копии "cp"
// изначальная структура "p" не изменяется
cp.Children[0] = Child{}
fmt.Println(p.Children) // [{Andy 18}]
```
