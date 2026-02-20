package main

import (
	"github.com/jsndz/gorbage/gc"
	"github.com/jsndz/gorbage/lang"
)

func GC(r *lang.Runtime) {
	gc.MarkAll(r)
	gc.Sweep(r)
}

func main() {
	r := lang.NewRuntime()
	obj1 := r.NewObject(lang.OBJ_INT)
	obj2 := r.NewObject(lang.OBJ_PAIR)
	r.Push(obj1)
	r.Push(obj2)
	GC(r)
}
