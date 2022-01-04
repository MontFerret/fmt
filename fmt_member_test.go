package fmt_test

import (
	"github.com/MontFerret/fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMemberFormatter(t *testing.T) {
	Convey("Member", t, func() {
		Convey("With plain names", func() {
			Convey("Source is param", func() {
				f := fmt.New()

				out := f.MustFormat(`LET a = @param. foo.  bar
RETURN a`)

				So(out, ShouldEqual, `LET a = @param.foo.bar

RETURN a`)
			})

			Convey("Source is variable", func() {
				f := fmt.New()

				out := f.MustFormat(`LET a = var. foo.  bar
RETURN a`)

				So(out, ShouldEqual, `LET a = var.foo.bar

RETURN a`)
			})

			Convey("Source is a func call", func() {
				f := fmt.New()

				out := f.MustFormat(`LET a = FOO(). foo.  bar
RETURN a`)

				So(out, ShouldEqual, `LET a = FOO().foo.bar

RETURN a`)
			})

			Convey("Source is an obj literal", func() {
				f := fmt.New()

				out := f.MustFormat(`LET a = { foo:   1}. foo.  bar
RETURN a`)

				So(out, ShouldEqual, `LET a = { foo: 1 }.foo.bar

RETURN a`)
			})
		})

		Convey("With computed names", func() {
			Convey("Array", func() {
				Convey("Source", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = arr[ 0 ].  bar  .baz
RETURN a`)

					So(out, ShouldEqual, `LET a = arr[0].bar.baz

RETURN a`)
				})

				Convey("Segment", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = obj. arr[ 0 ].  bar
RETURN a`)

					So(out, ShouldEqual, `LET a = obj.arr[0].bar

RETURN a`)
				})
			})

			Convey("Object", func() {
				Convey("Source", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = obj [ "boo"].  bar
RETURN a`)

					So(out, ShouldEqual, `LET a = obj["boo"].bar

RETURN a`)
				})

				Convey("Segment", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = obj. obj2 [ "boo"].  bar
RETURN a`)

					So(out, ShouldEqual, `LET a = obj.obj2["boo"].bar

RETURN a`)
				})
			})

			Convey("Func", func() {
				f := fmt.New()

				out := f.MustFormat(`LET a = obj. obj2 [ FOO(1,23)].  bar
RETURN a`)

				So(out, ShouldEqual, `LET a = obj.obj2[FOO(1, 23)].bar

RETURN a`)
			})
		})

		Convey("Line max len", func() {
			f := fmt.New()

			out := f.MustFormat(`LET a = aoooooo.boooo.doooo.feeeeee.fffffff.ddddddd.eeeeeeee.ddddddddd
RETURN a`)

			So(out, ShouldEqual, `LET a = aoooooo.boooo.doooo.feeeeee.fffffff.ddddddd.eeeeeeee.ddddddddd

RETURN a`)
		})

		Convey("Optional chaining", func() {
			Convey("Plain", func() {
				Convey("Source", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = @param?. foo.  bar.foo
RETURN a`)

					So(out, ShouldEqual, `LET a = @param?.foo.bar.foo

RETURN a`)
				})

				Convey("Segment", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = @param. foo?.  bar?.foo
RETURN a`)

					So(out, ShouldEqual, `LET a = @param.foo?.bar?.foo

RETURN a`)
				})
			})

			Convey("Computed", func() {
				Convey("Source", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = arr?.[0]  .bar.foo
RETURN a`)

					So(out, ShouldEqual, `LET a = arr?.[0].bar.foo

RETURN a`)

					out2 := f.MustFormat(`LET a = FOO()?. [0]  .bar.foo
RETURN a`)

					So(out2, ShouldEqual, `LET a = FOO()?.[0].bar.foo

RETURN a`)
				})

				Convey("Segment", func() {
					f := fmt.New()

					out := f.MustFormat(`LET a = @param. arr?.[0]  .bar.foo
RETURN a`)

					So(out, ShouldEqual, `LET a = @param.arr?.[0].bar.foo

RETURN a`)
				})
			})
		})
	})
}
