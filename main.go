package main

import "fmt"

func main() {
	d1_1 := Day1Exec("inputs/1.txt")
	fmt.Printf("Day 1 part 1 answer is %d\r\n", d1_1)
	d1_2 := Day1ExecII("inputs/1.txt")
	fmt.Printf("Day 1 part 2 answer is %d\r\n", d1_2)

	d2_1 := Day2Exec("inputs/2.txt")
	fmt.Printf("Day 2 part 1 answer is %d\r\n", d2_1)
	d2_2 := Day2ExecII("inputs/2.txt")
	fmt.Printf("Day 2 part 2 answer is %d\r\n", d2_2)

	d3_1 := Day3Exec("inputs/3.txt")
	fmt.Printf("Day 3 part 1 answer is %d\r\n", d3_1)
	d3_2 := Day3ExecII("inputs/3.txt")
	fmt.Printf("Day 3 part 2 answer is %d\r\n", d3_2)
}
