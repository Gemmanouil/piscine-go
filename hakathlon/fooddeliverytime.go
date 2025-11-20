package piscine

type food struct {
	preptime int
	name     string
}

func FoodDeliveryTime(order string) int {
	burger := food{preptime: 15, name: "burger"}
	nuggets := food{preptime: 12, name: "nuggets"}
	chips := food{preptime: 10, name: "chips"}
	menu := []food{burger, nuggets, chips}
	for _, item := range menu {
		if item.name == order {
			return 404
		}
	}
	return 404
}
