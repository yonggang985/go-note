//完成虚拟机10.x.x.x和物理机１72.x.x.x按照set对应匹配，方便运维操作。
//使用map进行操作，时间复杂度为o(n),之前用脚本grep操作轮训操作，时间复杂度为o(n^2)
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f1, err := os.Open("/home/zyg/golang/own/src/github.com/yonggang985/go-test/ugc.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f1.Close()
	m1 := make(map[string]string)
	br1 := bufio.NewReader(f1)

	for {
		a, _, c := br1.ReadLine()
		astr := string(a)
		if c == io.EOF {
			break
		}
		arr1 := strings.Fields(astr)
		var key string
		for k, v := range arr1 {
			//fmt.Println(k,v)
			if k == 0 {
				key = v
			}
			if k == 1 {
				m1[key] = v
			}
		}
	}

	f2, err := os.Open("/home/zyg/golang/own/src/github.com/yonggang985/go-test/set20")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f2.Close()
	m2 := make(map[string]string)
	str2 := make([]string, 0)
	br2 := bufio.NewReader(f2)
	for {
		a, _, c := br2.ReadLine()
		if c == io.EOF {
			break
		}
		str2 = append(str2, string(a))
	}
	for _, v := range str2 {
		val, ok := m1[v]
		if ok == true {
			m2[v] = val
		}
	}

	fileName := "/home/zyg/golang/own/src/github.com/yonggang985/go-test/set20_new"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	for k, v := range m2 {
		dstFile.WriteString(k + " " + v + "\n")
	}
}
