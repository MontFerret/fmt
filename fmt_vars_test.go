package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVariableFormatter(t *testing.T) {
	Convey("Variables", t, func() {
		Convey("Boolean", func() {
			f := fmt.New(fmt.WithSingleQuote(true))

			out := f.MustFormat(`
LET    a =     FALSE

RETURN     a

`)

			So(out, ShouldEqual, `LET a = FALSE

RETURN a`)
		})

		Convey("Integers", func() {
			f := fmt.New(fmt.WithSingleQuote(true))

			out := f.MustFormat(`
LET    a =     1

RETURN     a

`)

			So(out, ShouldEqual, `LET a = 1

RETURN a`)
		})

		Convey("Float", func() {
			f := fmt.New(fmt.WithSingleQuote(true))

			out := f.MustFormat(`
LET    a =     1.1

RETURN     a

`)

			So(out, ShouldEqual, `LET a = 1.1

RETURN a`)
		})

		Convey("None", func() {
			f := fmt.New(fmt.WithSingleQuote(true))

			out := f.MustFormat(`
LET    a =     NONE

RETURN     a

`)

			So(out, ShouldEqual, `LET a = NONE

RETURN a`)
		})
	})
}
