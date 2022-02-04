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

			Convey("With FOR body", func() {
				f := fmt.New()

				out := f.MustFormat(`
FOR link IN links
NAVIGATE(doc, link,   20000)
WAIT_ELEMENT(doc, '.c-entry-content', 15000)
LET texter      = ELEMENT(doc,    '.c-entry-content')
    RETURN texter.innerText
`)

				So(out, ShouldEqual, `FOR link IN links
    NAVIGATE(doc, link, 20000)
    WAIT_ELEMENT(doc, ".c-entry-content", 15000)
    LET texter = ELEMENT(doc, ".c-entry-content")

    RETURN texter.innerText`)
			})

			Convey("With LIMIT", func() {
				Convey("Primitive", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() LIMIT 1,    2      RETURN foo*2

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    LIMIT 1, 2
    RETURN foo * 2`)
				})

				Convey("Param", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() LIMIT @  skip,    @take      RETURN foo*2

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    LIMIT @skip, @take
    RETURN foo * 2`)
				})

				Convey("Func", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() LIMIT SKIP( 11,  [  1]),    @take      RETURN foo*2

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    LIMIT SKIP(11, [1]), @take
    RETURN foo * 2`)
				})
			})

			Convey("With SORT", func() {
				Convey("Singular", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() SORT    foo.bar     RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    SORT foo.bar
    RETURN foo.id`)
				})

				Convey("Singular with direction", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() SORT    foo.bar    ASC    RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    SORT foo.bar ASC
    RETURN foo.id`)
				})

				Convey("Plural", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() SORT    foo.bar,     foo.baz    RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    SORT foo.bar, foo.baz
    RETURN foo.id`)
				})

				Convey("Plural with singular direction (left)", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() SORT    foo.bar     ASC,     foo.baz    RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    SORT foo.bar ASC, foo.baz
    RETURN foo.id`)
				})

				Convey("Plural with singular direction (right)", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() SORT    foo.bar     ,     foo.baz  ASC  RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    SORT foo.bar, foo.baz ASC
    RETURN foo.id`)
				})

				Convey("Plural with plural direction", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() SORT    foo.bar     ASC,     foo.baz    DESC    RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    SORT foo.bar ASC, foo.baz DESC
    RETURN foo.id`)
				})
			})

			Convey("With FILTER", func() {
				Convey("Singular", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() FILTER    foo.bar >   1    RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    FILTER foo.bar > 1
    RETURN foo.id`)
				})

				Convey("Multiple one-line", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() FILTER    foo.bar >   1    AND foo.baz <    2   RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    FILTER foo.bar > 1 AND foo.baz < 2
    RETURN foo.id`)
				})

				Convey("Multiple multi-line", func() {
					f := fmt.New()

					out := f.MustFormat(`
FOR   foo IN     GET_DATA() FILTER    foo.bar >   1    FILTER    foo.baz <    2   RETURN foo.id

`)

					So(out, ShouldEqual, `FOR foo IN GET_DATA()
    FILTER foo.bar > 1
    FILTER foo.baz < 2
    RETURN foo.id`)
				})
			})
		})
	})
}
