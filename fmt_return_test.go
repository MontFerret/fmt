package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatter(t *testing.T) {
	Convey("Formatter", t, func() {
		Convey("StringLiteral", func() {
			Convey("Single quote", func() {
				Convey("Should convert", func() {
					f := fmt.New(fmt.WithSingleQuote(true))

					So(f.MustFormat("RETURN \"FOO\""), ShouldEqual, "RETURN 'FOO'")
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

		Convey("RETURN", func() {
			Convey("Boolean", func() {
				f := fmt.New()
				out := f.MustFormat("RETURN   TRUE")

				So(out, ShouldEqual, "RETURN TRUE")
			})
		})

		Convey("Variables", func() {
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

			Convey("Array", func() {
				Convey("Empty", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET    a =     [      ]

RETURN     a

`)

					So(out, ShouldEqual, `LET a = []
RETURN a`)
				})

				Convey("Not empty", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET    a =     [ 1, 2.5,      "foo",    
NONE     ]

RETURN     a

`)

					So(out, ShouldEqual, `LET a = [1, 2.5, "foo", NONE]
RETURN a`)
				})
			})

			Convey("Object", func() {
				Convey("Empty", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET    a =     {      }

RETURN     a

`)

					So(out, ShouldEqual, `LET a = {}
RETURN a`)
				})

				Convey("Not empty", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET    a =     { a: 1, b:   2.5,   c:   "foo",    
d: NONE}

RETURN     a

`)

					So(out, ShouldEqual, `LET a = { a: 1, b: 2.5, c: "foo", d: NONE }
RETURN a`)
				})

				Convey("Not empty with shorthand", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET foo = "bar"
LET    a =     { a: 1, b:   2.5,   foo}

RETURN     a

`)

					So(out, ShouldEqual, `LET foo = "bar"
LET a = { a: 1, b: 2.5, foo }
RETURN a`)
				})
			})
		})
	})
}
