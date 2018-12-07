package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

const GenesisInfo = "开始学习区块链！"

//创建区块
type Block struct {
	Version       uint64 //区块版本号
	PrevBlockHash []byte //前区块哈希
	MerKleRoot    []byte //
	TimeStamp     uint64 //时间戳
	Difficulity   uint64 //难度值
	Nonce         uint64 //随机数
	Data          []byte //交易信息
	Hash          []byte //当前区块哈希
}

//实现创建区块的方法
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:       02,
		PrevBlockHash: prevBlockHash,
		MerKleRoot:    []byte{},
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulity:   Bits,
		//Nonce:         10,
		Data: []byte(data),
		Hash: []byte{},
	}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

//序列化,
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	Error("编码错误：", err)
	return buffer.Bytes()
}

//反序列化,
func Deserialize(block *Block,data []byte) /**Block*/ {
	//fmt.Printf("解码的数据： %x\n", data)
	//var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Error("解码错误：", err)
	//return &block
}
