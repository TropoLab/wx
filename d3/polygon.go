package d3

type Polygon struct {
	A Vertex
	B Vertex
	C Vertex
}

func (p Polygon) Render() [9]int32 {
	return [9]int32{
		p.A.x,
		p.A.y,
		p.A.z,
		p.B.x,
		p.B.y,
		p.B.z,
		p.C.x,
		p.C.y,
		p.C.z,
	}
}
