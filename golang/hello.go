package main

import (
	"fmt"
	// "os"
)

type Worker struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello, World!")

	power, powValue := getPower("400", 2000)
	fmt.Printf("%s %d\n", power, powValue)

	var name string
	name = "frank"
	fmt.Println(name)

	// if len(os.Args) != 2 {
	// 	os.Exit(1)
	// }
	// fmt.Println(os.Args[1])

	worker := Worker{
		Name: "Frank",
		Age:  30,
	}

	worker.Age = 31
	fmt.Println(worker)

	w := &Worker{
		Name: "Super",
		Age:  50,
	}
	SuperWorker(w)
	w.WxWorker()
	fmt.Println(w.Age)
	forExample()
}

func (w *Worker) WxWorker() {
	w.Age += 4000
}

func SuperWorker(wx *Worker) {
	wx.Age += 150
}

func getPower(a string, b int) (string, int) {
	if a == "400" || a == "200" {
		return a, b
	} else {
		return "default", 0
	}
}

func forExample() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// for i := range [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
	// 	fmt.Println(i)
	// }

	m := map[string]int{"k1": 1, "k2": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
