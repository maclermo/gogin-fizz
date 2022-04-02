package dog

type DogParams struct {
	Name string `path:"name"`
}

type Dog struct {
	name   string
	age    int
	weight float32
	colors  []DogColor
}

type DogColor struct {
	Name string
	Hex  string
}
