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
func (p Polygon) Render() (v []float32) {
	for _, vertex := range p.Vertices {
		v = append(v, vertex.X)
		v = append(v, vertex.Y)
		v = append(v, vertex.Z)
	}
	return
}
