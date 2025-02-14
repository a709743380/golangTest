package controller

import (
	"fmt"
	"html/template"
	_bankService "main/service"
	"net/http"
)

type BankController struct {
	BankService _bankService.BankService
}

func NewBankController(service _bankService.BankService) *BankController {
	return &BankController{BankService: service}
}

func (c *BankController) GetBankData(w http.ResponseWriter, r *http.Request) {

	bankdata, err := c.BankService.GetBankData()
	if err != nil {
		http.Error(w, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	t, _ := handleLayout("templates/BankData.html")

	// 渲染模板并传递数据
	result := map[string]interface{}{
		"Title": "银行数据页面",
		"Model": bankdata,
	}

	err = t.ExecuteTemplate(w, "layout", result)
	if err != nil {
		http.Error(w, fmt.Sprintf("渲染模板失败: %v", err), http.StatusInternalServerError)
	}
}

func handleLayout(targethtml string) (*template.Template, error) {
	// 解析 layout 和 home 模板
	// 调用 template.ParseFiles 返回模板和错误
	tmpl, err := template.ParseFiles("templates/Layout/layoutBank.html", targethtml)
	if err != nil {
		return nil, fmt.Errorf("解析模板失败: %w", err)
	}

	return tmpl, nil
}
