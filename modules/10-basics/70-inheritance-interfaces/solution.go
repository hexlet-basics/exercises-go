package solution

type IVoiceable interface {
    Voice() string
}

type Cat struct {
    // … 
}

type Cow struct {
    // … 
}

type Dog struct {
	// … 
}

// BEGIN

func (c Cat) Voice() string {
	return "Мяу"
}

func (cw Cow) Voice() string {
	return "Мууу"
}

func (d Dog) Voice() string {
    return "Гав"
}

// END
