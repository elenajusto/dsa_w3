package main

import "log"

func main() {
	my_array := []int{64, 34, 25, 12, 22, 11, 90, 5}

	log.Println("initial array:", my_array)

	for index := range len(my_array) {

		for sub_index := range len(my_array) - index - 1 {

			// Establish the values
			check_1 := &my_array[sub_index]
			check_2 := &my_array[sub_index+1]

			// Check the values
			if *check_1 > *check_2 {
				// Swap if conditon met
				log.Print("bubbling ", *check_1)
				*check_1, *check_2 = *check_2, *check_1
			}
			log.Println("array after iteration:", my_array)
		}
	}

	log.Println("final array:", my_array)
}
