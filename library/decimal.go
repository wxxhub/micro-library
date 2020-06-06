package library

import (
	"github.com/shopspring/decimal"
)

func Yuan2penny(from int64) string {
	d := decimal.New(1,2)
	return decimal.NewFromInt(from).DivRound(d,2).String()
}

func Penny2Yuan(from string) int64 {
	d := decimal.New(1,2)//分转元乘以100
	d1 := decimal.New(1,0)//乘完之后，保留2为小数，需要这么一个中间参数
	//当乘以100后，仍然有小数位，取四舍五入法后，再取整数部分
	return decimal.NewFromFloat(price).Mul(d).DivRound(d1,0).IntPart()
}
