package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatterReturn(t *testing.T) {
	Convey("Return", t, func() {
		Convey("String", func() {
			Convey("Single quote", func() {
				Convey("Should convert", func() {
					f := fmt.New(fmt.WithSingleQuote(true))

					out := f.MustFormat("RETURN \"FOO\"")
					So(out, ShouldEqual, "RETURN 'FOO'")
				})

				Convey("Should convert escape", func() {
					f := fmt.New(fmt.WithSingleQuote(true))

					So(f.MustFormat("RETURN \"FOO \\\"BAR\\\"\""), ShouldEqual, "RETURN 'FOO \\'BAR\\''")
				})

				Convey("Should remove unnecessary escape", func() {
					f := fmt.New(fmt.WithSingleQuote(true))

					So(f.MustFormat("RETURN \"FOO \\\"BAR\\\" \\'BAZ\\'\""), ShouldEqual, "RETURN 'FOO \\'BAR\\' BAZ'")
				})
			})

			Convey("Double quote", func() {
				Convey("Should convert", func() {
					f := fmt.New(fmt.WithSingleQuote(false))

					So(f.MustFormat("RETURN 'FOO'"), ShouldEqual, "RETURN \"FOO\"")
				})

				Convey("Should convert escape", func() {
					f := fmt.New(fmt.WithSingleQuote(false))

					So(f.MustFormat("RETURN 'FOO \\'BAR\\''"), ShouldEqual, "RETURN \"FOO \\\"BAR\\\"\"")
				})

				Convey("Should remove unnecessary escape", func() {
					f := fmt.New(fmt.WithSingleQuote(false))

					So(f.MustFormat("RETURN 'FOO \\'BAR\\' \\\"BAZ\\\"'"), ShouldEqual, "RETURN \"FOO \\\"BAR\\\" BAZ\"")
				})
			})
		})

		Convey("Boolean", func() {
			Convey("Boolean True", func() {
				f := fmt.New()
				out := f.MustFormat("RETURN   TRUE")

				So(out, ShouldEqual, "RETURN TRUE")
			})

			Convey("Boolean False", func() {
				f := fmt.New()
				out := f.MustFormat("      RETURN        FALSE")

				So(out, ShouldEqual, "RETURN FALSE")
			})
		})

		Convey("Integer", func() {
			Convey("Positive", func() {
				f := fmt.New()
				out := f.MustFormat("RETURN   100")

				So(out, ShouldEqual, "RETURN 100")
			})

			Convey("Negative", func() {
				f := fmt.New()
				out := f.MustFormat("      RETURN        -9999")

				So(out, ShouldEqual, "RETURN -9999")
			})
		})

		Convey("Float", func() {
			Convey("Positive", func() {
				f := fmt.New()
				out := f.MustFormat("RETURN   1.1")

				So(out, ShouldEqual, "RETURN 1.1")
			})

			Convey("Negative", func() {
				f := fmt.New()
				out := f.MustFormat("      RETURN        -0.1")

				So(out, ShouldEqual, "RETURN -0.1")
			})
		})

		Convey("None", func() {
			f := fmt.New()
			out := f.MustFormat("RETURN   NONE")

			So(out, ShouldEqual, "RETURN NONE")
		})

		Convey("Namespace", func() {
			Convey("USE", func() {
				f := fmt.New()
				out := f.MustFormat(`
		USE T:: HTML::BAR::BAZ

 RETURN FALSE
`)

				So(out, ShouldEqual, `USE T::HTML::BAR::BAZ

RETURN FALSE`)
			})

			SkipConvey("RETURN", func() {
				f := fmt.New()
				out := f.MustFormat(`
RETURN    IO::NET::HTTP::  GET(@lab.cdn.content)

`)

				So(out, ShouldEqual, `RETURN IO::NET::HTTP::GET(@lab.cdn.content)`)
			})
		})
	})
}
