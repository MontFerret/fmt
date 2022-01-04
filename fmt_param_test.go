package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParamFormatter(t *testing.T) {
	Convey("Param", t, func() {
		Convey("Single", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET a =      @param

RETURN a
`)
			So(out, ShouldEqual, `LET a = @param

RETURN a`)
		})

		Convey("Nested", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET a =      @one.    two

RETURN a
`)
			So(out, ShouldEqual, `LET a = @one.two

RETURN a`)
		})
	})
}
