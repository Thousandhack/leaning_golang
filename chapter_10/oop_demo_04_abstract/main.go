package main

import "fmt"

// 定义一个结构体Account
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 方法
// 1.存款
func (account *Account) SaveMoney(money float64, pwd string) {
	// 检验输入密码
	if pwd != account.Pwd {
		fmt.Println("输入密码错误")
		return
	}
	if money <= 0 {
		fmt.Println("输入金额错误")
		return
	}
	account.Balance += money
	fmt.Printf("存款%v成功！",money)
}

// 2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	// 检验输入密码
	if pwd != account.Pwd {
		fmt.Println("输入密码错误")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("输入金额错误")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功！")
}


// 3. 查询余额
func (account *Account) Query(pwd string) {
	// 检验输入密码
	if pwd != account.Pwd {
		fmt.Println("输入密码错误")
		return
	}
	fmt.Printf("查询为：账号：%v,余额：%v \n", account.AccountNo, account.Balance)
}
func main() {

	// 实例化
	account := Account{
		AccountNo: "no1111",
		Pwd:       "666",
		Balance:   100,
	}

	//
	account.SaveMoney(10, "666")
	account.Query("666")
	account.WithDraw(20,"666")
	account.Query("666")
}
