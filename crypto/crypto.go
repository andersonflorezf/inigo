package crypto

import "fmt"

var (
	bitcoinPrice  float32
	ethereumPrice float32
	litecoinPrice float32
)

func init() {
	bitcoinPrice = 45000.32
	ethereumPrice = 218.9
	litecoinPrice = 231.1
}

// convierte de dolares a bitcoin
func ToBitcoin(dolars float32) {
	bitcoins := dolars / bitcoinPrice
	fmt.Printf("%.2f USD equivalen a ~%.2f Bitcoins", dolars, bitcoins)
}

// convierte de dolares a ethereum
func ToEthereum(dolars float32) {
	ethereum := dolars / ethereumPrice
	fmt.Printf("%.2f USD equivalen a ~%.2f Ethereums", dolars, ethereum)
}

// convierte de dolares a litecoin
func ToLitecoin(dolars float32) {
	litecoin := dolars / litecoinPrice
	fmt.Printf("%.2f USD equivalen a ~%.2f Litecoin", dolars, litecoin)
}
