package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

//创建区块链，使用bolt
type BlockChain struct {
	db   *bolt.DB //句柄
	tail []byte   //最后一个区块的哈希值
}

const BlockChainName = "blockChain.db"
const BlockBucketName = "blockBucket"
const LastHashKey = "lastHashKey"

//实现创建区块链的方法
func NewBlockChain() *BlockChain {
	//在创建的时候添加一个区块：创世块
	db, err := bolt.Open(BlockChainName, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	//defer db.Close()
	var tail []byte
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b == nil {
			fmt.Println("Bucket不存在，准备创建！", b)
			b, err = tx.CreateBucket([]byte(BlockBucketName))
			Error("tx.CreateBucket: ", err)
			//Bucket准备完毕，开始添加创世块
			genesisBlock := NewBlock(GenesisInfo, []byte{})
			b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			b.Put([]byte(LastHashKey), genesisBlock.Hash)

			tail = genesisBlock.Hash
		} else {
			tail = b.Get([]byte(LastHashKey))
		}

		return nil
	})
	return &BlockChain{db, tail}
}

//添加区块
func (bc *BlockChain) AddBlock(data string) {
	bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b == nil {
			log.Panic("Bucket不存在，请检查！")
		}
		//创建一个区块
		Block := NewBlock(data, bc.tail)
		b.Put(Block.Hash, Block.Serialize())
		b.Put([]byte(LastHashKey), Block.Hash)

		bc.tail = Block.Hash

		return nil
	})
}

//定义一个区块链的迭代器，包含db,current
type BlockChainIterator struct {
	db      *bolt.DB
	current []byte //当前所指向区块的哈希值
}

//创建迭代器，使用bc初始化
func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{db: bc.db, current: bc.tail}
}

//迭代器Next方法
func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b == nil {
			log.Panic("Bucket不存在，请检查！")
		}
		Deserialize(&block, b.Get(it.current))

		it.current = block.PrevBlockHash

		return nil
	})
	return &block
}
