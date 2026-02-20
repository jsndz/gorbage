package gc

import "github.com/jsndz/gorbage/lang"

func Sweep(r *lang.Runtime) {
	obj := r.FirstObject

	for obj != nil {
		if !obj.Marked {
			unreached := obj

			obj = unreached.Next

			unreached = nil
		} else {
			obj.Marked = false
			obj = obj.Next
		}
	}
}
