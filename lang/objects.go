package lang

// strong enum
type ObjectType int

const (
	OBJ_INT  ObjectType = 0
	OBJ_PAIR ObjectType = 1
)

type Object struct {
	oType    ObjectType
	IntValue int
	Head     *Object
	Tail     *Object
}

func (r *Runtime) NewObject(t ObjectType) *Object {
	return &Object{
		oType: t,
	}

}

func (r *Runtime) PushInt(val int) {
	obj := r.NewObject(OBJ_INT)
	obj.IntValue = val
	r.Push(obj)
}

func (r *Runtime) PushPair(val int) *Object {
	obj := r.NewObject(OBJ_PAIR)
	obj.Tail = r.Pop()
	obj.Head = r.Pop()

	r.Push(obj)
	return obj
}

//  If we had a parser and an interpreter that called those functions, we’d have an honest to God language on our hands.
// And, if we had infinite memory, it would even be able to run real programs
