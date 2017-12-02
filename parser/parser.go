package parser

import (
	"fmt"
	"math/big"
	//"strings"
)

//Convert hexadecimal strings into integers
func Hex2Int(hexStr string) (*big.Int) {
	//remove 0x prefix
	//cleaned := strings.Replace(hexStr, "0x", "", -1)

	//parse string to int
	result := new(big.Int)
	result, _ = result.SetString(hexStr, 16)

	//For debug purposes take this out later
	fmt.Println(result)
	
	return result
}

//Adjusts number to account for decimals
func AdjustIntForPrecision(i *big.Int, precision int) (*big.Float) {
	//Convert int to float
	f1 := new(big.Float).SetInt(i)

	//For debug purposes take this out later
	fmt.Println(f1)

	if (precision == 0) {
		return f1
	}

	//Convert precision into exponentiated int
	i2 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	//Convert precision int to float
	f2 := new(big.Float).SetInt(i2)

	//Divide unadjusted number by precision to get true val
	f1.Quo(f1, f2)

	//For debug purposes take this out later
	fmt.Printf("Adjusted Value = %s\n", f1.Text('f', 8))

	return f1
}

func AdjustFloatForPrecision(f *big.Float, precision int) (*big.Float) {
	adjustedVal := new(big.Float)
	if (precision == 0) {
		return f
	}

	//Convert precision into exponentiated int
	i2 := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	//Convert precision int to float
	f2 := new(big.Float).SetInt(i2)

	//Divide unadjusted number by precision to get true val
	adjustedVal.Quo(f, f2)

	//For debug purposes take this out later
	fmt.Printf("Adjusted Value = %s\n", adjustedVal.Text('f', 8))

	return adjustedVal
}

//Converts time in hours into approximate number of blocks
func Hours2Block(time int) (int) {
	seconds := time * 60 * 60
	blocks := seconds / 14
	return blocks
}