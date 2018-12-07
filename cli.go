package main

import (
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

const Usage = `
    addBlock --data DATA "add a block"
    printChain "print block Chain"
`

func (cli *CLI) Run() {
	cmds := os.Args

	if len(cmds) < 2 {
		fmt.Println(Usage)
		os.Exit(1)
	}
	switch cmds[1] {
	case "addBlock":
		fmt.Printf("添加区块命令被调用，数据：%s\n", cmds[2])
		data := cmds[2]
		cli.AddBlock(data)
	case "printChain":
		fmt.Println("打印区块命令被调用")
		cli.PrintChain()
	default:
		fmt.Println("无效命令被调用，请检查")
		fmt.Println(Usage)
	}
}
