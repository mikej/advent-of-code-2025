package main

type Beam struct {
	x        int
	hasSplit bool
}

type Splitter struct {
	x, y int
}

type Manifold struct {
}

func NewManifold(input []string) *Manifold {
	return &Manifold{}
}

func (m *Manifold) Run() {

}

func (m *Manifold) SplitCount() int {
	return 0
}

func main() {
}
