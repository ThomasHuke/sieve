//Package sieve的主要目的是汇总常见的安全筛查机制，例如说用户的用户名安全输入，邮箱的安全输入，密码的安全输入等。
package sieve

import (
	"fmt"
)

//Sieve 主文件。使用这个筛查器可以很好的检查要检查的东西。
//type的命名使用int格式。在doc中会有详细的介绍。
func Control(ty Sie, tyN int, src string) (string, error) {
	src, err := ty.DealWith(tyN, src)
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	return src, nil
}

type number0 {
    src string
}
