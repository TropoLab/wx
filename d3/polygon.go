package d3

type Polygon struct {
	A Vertex
	B Vertex
	C Vertex
}

func (p Polygon) Render() []float32 {
	return []float32{
		p.A.X,
		p.A.Y,
		p.A.Z,
		p.B.X,
		p.B.Y,
		p.B.Z,
		p.C.X,
		p.C.Y,
		p.C.Z,
	}
}
