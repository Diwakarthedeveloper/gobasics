/*An interface is a custom type that is used to specify a set of one or more method signatures and the
interface is abstract, so you are not allowed to create an instance of the interface. But you are allowed
to create a variable of an interface type and this variable can be assigned with a concrete type value that
 has the methods the interface requires. Or in other words, the interface is a behavior that a type can do.
*/

// Defining interface
type MyInterface interface {
	Method1(param1, param2 int) int
	Method2(param1 string) string
}

// Defining a struct
type MyStruct struct {
	// fields
}

// Implementing interface methods on struct
func (m MyStruct) Method1(param1, param2 int) int {
	// implementation
}

func (m MyStruct) Method2(param1 string) string {
	// implementation
}
