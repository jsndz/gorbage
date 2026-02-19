package lang

// this is for imitating the a program running / a virtual machine
// The Runtime’s role in this story is to have a stack that stores the variables that are currently in scope
// Runtime can be stack based or register based

// create a Runtime
// and provide methods
const MAX_SIZE = 256

type Runtime struct {
	stackSize int
	stack     []*Object
}

func NewRuntime() *Runtime {
	return &Runtime{
		stackSize: 0,
		stack:     make([]*Object, MAX_SIZE),
	}
}

func (Runtime *Runtime) Push(obj *Object) bool {
	Runtime.stack[Runtime.stackSize] = obj
	Runtime.stackSize++
	return true
}

func (Runtime *Runtime) Pop() *Object {
	var obj *Object = Runtime.stack[Runtime.stackSize]
	Runtime.stackSize--
	return obj
}
