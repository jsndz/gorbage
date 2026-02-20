package gc

import "github.com/jsndz/gorbage/lang"

func MarkAll(r *lang.Runtime) {
	for i := 0; i < r.StackSize; i++ {
		Mark(r.Stack[i])
	}
}

func Mark(obj *lang.Object) {
	if obj == nil {
		return
	}
	if obj.Marked {
		return
	}
	obj.Marked = true
	if obj.OType == lang.OBJ_PAIR {
		Mark(obj.Head)
		Mark(obj.Tail)
	}

}
