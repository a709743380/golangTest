// main.go
package main

import (
	"fmt"
	"log"
	db "main/MSDB"
	"main/controller"
	"main/repository"
	"main/service"
	"net/http"
)

func main() {
	// 使用配置文件連接資料庫
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("初始化資料庫連接失敗: %v", err)
	}
	defer dbConn.Close()
	// 按順序注入 開始
	// bankRepository := &repository.BankRepository{
	// 	DB: dbConn, // 注入連線
	// }
	// bankService := &service.BankServiceImpl{
	// 	Repository: bankRepository, // 注入 Repository
	// }
	// controller := &controller.BankController{
	// 	BankService: bankService, // 注入服務
	// }
	// 按順序注入 結束

	//创建各层实例
	bankRepository := repository.NewBankRepository(dbConn)
	bankService := service.NewBankServiceImpl(bankRepository)
	controller := controller.NewBankController(bankService)

	// 設定路由
	http.HandleFunc("/", controller.GetBankData)
	// 設定路由

	// 啟動 HTTP 伺服器，監聽 8080 埠口
	fmt.Println("伺服器正在啟動，監聽埠口 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("啟動伺服器失敗: %v", err)
	}
}
