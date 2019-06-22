package test

import (
	"github.com/shopspring/decimal"
	"time"
)

//红包商品表,默认有符号类型字段
type GoodsSigned struct {
	Goods
}

//红包商品表,无符号类型字段
type GoodsUnsigned struct {
	Goods
}

/*
例子中金额字段使用了decimal类型，这里使用的decimal类型是由下面这个第三方库提供的：github.com/shopspring/decimal
可以使用 go get 进行安装：
go get -u github.com/shopspring/decimal

使用decimal的目的是让浮点数字计算精确，go语言中内置的类型float32和float64在计算时会丢失精度，在金额计算中非常不严谨的。
比如下面这段代码，结果输出是0.09999999999999998:

var a float64 = 0.9
c := float64(1) - a
fmt.Println(c)
*/
//红包商品
type Goods struct {
	RemainAmount   decimal.Decimal `db:"remain_amount"`        //红包剩余金额额
	RemainQuantity int             `db:"remain_quantity"`      //红包剩余数量
	EnvelopeNo     string          `db:"envelope_no,uni"`      //红包编号,红包唯一标识
	CreatedAt      time.Time       `db:"created_at,omitempty"` //创建时间
	UpdatedAt      time.Time       `db:"updated_at,omitempty"` //更新时间
}
