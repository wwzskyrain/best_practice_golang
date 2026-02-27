package main

func getName() string {
	return "Yueyi"
}

func main() {
	name := getName()
	nameAddr := &name

	println(name)
	println(&name)
	println(nameAddr)
	println(*nameAddr)
}
