package main

import "fmt"

type Behavior interface {
	Eat()
	Run()
}

type Person struct {
}

func (Person) Eat() {
	println("eat")
}

func (Person) Run() {
	println("run")
}

type Inter interface {
	Ping()
	Pang()
}

type St struct {
}

func (St) Ping() {
	println("ping")
}

func (*St) Pang() {
	println("pang")
}

// Test01StWithDesignOrJustVar St类型变量，无论赋值或者不复制，都可以调用方法，而且可以调用St和*St的方法。 St类型变量不能赋值以nil
func Test01StWithDesignOrJustVar() {
	var st St = St{}
	st.Ping()
	st.Pang()

	var stVar St
	stVar.Ping()
	stVar.Pang()

	// var stWithNil St = nil         //会编译报错，因为St类型变量不能赋值为nil
	// fmt.Printf("stWithNil = %+v", stWithNil)
}

// Test02StPointerWithAssignOrOnlyVarOrAssignNil Test02StPointerWithAssignOrOnlyVar 可以看出来
// 1.指针类型的变量，无论不赋值、赋值为nil、赋值为空壳结构体，都可以调用其*St的方法，不同之处在于如果是一个空壳，则可以调用St的方法。
// 2.未赋值的指针类型变量，调用结构体的方法，会报空指针。只有这一种类型的空指针
func Test02StPointerWithAssignOrOnlyVarOrAssignNil() {
	var stP *St = &St{}
	stP.Ping()
	stP.Pang()

	var stPointerWithOnlyVar *St
	// stPointerWithOnlyVar.Ping() // 未实现Ping方法，所以会报空指针——nil pointer dereference
	stPointerWithOnlyVar.Pang()

	var stPointerWithNil *St = nil
	// stPointerWithNil.Ping() // 未实现Ping方法，所以会报空指针——nil pointer dereference
	stPointerWithNil.Pang()
}

// Test03TypeStAndStarSt 类型判断时，St与*St是两个不同的类型
func Test03TypeStAndStarSt() {
	i2 := getInterface(-1)
	// i2 := GetInterface(1) //确实是两个不同类型的变量类型
	switch v := i2.(type) {
	case St:
		fmt.Printf("st, v = %+v", v)
	case *St:
		fmt.Printf("*st, v = %+v", v)
	default:
		fmt.Printf("default, v = %+v", v)
	}
}

func getInterface(i int) interface{} {
	if i < 0 {
		return St{}
	} else {
		return &St{}
	}
}

func TestOrigin1() {
	var stP *St = nil
	var it Inter = stP // 这里it的动态类型是*St，但是动态值是nil
	fmt.Printf("%p\n", stP)
	fmt.Printf("%p\n", it)

	var itOnlyVar Inter

	itOnlyVar.Ping()

	if it != nil {
		it.Pang()
		// 调用ping方法会panic的。方法转化为函数调用，第一个参数是St类型，但是*St是nil，无法获取指针所指的对象的值，所以导致panic
		// 【王卫振注】而如果变量是St类型，而非指针类型，所以不存在
		// it.Ping()
	}
}

// Test04Interface 发现：
// 1.接口类型变量，如果不赋值，则会空指针。
// 2.接口类型变量能赋以什么值呢
func Test04Interface() {

	var itOnlyVar Inter
	fmt.Printf("itOnlyVar = %p\n", itOnlyVar)
	// itOnlyVar.Pang() // 这段代码会panic，因为itOnlyVar的动态类型是nil，而nil没有Pang方法

	var itAssignWithNil Inter = nil
	fmt.Printf("itAssignWithNil = %p\n", itAssignWithNil)
	// itAssignWithNil.Pang() // 这段代码会panic，因为itAssignWithNil的动态类型是nil，而nil没有Pang方法

	// var itAssignWithStNil Inter = (*stP)(nil) //nil可以转化为*St的形式吗？
	// fmt.Printf("itAssignWithCastNil = %p\n", itAssignWithNil)
	// itAssignWithNil.Pang() // 这段代码会panic，因为itAssignWithNil的动态类型是nil，而nil没有Pang方法

	var stP *St = nil
	var it Inter = stP // 这里it的动态类型是*St，但是动态值是nil
	fmt.Printf("stP = %p\n", stP)
	fmt.Printf("it = %p\n", it)
	// 	stP的调用我们前面已经盘过一遍， 而it则和它一样的，只能调用Pang方法，而不能调用Ping方法，会空指针
	it.Pang()
	// it.Ping() // 这段代码会panic，因为it的动态类型是*St，而*St没有Ping方法

	var itToSt interface{} = St{}
	fmt.Printf("itToSt = %+v", itToSt)
}

func TestInterfaceToAssign() {
	var b Behavior = Person{}
	fmt.Printf("b = %+v \n", b)

	var bPointer Behavior = &Person{}
	fmt.Printf("bPointer = %+v  %p\n", bPointer, bPointer)
}

func GenSt(i int) St {
	if i > 0 {
		return St{}
	} else {
		return St{}
		// return nil //这编译是不通过的。
	}
}

func TestWWZ01() {
	var it Inter = nil
	if it != nil {
		it.Ping()
	}

	var st *St = nil
	// var i Inter = st
	var i interface{} = st
	if i != nil {
		i.(*St).Pang()
		// i.Pang()
	}

	st.Pang()

	// 神奇，即使不辅助，也能调用。哈哈，原来非接口变量没有nil空值的概念
	var st2 St
	st2.Pang()
	st2.Ping()
}

func main() {
	// Test01StWithDesignOrJustVar()
	// Test02StPointerWithAssignOrOnlyVarOrAssignNil()
	// Test03TypeStAndStarSt()
	// Test04Interface()
	TestInterfaceToAssign()
}
