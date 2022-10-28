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

	update := make(map[string]interface{})
	update["Use"] = 310
	stor.UpdateStudent("Cody", update)

	fmt.Println(stor.Student("Cody"))

	stor.DB.Drop(stor.Ctx)

}
