package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatter(t *testing.T) {
	Convey("Formatter", t, func() {
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

				Convey("Nested", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET    a =     [ 1, 2.5,      "foo",    
NONE   , [ 1 , TRUE, FALSE,
"asda"
]  ]

RETURN     a

`)

					So(out, ShouldEqual, `LET a = [1, 2.5, "foo", NONE, [1, TRUE, FALSE, "asda"]]

RETURN a`)
				})

				Convey("Too long", func() {
					f := fmt.New(fmt.WithPrintWidth(40))

					out := f.MustFormat(`
LET    a =     ["Lorem ipsum dolor sit amet", "Lorem ipsum dolor sit amet", "sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"]

RETURN     a

`)

					So(out, ShouldEqual, `LET a = [
    "Lorem ipsum dolor sit amet",
    "Lorem ipsum dolor sit amet",
    "sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
]

RETURN a`)
				})

				Convey("Too long nested", func() {
					f := fmt.New(fmt.WithPrintWidth(40))

					out := f.MustFormat(`
LET    a =     ["Lorem ipsum dolor sit amet", "Lorem ipsum dolor sit amet", ["sed do eiusmod tempor incididunt ut labore et dolore magna aliqua", "quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat", ["foo","bar"]]]

RETURN     a

`)

					So(out, ShouldEqual, `LET a = [
    "Lorem ipsum dolor sit amet",
    "Lorem ipsum dolor sit amet",
    [
        "sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
        "quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat",
        ["foo", "bar"]
    ]
]

RETURN a`)
				})
			})

			SkipConvey("Object", func() {
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

				Convey("With nested object", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET foo = "bar"
LET    a =     { a: 1, b:  { c: "d"  }  }

RETURN     a

`)

					So(out, ShouldEqual, `LET foo = "bar"
LET a = { a: 1, b: { c: "d" } }
RETURN a`)
				})
			})

			SkipConvey("Functions", func() {
				Convey("Without arguments", func() {
					f := fmt.New(fmt.WithSingleQuote(true))

					out := f.MustFormat(`
LET    a =     FOO(   )

RETURN     a

`)

					So(out, ShouldEqual, `LET a = FOO()
RETURN a`)
				})
			})
		})
	})
}
