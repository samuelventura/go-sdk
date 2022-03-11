package main

func main() {
	//fatal error: all goroutines are asleep - deadlock!
	//goroutine 1 [select (no cases)]:
	// select {}

	//panic: in go routine
	//exit status 2
	// go func() { panic("in go routine") }()

	//fatal error: all goroutines are asleep - deadlock!
	//goroutine 1 [chan receive]:
	// o := make(chan interface{})
	// c := make(chan interface{})
	// go func() {
	// 	defer func() { o <- "exit" }()
	// 	for d := range c {
	// 		fmt.Println(d)
	// 	}
	// }()
	// fmt.Println(<-o)
}
