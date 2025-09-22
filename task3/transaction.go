package task3

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Accounts struct {
	ID      uint64
	Balance float64
	UserID  uint64
}

type Transaction struct {
	ID            uint64
	FromAccountId uint64
	ToAccountId   uint64
	Amount        float64
}

func RunTransaction(db *gorm.DB) {

	//a := Accounts{}
	//a.Balance = 1000
	//a.UserID = 10
	//b := Accounts{}
	//b.Balance = 100
	//b.UserID = 20
	//
	//db.AutoMigrate(&Accounts{}, &Transaction{})
	//
	//db.Create(&a)
	//db.Create(&b)

	err := db.Transaction(func(tx *gorm.DB) error {

		ua := Accounts{}
		uar := tx.Find(&ua, "user_id", 10)
		ub := Accounts{}
		ubr := tx.Find(&ub, "user_id", 20)
		fmt.Println(ua)
		fmt.Println(ub)
		if uar.Error != nil || ua.Balance < 100 {
			return errors.New("a 账户余额不足")
		}
		if ubr.Error != nil {
			return errors.New("b 账户不存在")
		}

		ua.Balance -= 100
		tx.Save(&ua)
		ub.Balance += 100
		tx.Save(&ub)
		fmt.Println("__")
		tran := Transaction{}
		tran.FromAccountId = ua.ID
		tran.ToAccountId = ub.ID
		tran.Amount = 100
		tx.Create(&tran)

		return errors.New("异常处理 回滚了吗")
	})
	if err != nil {
		fmt.Println(err)
		return
	}

}
