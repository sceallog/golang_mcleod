package main

//● “defer” multiple functions in main
//	○ show that a deferred func runs after the func containing it exits.
//	○ determine the order in which the multiple defer funcs run

func main() {
	defer foo("This is foo")
	bar()
	baz()
	defer biz()
	buz()
}

func foo(text string) {
	println(text)
}
func bar() {
	defer foo("This is a deferred foo inside bar")
	println("This is bar")
}
func baz() {
	println("This is baz")
}
func biz() {
	println("This is biz")
}
func buz() {
	println("This is buz")
}
