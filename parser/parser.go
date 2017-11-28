package parser

import (
	"fmt"
	"math/big"
	"strings"
)

//Convert hexadecimal strings into integers
func Hex2Int(hexStr string) *big.Int {
	//remove 0x prefix
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	//parse string to int
	result := new(big.Int)
	result, _ = result.SetString(cleaned, 16)

	//For debug purposes take this out later
	fmt.Println(result)
	
	return result
}

//Adjusts number to account for decimals
func AdjustForPrecision(i *big.Int, precision int64) *big.Float {
	//Convert int to float
	f1 := new(big.Float).SetInt(i)

	//For debug purposes take this out later
	fmt.Println(f1)

	//Convert precision into exponentiated int
	i2 := new(big.Int).Exp(big.NewInt(10), big.NewInt(precision), nil)

	//Convert precision int to float
	f2 := new(big.Float).SetInt(i2)

	//Divide unadjusted number by precision to get true val
	f1.Quo(f1, f2)

	//For debug purposes take this out later
	fmt.Println(f1)

	return f1
}

//Converts time in hours into approximate number of blocks
func Hours2Block(time int) (int) {
	seconds := time * 60 * 60
	blocks := seconds / 14
	return blocks
}