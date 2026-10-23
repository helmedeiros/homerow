package memstore_test

import (
	"testing"

	"github.com/helmedeiros/homerow/internal/adapters/memstore"
	"github.com/helmedeiros/homerow/internal/app"
	"github.com/helmedeiros/homerow/internal/app/storetest"
)

func TestMemoryStoreHonoursTheContract(t *testing.T) {
	storetest.Contract(t, func(*testing.T) app.ProgressStore { return memstore.New() })
}
