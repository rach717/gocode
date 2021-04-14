package main

import (
	"fmt"
	"strings"
)




func  CheckPalindrome(ini string) bool {

	stringValue := strings.ReplaceAll(ini, " ", "")
	stringValue= strings.ToUpper(stringValue)
	fmt.Println("inputstring in function:", stringValue)
	result:=""
	for _, v := range stringValue{
		result = string(v) + result
		fmt.Println(result)
	}
	fmt.Println("result"+result)
	return stringValue== result
}
func main() {

	// fmt.Println("enter the string")
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("please, type some text:")
	// //user input
	// stringValue, _ := reader.ReadString('\n')
	// hc := hex.EncodeToString([]byte(stringValue))
	// fmt.Println(stringValue + "===>" + hc)
	// stringValue = strings.TrimSuffix(stringValue, "\r\n")
	// hc = hex.EncodeToString([]byte(stringValue))
	// fmt.Println(stringValue + "===>" + hc)
	// stringValue = strings.TrimPrefix(stringValue, "\r\n")
	// hc = hex.EncodeToString([]byte(stringValue))
	// fmt.Println(stringValue + "===>" + hc)
	// originalValue := stringValue
	// //fmt.Printf("original value %v\n", originalValue)
	// //remove spaces
	// stringValue = strings.ReplaceAll(originalValue, " ", "")
	//stringValue:="abc"
	//a := strings.ToUpper(stringValue)
	//s := &reversestruct{a}
	//hc := hex.EncodeToString([]byte(s.inputstring))
	//fmt.Println(s.inputstring + "===>" + hc)
	// fmt.Println("------------------")
	// hc = hex.EncodeToString([]byte("a"))
	// fmt.Println("===>" + hc)
	//fmt.Println("before calling:", s.inputstring)
	fmt.Println("hello")

}