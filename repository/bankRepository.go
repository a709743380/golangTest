package repository

import (
	"database/sql"
	"fmt"
	model "main/model"
	"reflect"
)

type BankRepository struct {
	DB *sql.DB
}

func NewBankRepository(db *sql.DB) *BankRepository {
	return &BankRepository{DB: db}
}

func (d *BankRepository) GetBankData() ([]model.Bank, error) {
	// SQL 插入语句
	query := `
SELECT TOP(100)
	[BankCode]
	,[BankName]
	,[ShortName]	
FROM [Test].[dbo].[BankData]`

	// 执行插入操作
	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	defer rows.Close()

	var banks []model.Bank
	columns, _ := rows.Columns() // 获取查询结果的列名
	values := make([]interface{}, len(columns))

	for rows.Next() {

		var eachData model.Bank
		for i := range values {
			values[i] = reflect.ValueOf(&eachData).Elem().Field(i).Addr().Interface()
		}
		// 扫描到结构体字段
		if err := rows.Scan(values...); err != nil {
			return nil, fmt.Errorf("扫描数据失败: %v", err)
		}

		banks = append(banks, eachData)
	}

	return banks, nil
}
