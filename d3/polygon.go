package d3

type Polygon struct {
	Vertices []Vertex
}

// Add a vertex to the polygon
func (p *Polygon) AddVertex(x float32, y float32, z float32) {
	v := Vertex{X: x, Y: y, Z: z}
	p.Vertices = append(p.Vertices, v)
}

// Render the polygon as a float32 slice
func (p Polygon) Render() (r []float32) {
	for _, vertex := range p.Vertices {
		r = append(r, vertex.X)
		r = append(r, vertex.Y)
		r = append(r, vertex.Z)
	}
	return
}
