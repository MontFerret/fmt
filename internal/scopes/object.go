package scopes

//
//import (
//	"github.com/MontFerret/fmt/internal/core"
//)
//
//type (
//	ObjectScopeItem struct {
//		name        string
//		isShorthand bool
//	}
//
//	ObjectScope struct {
//		buff        []core.Token
//		maxLineLen  uint64
//		currLineLen uint64
//	}
//)
//
//func NewObjectScope(maxLineLen uint64) Scope {
//	return &ObjectScope{
//		buff:       make([]core.Token, 0, 10),
//		maxLineLen: maxLineLen,
//	}
//}
//
//func (i *ObjectScopeItem) String() string {
//	if i.isShorthand {
//		return i.name
//	}
//
//	return i.name + ": "
//}
//
//func (s *ObjectScope) Start(stream *core.Stream) Scope {
//	stream.Write("{")
//
//	return s
//}
//
//func (s *ObjectScope) Write(item core.Token) Scope {
//	s.buff = append(s.buff, item)
//	s.currLineLen += uint64(item.Len()) + 2 // len + comma + white space
//
//	return s
//}
//
//func (s *ObjectScope) End(stream *core.Stream) Scope {
//	//size := len(s.buff)
//	//last := size - 1
//	//
//	//if s.bracketSpacing && size > 0 {
//	//	s.out.Write(" ")
//	//}
//	//
//	//for i, item := range s.buff {
//	//	s.out.Write(item.String())
//	//
//	//	if i != last {
//	//		if _, ok := item.(*ObjectScopeItem); !ok {
//	//			s.out.Write(", ")
//	//		}
//	//	}
//	//}
//	//
//	//if s.bracketSpacing && size > 0 {
//	//	s.out.Write(" ")
//	//}
//	//
//	//s.out.WriteString("}")
//
//	return s
//}
