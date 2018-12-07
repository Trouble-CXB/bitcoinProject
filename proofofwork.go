package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义一个工作量证明
type ProofOfWork struct {
	block  *Block
	target *big.Int //储存哈希值，内置了一些方法Cmp：比较方法
	//SetBytes :  把bytes转化成big.Int
	//SetString： 把String转化成big.Int
}

//难度值定义与推导
const Bits = 4

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//写难度值。难度值是推导出来的，先写成固定。
	//targetStr := "0010000000000000000000000000000000000000000000000000000000000000"
	bigIntTmp := big.NewInt(1)
	//bigIntTmp.Lsh(bigIntTmp,256)
	//bigIntTmp.Rsh(bigIntTmp,16)
	bigIntTmp.Lsh(bigIntTmp, 256-Bits*4)

	pow.target = bigIntTmp

	return &pow
}

//pow的运算函数，为了获取挖坑的随机数，同时返回区块的哈希值
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//1、获取block数据2、拼接nonce  调用函数prepareData()
	//3、sha256
	//4、与难度值比较
	//5、哈希值大于难度值，nonce++
	//6、哈希值小于难度值，返回数据
	var nonce uint64
	var hash [32]byte
	var bigIntTmp big.Int
	for {
		//fmt.Printf("%x\r", hash)
		hash = sha256.Sum256(pow.prepareData(nonce))
		bigIntTmp.SetBytes(hash[:])
		if bigIntTmp.Cmp(pow.target) == -1 {
			//此时哈希值bigIntTmp 小与 难度值pow.target
			fmt.Printf("挖矿成功！nonce：%-8d,哈希值：%x\n", nonce, hash)
			break
		} else {
			nonce++
		}
	}
	return hash[:], nonce
}

//block数据拼接nonce
func (pow *ProofOfWork) prepareData(nonce uint64) []byte {
	block := pow.block
	tmp := [][]byte{
		UintToByte(block.Version),
		block.PrevBlockHash,
		block.PrevBlockHash,
		UintToByte(block.TimeStamp),
		UintToByte(block.Difficulity),
		UintToByte(nonce),
		block.Data,
	}
	data := bytes.Join(tmp, []byte{})
	return data
}

//校验哈希值
func (pow *ProofOfWork) IsValid() bool {

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)

	var tmp big.Int
	tmp.SetBytes(hash[:])

	return tmp.Cmp(pow.target) == -1
}
