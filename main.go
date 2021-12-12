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

	d4_1 := Day4Exec("inputs/4.txt")
	fmt.Printf("Day 4 part 1 answer is %d\r\n", d4_1)
	d4_2 := Day4ExecII("inputs/4.txt")
	fmt.Printf("Day 4 part 2 answer is %d\r\n", d4_2)

	d5_1 := Day5Exec("inputs/5.txt")
	fmt.Printf("Day 5 part 1 answer is %d\r\n", d5_1)
	d5_2 := Day5ExecII("inputs/5.txt")
	fmt.Printf("Day 5 part 2 answer is %d\r\n", d5_2)

	d6_1 := Day6Exec("inputs/6.txt", 80)
	fmt.Printf("Day 6 part 1 answer is %d\r\n", d6_1)
	d6_2 := Day6Exec("inputs/6.txt", 256)
	fmt.Printf("Day 6 part 2 answer is %d\r\n", d6_2)

	d7_1 := Day7Exec("inputs/7.txt")
	fmt.Printf("Day 7 part 1 answer is %d\r\n", d7_1)
	d7_2 := Day7ExecII("inputs/7.txt")
	fmt.Printf("Day 7 part 2 answer is %d\r\n", d7_2)

	d8_1 := Day8Exec("inputs/8.txt")
	fmt.Printf("Day 8 part 1 answer is %d\r\n", d8_1)
	d8_2 := Day8ExecII("inputs/8.txt")
	fmt.Printf("Day 8 part 2 answer is %d\r\n", d8_2)

	d9_1 := Day9Exec("inputs/9.txt")
	fmt.Printf("Day 9 part 1 answer is %d\r\n", d9_1)
	d9_2 := Day9ExecII("inputs/9.txt")
	fmt.Printf("Day 9 part 2 answer is %d\r\n", d9_2)
}
