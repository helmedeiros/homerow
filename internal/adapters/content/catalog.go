package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

type Catalog struct{}

func (Catalog) Course() curriculum.Course {
	return curriculum.Course{Modules: []curriculum.Module{HomeRow()}}
}
