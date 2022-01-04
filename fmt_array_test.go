package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArrayFormatter(t *testing.T) {
	Convey("Array", t, func() {
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

		Convey("Of objects", func() {
			f := fmt.New(fmt.WithPrintWidth(40))

			out := f.MustFormat(`
LET    a =     [{ a: 1, b:  { c: "d"  }  }, { a: 1, b:  { c: "d"  }  },   { a: 1, b:  { c: "d"  }  },   { a: 1, b:  { c: "d"  }  }]

RETURN     a

`)
			So(out, ShouldEqual, `LET a = [
    { a: 1, b: { c: "d" } },
    { a: 1, b: { c: "d" } },
    { a: 1, b: { c: "d" } },
    { a: 1, b: { c: "d" } }
]

RETURN a`)
		})
	})
}
