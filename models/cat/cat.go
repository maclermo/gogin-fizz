package cat

type CatParams struct {
	Name string `path:"name"`
}

type Cat struct {
	name   string
	age    int
	weight float32
	colors  []CatColor
}

type CatColor struct {
	Name string
	Hex  string
}
