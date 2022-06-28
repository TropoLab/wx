package d3

import "log"

type Mesh struct {
	Polygons []Polygon
}

// convert the mesh to an object
func (m Mesh) ToObj() error {
	log.Printf("polygons: %v\n", m.Polygons)
	return nil
}

// render the mesh as a float32 slice
func (m Mesh) Render() (r []float32) {
	for _, polygon := range m.Polygons {
		r = append(r, polygon.Render()...)
	}
	return
}
