package fmt_test

import (
	"github.com/MontFerret/fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOperatorsFormatter(t *testing.T) {
	Convey("Single Operators", t, func() {
		Convey("Range Operator", func() {
			Convey("Variable", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET     foo =     1..    2

RETURN     foo
`)

				So(out, ShouldEqual, `LET foo = 1..2

RETURN foo`)
			})

			Convey("Return", func() {
				f := fmt.New()

				out := f.MustFormat(`

RETURN     1..    2
`)

				So(out, ShouldEqual, `RETURN 1..2`)
			})
		})

		Convey("Unary Operator", func() {
			Convey("Variable", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET     foo = -   1

RETURN     foo
`)

				So(out, ShouldEqual, `LET foo = -1

RETURN foo`)
			})

			Convey("Return", func() {
				f := fmt.New()

				out := f.MustFormat(`

RETURN    !    TRUE
`)

				So(out, ShouldEqual, `RETURN !TRUE`)
			})
		})

		Convey("Logical AND Operator", func() {
			Convey("Variable", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET     foo = TRUE     &&      FALSE

RETURN     foo
`)

				So(out, ShouldEqual, `LET foo = TRUE && FALSE

RETURN foo`)
			})

			Convey("Return", func() {
				f := fmt.New()

				out := f.MustFormat(`

RETURN      TRUE     AND      FALSE
`)

				So(out, ShouldEqual, `RETURN TRUE AND FALSE`)
			})
		})

		Convey("Logical OR Operator", func() {
			Convey("Variable", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET     foo = TRUE     ||      FALSE

RETURN     foo
`)

				So(out, ShouldEqual, `LET foo = TRUE || FALSE

RETURN foo`)
			})

			Convey("Return", func() {
				f := fmt.New()

				out := f.MustFormat(`

RETURN      TRUE     OR      FALSE
`)

				So(out, ShouldEqual, `RETURN TRUE OR FALSE`)
			})
		})

		Convey("Ternary Operator", func() {
			Convey("Full", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo = TRUE ?     FALSE     : TRUE

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = TRUE ? FALSE : TRUE

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN      TRUE ?     FALSE     : TRUE
`)

					So(out, ShouldEqual, `RETURN TRUE ? FALSE : TRUE`)
				})
			})

			Convey("Short", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo = TRUE ?          : TRUE

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = TRUE ? : TRUE

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN      TRUE ?          : TRUE
`)

					So(out, ShouldEqual, `RETURN TRUE ? : TRUE`)
				})
			})
		})

		Convey("Regexp Operator", func() {
			Convey("Variable", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET     foo =     "foo"    !~     "[a-z]+bar$"

RETURN     foo
`)

				So(out, ShouldEqual, `LET foo = "foo" !~ "[a-z]+bar$"

RETURN foo`)
			})

			Convey("Return", func() {
				f := fmt.New()

				out := f.MustFormat(`

RETURN      "foo"     !~     "[a-z]+bar$"
`)

				So(out, ShouldEqual, `RETURN "foo" !~ "[a-z]+bar$"`)
			})
		})

		Convey("Equality Operator", func() {
			Convey("Variable", func() {
				f := fmt.New()

				out := f.MustFormat(`
LET     foo =     0     ==     NONE 

RETURN     foo
`)

				So(out, ShouldEqual, `LET foo = 0 == NONE

RETURN foo`)
			})

			Convey("Return", func() {
				f := fmt.New()

				out := f.MustFormat(`

RETURN      0      ==       NONE 
`)

				So(out, ShouldEqual, `RETURN 0 == NONE`)
			})
		})

		Convey("Array Operator", func() {
			Convey("IN", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =     1.5       IN       [ 2, 3, 1.5 ] 

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = 1.5 IN [2, 3, 1.5]

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    1.5     IN   [2, 3,1.5 ] 
`)

					So(out, ShouldEqual, `RETURN 1.5 IN [2, 3, 1.5]`)
				})

				Convey("NOT IN", func() {
					Convey("Variable", func() {
						f := fmt.New()

						out := f.MustFormat(`
LET     foo =     1.5   NOT    IN       [ 2, 3, 1.5 ] 

RETURN     foo
`)

						So(out, ShouldEqual, `LET foo = 1.5 NOT IN [2, 3, 1.5]

RETURN foo`)
					})

					Convey("Return", func() {
						f := fmt.New()

						out := f.MustFormat(`

RETURN    1.5  NOT   IN   [2, 3,1.5 ] 
`)

						So(out, ShouldEqual, `RETURN 1.5 NOT IN [2, 3, 1.5]`)
					})
				})

				Convey("ALL IN", func() {
					Convey("Variable", func() {
						f := fmt.New()

						out := f.MustFormat(`
LET     foo =    [ 1, 2, 3 ]     ALL     IN       [ 2, 3, 4 ]

RETURN     foo
`)

						So(out, ShouldEqual, `LET foo = [1, 2, 3] ALL IN [2, 3, 4]

RETURN foo`)
					})

					Convey("Return", func() {
						f := fmt.New()

						out := f.MustFormat(`

RETURN   [ 1, 2, 3   ]    ALL    IN    [ 2,     3,   4 ]
`)

						So(out, ShouldEqual, `RETURN [1, 2, 3] ALL IN [2, 3, 4]`)
					})
				})

				Convey("ANY IN", func() {
					Convey("Variable", func() {
						f := fmt.New()

						out := f.MustFormat(`
LET     foo =    [ 1, 2, 3 ]     ANY     IN       [ 2, 3, 4 ]

RETURN     foo
`)

						So(out, ShouldEqual, `LET foo = [1, 2, 3] ANY IN [2, 3, 4]

RETURN foo`)
					})

					Convey("Return", func() {
						f := fmt.New()

						out := f.MustFormat(`

RETURN   [ 1, 2, 3   ]    ANY    IN    [ 2,     3,   4 ]
`)

						So(out, ShouldEqual, `RETURN [1, 2, 3] ANY IN [2, 3, 4]`)
					})
				})

				Convey("NONE IN", func() {
					Convey("Variable", func() {
						f := fmt.New()

						out := f.MustFormat(`
LET     foo =    [ 1, 2, 3 ]     NONE     IN       [ 2, 3, 4 ]

RETURN     foo
`)

						So(out, ShouldEqual, `LET foo = [1, 2, 3] NONE IN [2, 3, 4]

RETURN foo`)
					})

					Convey("Return", func() {
						f := fmt.New()

						out := f.MustFormat(`

RETURN   [ 1, 2, 3   ]    NONE    IN    [ 2,     3,   4 ]
`)

						So(out, ShouldEqual, `RETURN [1, 2, 3] NONE IN [2, 3, 4]`)
					})
				})
			})

			Convey("Equality Operator", func() {
				Convey("ALL <Eq>", func() {
					Convey("Variable", func() {
						f := fmt.New()

						out := f.MustFormat(`
LET     foo =    [ 1, 2, 3 ]     ALL     ==       [ 2, 3, 4 ]

RETURN     foo
`)

						So(out, ShouldEqual, `LET foo = [1, 2, 3] ALL == [2, 3, 4]

RETURN foo`)
					})

					Convey("Return", func() {
						f := fmt.New()

						out := f.MustFormat(`

RETURN   [ 1, 2, 3   ]    ALL    !=    [ 2,     3,   4 ]
`)

						So(out, ShouldEqual, `RETURN [1, 2, 3] ALL != [2, 3, 4]`)
					})
				})

				Convey("ANY <Eq>", func() {
					Convey("Variable", func() {
						f := fmt.New()

						out := f.MustFormat(`
LET     foo =    [ 1, 2, 3 ]     ANY     >       [ 2, 3, 4 ]

RETURN     foo
`)

						So(out, ShouldEqual, `LET foo = [1, 2, 3] ANY > [2, 3, 4]

RETURN foo`)
					})

					Convey("Return", func() {
						f := fmt.New()

						out := f.MustFormat(`

RETURN   [ 1, 2, 3   ]    ANY    <    [ 2,     3,   4 ]
`)

						So(out, ShouldEqual, `RETURN [1, 2, 3] ANY < [2, 3, 4]`)
					})
				})

				Convey("NONE <Eq>", func() {
					Convey("Variable", func() {
						f := fmt.New()

						out := f.MustFormat(`
LET     foo =    [ 1, 2, 3 ]     NONE     ==       [ 2, 3, 4 ]

RETURN     foo
`)

						So(out, ShouldEqual, `LET foo = [1, 2, 3] NONE == [2, 3, 4]

RETURN foo`)
					})

					Convey("Return", func() {
						f := fmt.New()

						out := f.MustFormat(`

RETURN   [ 1, 2, 3   ]    NONE    !=    [ 2,     3,   4 ]
`)

						So(out, ShouldEqual, `RETURN [1, 2, 3] NONE != [2, 3, 4]`)
					})
				})
			})
		})

		Convey("Like Operator", func() {
			Convey("Plain", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =    "foo"    LIKE      "f*"

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = "foo" LIKE "f*"

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    "foo"     LIKE      "f*"
`)

					So(out, ShouldEqual, `RETURN "foo" LIKE "f*"`)
				})
			})

			Convey("With unary", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =     "foo"    NOT    LIKE "f*"

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = "foo" NOT LIKE "f*"

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    "foo"      NOT     LIKE      "f*"
`)

					So(out, ShouldEqual, `RETURN "foo" NOT LIKE "f*"`)
				})
			})
		})

		Convey("Multiplicative Operator", func() {
			Convey("Multi", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =    1 *    2

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = 1 * 2

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    1      *       2
`)

					So(out, ShouldEqual, `RETURN 1 * 2`)
				})
			})

			Convey("Div", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =    1     /    2

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = 1 / 2

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    1      /       2
`)

					So(out, ShouldEqual, `RETURN 1 / 2`)
				})
			})

			Convey("Mod", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =    1     %    2

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = 1 % 2

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    1      %       2
`)

					So(out, ShouldEqual, `RETURN 1 % 2`)
				})
			})
		})

		Convey("Additive Operator", func() {
			Convey("Plus", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =    1  +    2

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = 1 + 2

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    1      +      2
`)

					So(out, ShouldEqual, `RETURN 1 + 2`)
				})
			})

			Convey("Minus", func() {
				Convey("Variable", func() {
					f := fmt.New()

					out := f.MustFormat(`
LET     foo =    1     -    2

RETURN     foo
`)

					So(out, ShouldEqual, `LET foo = 1 - 2

RETURN foo`)
				})

				Convey("Return", func() {
					f := fmt.New()

					out := f.MustFormat(`

RETURN    1      -       2
`)

					So(out, ShouldEqual, `RETURN 1 - 2`)
				})
			})
		})
	})

	Convey("Grouped Operators", t, func() {
		Convey("Math Operators", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET     foo = (   1    +  2   )    *   2

RETURN     foo
`)

			So(out, ShouldEqual, `LET foo = (1 + 2) * 2

RETURN foo`)
		})

		Convey("Ternary Operators", func() {
			f := fmt.New()

			out := f.MustFormat(`
LET     foo = (1 +  1   ) &&     ( TRUE && FALSE   ) ? TRUE    : FALSE

RETURN     foo
`)

			So(out, ShouldEqual, `LET foo = (1 + 1) && (TRUE && FALSE) ? TRUE : FALSE

RETURN foo`)
		})
	})
}
