package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("оплата криптовалютой")
	fmt.Println("сумма:", usd, "usdt")

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Println("отмена крипто операции id=", id)
}
