package main

import "log"

func main() {
	my_array := []int{64, 34, 25, 12, 22, 11, 90, 5}

	log.Println("initial array:", my_array)

	for index := range len(my_array) {

		log.Println("index being checked: ", index)

		for sub_index := range len(my_array) - index - 1 {

			log.Print("sub_index being checked: ", sub_index, " with value: ", my_array[sub_index])

			// Establish the values
			check_1 := &my_array[sub_index]
			check_2 := &my_array[sub_index+1]

			// Check the values
			log.Println("comparing sub index ", sub_index, " with value ", *check_1, " with sub index ", sub_index+1, " with value ", *check_2)
			if *check_1 > *check_2 {
				// Swap if conditon met
				log.Print("preforming swap")
				*check_1, *check_2 = *check_2, *check_1
			}

			log.Println("array after iteration:", my_array)
		}
	}

	log.Println("final array:", my_array)
}
