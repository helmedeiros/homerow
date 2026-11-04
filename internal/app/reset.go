package app

import (
	"fmt"

	"github.com/helmedeiros/homerow/internal/domain/progress"
)

func (t *Trainer) Reset() error {
	if err := t.store.Save(progress.NewBoard(nil)); err != nil {
		return fmt.Errorf("resetting progress: %w", err)
	}
	return nil
}
