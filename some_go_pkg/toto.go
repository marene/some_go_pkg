package somegopkg

import (
	"fmt"

	"github.com/marene/some_go_pkg/internal/bar"
	"github.com/marene/some_go_pkg/internal/foo"
)

func MainExposedFunc() string {
	return fmt.Sprintf("%s - %d", foo.SomeFunc(), bar.SomeOtherFunc())
}
