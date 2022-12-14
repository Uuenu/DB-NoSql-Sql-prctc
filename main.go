package main

import (
	"fmt"
	"test-go/storage/mongodb"
)

func main() {
	stor := mongodb.New()

	stor.SaveStudent("Cody", "Van Goth", "m3233", "codyvangoth@gmail.com", 291, 1999, false)
	stor.SaveStudent("Body", "Pack", "m2312", "hello@world.com", 271, 2001, true)
	stor.SaveStudent("Liza", "Dedusheva", "m3233", "lizadedusheva@gmail.com", 310, 1999, false)
	stor.SaveStudent("Isabel", "Hello", "m3233", "isabelhello@gmail.com", 263, 2000, true)

	updateStd := stor.Student("Cody")
	updateStd.Use = 310
	stor.UpdateStudent(updateStd)

	fmt.Println(stor.TableStudents())

	stor.DB.Drop(stor.Ctx)

}
