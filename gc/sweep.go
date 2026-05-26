package gc

import (
	"fmt"

	"github.com/jsndz/gorbage/lang"
)

func Sweep(r *lang.Runtime) {
	obj := r.FirstObject

	for obj != nil {
		if !obj.Marked {
			unreached := obj
			fmt.Println("sweeping object")
			obj = unreached.Next

			unreached = nil
		} else {
			fmt.Println("sweeping object")
			obj.Marked = false
			obj = obj.Next
		}
	}
}
