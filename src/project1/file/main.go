package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("d:/nginx.conf")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//输出文件
	fmt.Printf("file=%v\n", file)

	//关闭文件
	//当函数退出时，要及时关闭文件
	defer file.Close()
	//err = file.Close()
	//if err != nil {
	//	fmt.Println("close file err=", err)
	//}
	//readBuf(file)
	//
	////一次性读取内容,适合文件不大
	//readDirect()

	//创建并写入文件
	createAndWrite()

	//判断文件是否存在
	var exists bool
	exists, err = PathExists("d:/")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("exists=", exists)

	//CopyFile("d:/abc.txt", "e:/abc.txt")

	//统计文件内字符的个数
	countMap := make(map[string]int)
	CountChar("e:/abc.txt", countMap)
	for k, v := range countMap {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}
}

//带缓冲区读取文件
func readBuf(file *os.File) {
	//带缓存的方式读取文件
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		//读到换行符就结束
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			//读到文件末尾就结束
			break
		}
		//输出内容
		fmt.Println(str)
	}
	fmt.Println("文件读取结束")
}

//一次性读取文件，适合文件不大的情况
func readDirect() {
	content, err := ioutil.ReadFile("d:/nginx.conf")
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))
}

//创建文件并写入
func createAndWrite() {
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}
	//及时关闭file
	defer file.Close()
	str := "hello,Gardon"
	//带缓冲区的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str + "\n")
	}
	//因为writer是带缓存的，只有调用Flush之后才真正的写入
	writer.Flush()
}

// PathExists 判断路径是否存在
//如果存在返回true，否则返回false
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CopyFile 复制文件
//src：源文件
//dst：目标文件
//返回：written写入字节数
//err异常信息
func CopyFile(src string, dst string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("open file err=%v\r\n", err)
		return 0, err
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\r\n", err)
		return 0, err
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

// CountChar 统计文件中字符串的个数
func CountChar(file string, countMap map[string]int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("open file error:", err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)
		arr := []rune(line)
		for _, v := range arr {
			fmt.Printf("%c\n", v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'B':
				countMap["charCount"]++
			case v >= '0' && v <= '9':
				countMap["numCount"]++
			case v == ' ':
				countMap["spaceCount"]++
			default:
				countMap["otherCount"]++
			}
		}
		if err == io.EOF {
			break
		}

	}
}
