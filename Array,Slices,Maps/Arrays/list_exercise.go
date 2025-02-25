package lists

import (
	"fmt"
)

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"Bouldering", "Painting", "Taekwondo"}
	fmt.Println(hobbies, "======== Task 1")

	fmt.Println(hobbies[0], "===== Task 2.1")
	new_hobbies := hobbies[1:3]
	fmt.Println(new_hobbies, "======== Task 2.2")

	hobbies_slice1 := hobbies[:2]
	hobbies_slice2 := hobbies[0:2]
	fmt.Println(hobbies_slice1, hobbies_slice2, "======== Task 3")

	hobbies_slice3 := hobbies_slice2[1:3]
	fmt.Println(hobbies_slice3, "======== Task 4")

	var course_goal = []string{"Learn Go-Lang", "Build 2D Video Game"}
	fmt.Println(course_goal, "======= Task 5")

	course_goal[1] = "Paint a Mural"
	course_goal = append(course_goal, "Learn Blockchain")
	fmt.Println(course_goal, "======= Task 6")

	dynamic_hobbies := []Product{{"title", 1, 3.4}, {"title2", 2, 4.4}}
	dynamic_hobbies = append(dynamic_hobbies, Product{"title3", 3, 5.4})
	fmt.Println(dynamic_hobbies, "====== Task 7")

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
