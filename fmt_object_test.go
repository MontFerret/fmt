package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestObjectFormatter(t *testing.T) {
	Convey("Object", t, func() {
		Convey("Empty", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET    a =     {      }

RETURN     a

`)

			So(out, ShouldEqual, `LET a = {}

RETURN a`)
		})

		Convey("With primitives", func() {
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

		Convey("Not empty with computed", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET foo = "bar"
LET    a =     { a: 1, b:   2.5,   [foo]:     "bar"}

RETURN     a

`)

			So(out, ShouldEqual, `LET foo = "bar"
LET a = { a: 1, b: 2.5, [foo]: "bar" }

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

		Convey("With nested array", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET foo = "bar"
LET    a =     { a: 1, b:  { c: "d"  }, d:    [    1,2,   3,4,   5]  }

RETURN     a

`)

			So(out, ShouldEqual, `LET foo = "bar"
LET a = { a: 1, b: { c: "d" }, d: [1, 2, 3, 4, 5] }

RETURN a`)
		})

		Convey("Too long", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET foo = "bar"
LET    a =     { a: "Lorem ipsum dolor sit amet", b:  "Lorem ipsum dolor sit amet", c:    "Lorem ipsum dolor sit amet"  }

RETURN     a

`)

			So(out, ShouldEqual, `LET foo = "bar"
LET a = {
    a: "Lorem ipsum dolor sit amet",
    b: "Lorem ipsum dolor sit amet",
    c: "Lorem ipsum dolor sit amet"
}

RETURN a`)
		})

		Convey("Too long nested", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET foo = "bar"
LET    a =     { a: "foo", b:  { c: "Lorem ipsum dolor sit amet", d:    "Lorem ipsum dolor sit amet" }  }

RETURN     a

`)

			So(out, ShouldEqual, `LET foo = "bar"
LET a = {
    a: "foo",
    b: { c: "Lorem ipsum dolor sit amet", d: "Lorem ipsum dolor sit amet" }
}

RETURN a`)
		})
	})
}
