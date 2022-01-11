package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatterOptions(t *testing.T) {
	Convey("Formatter Options", t, func() {
		Convey("CaseMode", func() {
			Convey("When CaseModeUpper", func() {
				f := fmt.New(fmt.WithCaseMode(fmt.CaseModeUpper))

				out := f.MustFormat("return  2    in [ 1, 33]")
				So(out, ShouldEqual, "RETURN 2 IN [1, 33]")
			})

			Convey("When CaseModeLower", func() {
				f := fmt.New(fmt.WithCaseMode(fmt.CaseModeLower))

				out := f.MustFormat("RETurn  2    in [ 1, 33]")
				So(out, ShouldEqual, "return 2 in [1, 33]")
			})

			Convey("When CaseModeIgnore", func() {
				f := fmt.New(fmt.WithCaseMode(fmt.CaseModeIgnore))

				out := f.MustFormat("RETurn  2    In [ 1, 33]")
				So(out, ShouldEqual, "RETurn 2 In [1, 33]")
			})
		})
	})
}
