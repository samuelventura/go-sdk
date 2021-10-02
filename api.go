package go_sdk

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

type Context = context.Context
type Logger = func(string, ...Any)
type Output = func(...Any)
type Map = map[string]Any
type Queue = chan Action
type Channel = chan Any
type Any = interface{}
type Action = func()
type Dispatch = func(Mutation)
type Factory = func(Context) Dispatch

type Log interface {
	Prefixed(prefix ...Any) Log
	Log(string, ...Any)
	Trace(...Any)
	Debug(...Any)
	Info(...Any)
	Warn(...Any)
	Error(...Any)
}

func NopAction()                  {}
func NopOutput(...Any)            {}
func NopDispatch(Mutation)        {}
func NopFactory(Context) Dispatch { return NopDispatch }

type Mutation struct {
	Origin string
	Name   string
	Args   Any
}

func M(name string, origin string, args Any) Mutation {
	return Mutation{Name: name, Origin: origin, Args: args}
}

func Mn(name string) Mutation {
	return Mutation{Name: name}
}

func Mns(name string, origin string) Mutation {
	return Mutation{Name: name, Origin: origin}
}

func Mna(name string, args Any) Mutation {
	return Mutation{Name: name, Args: args}
}

func Mnsa(name string, origin string, args Any) Mutation {
	return Mutation{Name: name, Origin: origin, Args: args}
}

func (m Mutation) String() string {
	buf := new(strings.Builder)
	var typ string
	switch m.Args.(type) {
	case nil:
		typ = "<nil>"
	default:
		typ = reflect.TypeOf(m.Args).String()
	}
	fmt.Fprintf(buf, "{%s,%s,%s,%v}", m.Name, m.Origin, typ, m.Args)
	return buf.String()
}
