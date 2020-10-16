package main

import "fmt"

//使用map[string]map[string]string 的map类型
//key: 表示用户名，时唯一的，不可重复
//如果某个用户名存在 就将其密码修改为“888888” 如果不存在 就增加这个用户信息 包括：
// nickname, pwd
//编写一个函数modifyUser(users map[string]map[string]string, name string)完成上述功能
func modifyUser(users map[string]map[string]string, name string, nickname string) {
	//判断users中是否有 name
	v, ok := users[name]
	if ok {
		v["pwd"] = "888888"
		v["nickname"] = nickname
	} else {

		users[name] = make(map[string]string, 2)
		users[name] ["pwd"] = "888888"
		users[name] ["nickname"] = nickname
	}


}
func main() {
	users := make(map[string]map[string]string, 10)
	users["Daniel"] = make(map[string]string, 2)
	users["Daniel"]["nickname"] = "Danielliu000"
	users["Daniel"]["pwd"] = "000"

	modifyUser(users, "Daniel", "luck123")
	modifyUser(users, "Leo", "leo222")
	fmt.Println(users)

}
