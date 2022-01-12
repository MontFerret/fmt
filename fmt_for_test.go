package fmt_test

import (
	"github.com/MontFerret/fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestForFormatter(t *testing.T) {
	Convey("For", t, func() {
		Convey("For-in", func() {
			Convey("Simple FOR-IN-RETURN", func() {
				f := fmt.New()

				out := f.MustFormat(`
FOR   foo IN     1..10      RETURN foo*2

`)

				So(out, ShouldEqual, `FOR foo IN 1..10
    RETURN foo * 2`)
			})

			Convey("With LIMIT", func() {
				f := fmt.New()

				out := f.MustFormat(`
FOR   foo IN     GET_DATA() LIMIT 1,    2      RETURN foo*2

`)

				So(out, ShouldEqual, `FOR foo IN GET_DATA()
    LIMIT 1, 2
    RETURN foo * 2`)
			})
		})
	})
}
