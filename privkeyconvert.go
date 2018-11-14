package main

import "fmt"
import "crypto/sha256"
import "github.com/btcsuite/btcutil/base58"
import "encoding/hex"



var inputtype string
var input string
func main() {
	fmt.Scanln(&inputtype, &input)

	if inputtype == "hex"{
		input = "ef"+input       
		data, _ := hex.DecodeString(input)
		temp := sha256.Sum256(data)
		tempnew := sha256.Sum256(temp[:])
		str := hex.EncodeToString(tempnew[:])
		str = hex.EncodeToString(data[:])+str[:8]
		final,_:=hex.DecodeString(str)
		fmt.Printf(base58.Encode(final))
		fmt.Printf("\n")
	}


	if inputtype == "wif" {
		str := hex.EncodeToString(base58.Decode(input))
		str = substring(str, 0, len(str)-8)
		str = substring(str, 2, len(str))
		fmt.Printf("\n",str)
	}
}


func substring(source string, start int, end int) string {
    var r = []rune(source)
    length := len(r)

    if start < 0 || end > length || start > end {
        return ""
    }

    if start == 0 && end == length {
        return source
    }

    return string(r[start : end])
}
