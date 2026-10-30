package content

import "github.com/helmedeiros/homerow/internal/domain/curriculum"

type Catalog struct{}

func (Catalog) Module() curriculum.Module {
	return HomeRow()
}
