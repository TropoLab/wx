package d3

import "log"

type Mesh struct {
	polygons []Vertex
}

func (m Mesh) ToObj() error {
	log.Printf("polygons: %v\n", m.polygons)
	return nil
}
