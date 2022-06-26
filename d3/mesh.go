package d3

import "log"

type Mesh struct {
	Polygons []Polygon
}

func (m Mesh) ToObj() error {
	log.Printf("polygons: %v\n", m.Polygons)
	return nil
}
