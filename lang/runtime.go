package lang

// this is for imitating the a program running / a virtual machine
// The Runtime’s role in this story is to have a Stack that stores the variables that are currently in scope
// Runtime can be Stack based or register based

// create a Runtime
// and provide methods
const MAX_SIZE = 256

type Runtime struct {
	StackSize   int
	Stack       []*Object
	FirstObject *Object
}

func NewRuntime() *Runtime {
	return &Runtime{
		StackSize:   0,
		Stack:       make([]*Object, MAX_SIZE),
		FirstObject: nil,
	}
}

func (Runtime *Runtime) Push(obj *Object) bool {
	Runtime.Stack[Runtime.StackSize] = obj
	Runtime.StackSize++
	return true
}

func (Runtime *Runtime) Pop() *Object {
	var obj *Object = Runtime.Stack[Runtime.StackSize]
	Runtime.StackSize--
	return obj
}
