//Package sieve rule文件就是这个包中的规则制定包。
package sieve

//Pattern 代表了这个正则表达式的具体内容
var Pattern = map[string]string{
	"number0": "^\\w+@\\w+\\.\\w+$",
	"number1": "^\\w+$",
}
