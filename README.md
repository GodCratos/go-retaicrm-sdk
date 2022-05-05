# package main

# import (
#	"log"
#
#	retailcrm "github.com/godcratos/go-retaicrm-sdk"
# )

# func main() {
# 	client := retailcrm.NewClient("ecobar2.retailcrm.ru", "m1DeSfOYo71qlc6HLPs1YFdcX6mUd6TT", true)
# 	filter := retailcrm.OrdersRequest{
# 		Filter: retailcrm.OrdersFilter{
# 			Ids: []int{
# 				156138,
# 			},
# 		},
# 		Page: 1,
# 	}
# 	orders, _, err := client.OrdersGet(filter)
# 	if err != nil {
# 		log.Println(err)
# 		return
# 	}
# 	log.Println(orders)
# }
