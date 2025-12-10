package piscine

import "fmt"

func main() {
	// command_and_arguments := os.Args
	// fmt.Println(len(command_and_arguments))
	// // var arr []string
	// if len(command_and_arguments) < 1 || len(command_and_arguments)>1 {
	// 	fmt.Println("error")
	// } else {
	// 	fmt.Println(command_and_arguments)
	// }
	var slice1 []int
	var n []int
	slice1 = []int{1, 2, 3, 4, 7}
	slice2 := []int{6, 7, 678}
	fmt.Println(slice1)
	fmt.Println(slice2)
	n = append(slice1, 6)
	//for i,f:=range slice2{
	n = append(n, slice2...)
	//}
	//n = append(n,7)
	//n = append(n,678)
	fmt.Println(n)
}
