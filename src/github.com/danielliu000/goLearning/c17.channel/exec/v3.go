package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//练习
//1. 开启一个协程writeDataToFile, 随机生成1000个数据，并存放到文件中
//2. 当writeDataToFile完成1000个数据到文件后， 让sort协程从文件中读取1000个数据
//   并完成排序， 重新写入到另一个文件
//3. 考察点： 协程、管道、文件综合使用
//4. 扩展功能： 开10个协程writeDataToFile，每个协程随机生成1000个数据，并存放到10个文件中
//5. 当10个文件都生成了， 让10个sort协程从10个文件中读取1000个文件，并完成排序，重新写入到10个结果文件

func writeDataToChan(dataChan chan int, num int, exitChan chan bool) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	for i := 1; i <= num; i++ {
		randNum := rand.Intn(num)
		dataChan <- randNum
	}
	//close(dataChan)
	exitChan <- true
	wg2.Done()
}
func writeDataToFile(dataChan chan int, filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for v := range dataChan {
		_, err = writer.WriteString(strconv.Itoa(v) + "\n")
		if err != nil {
			fmt.Printf("write file error: %v\n", err)
			return
		}
		err = writer.Flush()
		if err != nil {
			fmt.Printf("write file error: %v\n", err)
			return
		}
	}
	wg2.Done()
}
func sortDataToSlice(readFilePath string, num int) []int {
	data := make([]int, num, num)
	file, _ := os.Open(readFilePath)
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.ReplaceAll(str, "\n", "")
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil
		}
		data = append(data, num)
	}
	sort.Ints(data)
	return data

}
func writeSortDataToFile(srcFilePath string, dstFilePath string, num int) {
	file, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file Error: %v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	data := make([]int, num, num)
	data = sortDataToSlice(srcFilePath, num)
	for _, v := range data {
		if v == 0 {
			continue
		}
		_, err = writer.WriteString(strconv.Itoa(v) + "\n")
		if err != nil {
			fmt.Printf("write file Error: %v", err)
			return
		}
		err = writer.Flush()
		if err != nil {
			fmt.Printf("write file Error: %v", err)
			return
		}
	}

	wg2.Done()
}

var wg2 = sync.WaitGroup{}

func main() {
	goRoutineNum := 10
	num := 1000
	dataChan := make(chan int, num)
	exitChan := make(chan bool, goRoutineNum)


	//sortChan := make(chan int, 1000)

	for i := 1; i <= goRoutineNum; i++ {
		wg2.Add(1)
		go writeDataToChan(dataChan, num, exitChan)

		filePath1 := fmt.Sprintf("./randNum_%v.txt", i)

		wg2.Add(1)
		go writeDataToFile(dataChan, filePath1)

		filePath2 := fmt.Sprintf("./sortNum_%v.txt", i)

		wg2.Add(1)
		go writeSortDataToFile(filePath1, filePath2, num)
	}

	wg2.Add(1)
	go func() {
		for i := 1; i <= goRoutineNum; i++ {
			<-exitChan
		}
		close(dataChan)
		close(exitChan)
		wg2.Done()
	}()

	wg2.Wait()
}
