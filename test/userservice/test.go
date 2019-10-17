package main

import "fmt"

type Main map[string]interface{}
type Edit map[string]interface{}

func main() {
	m := make(Main)
	e := make(Edit)
	m["mchtCd"] = "111"
	e["mchtCd"] = "112"

	for k, _ := range m {
		str := m[k].(string) + "," + e[k].(string) + "," + fmt.Sprint(m[k].(string) == e[k].(string))
		fmt.Println(str)
	}
}
