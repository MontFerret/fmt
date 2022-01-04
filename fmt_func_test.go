package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFunctionFormatter(t *testing.T) {
	Convey("Functions", t, func() {
		Convey("Without arguments", func() {
			Convey("Without namespace", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =     FOO(   )

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO()

RETURN a`)
			})

			Convey("With namespace", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =  FOO:: BAR::   BAZ(   )

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO::BAR::BAZ()

RETURN a`)
			})

			Convey("With error suppression", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =  FOO:: BAR::   BAZ(   ) ?

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO::BAR::BAZ()?

RETURN a`)
			})
		})

		Convey("With arguments", func() {
			Convey("Primitive types", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =     FOO( TRUE, "FOO",   1 )

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO(TRUE, "FOO", 1)

RETURN a`)
			})

			Convey("Complex types", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =     FOO( [ 1, 2,  3,  4], { foo:     "bar", }   )

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO([1, 2, 3, 4], { foo: "bar" })

RETURN a`)
			})

			Convey("Nested call", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =     FOO( BAZ( [ 1, 2,  3,  4], { foo:     "bar", } ), GAZ()  )

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO(BAZ([1, 2, 3, 4], { foo: "bar" }), GAZ())

RETURN a`)
			})

			Convey("Too long", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =     FOO( 'Lorem ipsum dolor sit amet',        'Lorem ipsum dolor sit amet',     'Lorem ipsum dolor sit amet', 'Lorem ipsum dolor sit amet')

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO(
    "Lorem ipsum dolor sit amet",
    "Lorem ipsum dolor sit amet",
    "Lorem ipsum dolor sit amet",
    "Lorem ipsum dolor sit amet"
)

RETURN a`)
			})

			Convey("Too long nested", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET    a =     FOO( BAR('Lorem ipsum dolor sit amet',        'Lorem ipsum dolor sit amet',     'Lorem ipsum dolor sit amet', 'Lorem ipsum dolor sit amet')?)

RETURN     a

`)

				So(out, ShouldEqual, `LET a = FOO(
    BAR(
        "Lorem ipsum dolor sit amet",
        "Lorem ipsum dolor sit amet",
        "Lorem ipsum dolor sit amet",
        "Lorem ipsum dolor sit amet"
    )?
)

RETURN a`)
			})
		})
	})
}
