package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

//工具文件

func UintToByte(num uint64) []byte {
	//使用binary.Write来进行编码
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func Error(str string,err error)  {
	if err!=nil {
		fmt.Println(str,err)
		os.Exit(1)
	}
}