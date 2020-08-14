package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "hello"
	str2 := "hello 丹尼尔"

	//1. len按字节返回
	fmt.Println("length of str1:", len(str1)) //一个字母占1字节
	fmt.Println("length of str2:", len(str2)) //一个汉字占3个字节

	//2. string 遍历 如果有中文  则需要转换后遍历 r:= []rune(str)
	r := []rune(str2)
	for _, val := range r {
		fmt.Printf("%c\n", val)
	}
	//3. 字符串转整数
	n1, err1 := strconv.Atoi("12")
	fmt.Println(n1, err1) // n1 = 12, err1 : <nil>
	n2, err2 := strconv.Atoi(str1)
	fmt.Println(n2, err2) // n2 = 0 err2: strconv.Atoi: parsing "hello": invalid syntax

	//4. 整数转字符串
	str := strconv.Itoa(123)
	fmt.Printf("str3=%v, type=%T\n", str,str)

	//5. 字符串转byte []byte
	bytes := []byte("hello go")
	fmt.Printf("bytes=%c\n",bytes)

	//6 byte转字符串
	str = string([]byte{97,98,99})
	fmt.Printf("str=%v\n", str)

	//7. 十进制转2、8、16 进制
	str = strconv.FormatInt(128, 2)
	fmt.Printf("str = %v\n", str)
	str = strconv.FormatInt(128, 8)
	fmt.Printf("str = %v\n", str)
	str = strconv.FormatInt(128, 16)
	fmt.Printf("str = %v\n", str)

	//8. 查找子串是否在指定的字符串中
	b := strings.Contains("hello Daniel", "daniel")
	fmt.Printf("是否包含: %t\n", b)

	//9. 统计字符串中有几个指定的字串

	c := strings.Count("hello Daniel", "l")
	fmt.Printf("子串l的个数：%d\n", c)

	//10. 不区分大小写判断字符串是否相等
	b = strings.EqualFold("Daniel", "daniel")
	fmt.Printf("不区分大小写是否相等：%t\n", b)

	//11. 返回一个字串在指定字符串第一次出现的位置，没有则返回-1 类似indexof
	// 用来找下标 或者 看字串是否在字符串中
	index1 := strings.Index("Hello Daniel", "el")
	index2 := strings.Index("Hello Daniel", "z")
	fmt.Printf("index1= %v\n", index1)
	fmt.Printf("index2= %v\n", index2)

	//12. 返回字串在字符串最后出现的位置
	index1 = strings.LastIndex("Hello Daniel", "el")
	index2 = strings.LastIndex("Hello Daniel", "z")
	fmt.Printf("index1= %v\n", index1)
	fmt.Printf("index2= %v\n", index2)

	//13. 替换字串
	//(source string, old-substr, new-substr, n) n表示替换次数 ，-1 为全部替换
	//不会改变原串
	str1 = "Hello,Daniel"
	str2 = strings.Replace(str1, "l", "A", -1)
	fmt.Printf("str1=%v, str2=%v\n", str1, str2)

	//14. 分割字符串
	strArr := strings.Split(str1, ",")  // 返回一个数组
	fmt.Printf("str1=%v\n", strArr)

	//15. 大小写转换
	str1 = strings.ToUpper(str1)
	str2 = strings.ToLower(str1)
	fmt.Printf("str1=%v, str2=%v\n",str1, str2)
	fmt.Println("------------------------------")

	//16. 去除指定内容
	// Trim 为去除左右边界字符串，
	str = "    hello     world !        "
	str1 = strings.TrimSpace(str) // 去除两边空格
	str2 = strings.Trim(str, "hd !") //匹配到了右边界的"d !”, 以及左边界的“ h 和 空格”
	str3 := strings.TrimLeft(str, " e")//只能匹配到左边界的 空格
	str4 := strings.TrimRight(str, " !")//匹配到了右边界的空格和！

	fmt.Printf("str1 =%v\n", str1 + "Daniel")
	fmt.Printf("str2 =%v\nstr3 =%v\nstr4 =%v\n",str2, str3, str4 + "Daniel")

	//17. 判断字符串是否以什么开头什么结尾
	str1 = "ftp://192.168.1.1"
	str2 = "/home/daniel/pictures/abc.jpg"

	b1 := strings.HasPrefix(str1, "ftp")
	b2 := strings.HasSuffix(str2, ".txt")

	fmt.Printf("%v是否以ftp开头: %t\n",str1, b1)
	fmt.Printf("%v是否以.txt结尾: %t\n", str2, b2)


}
