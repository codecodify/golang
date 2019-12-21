package main

import "wechat/initRouter"

func main() {
	router := initRouter.SetRouter()

	_ = router.Run(":8080")
}
