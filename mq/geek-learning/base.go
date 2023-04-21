package main

import "sync"

var zRecCount = uint32(0)  // 张大爷听到了多少句话
var lRecCount = uint32(0)  // 李大爷听到了多少句话
var total = uint32(100000) // 总共需要遇见多少次

var z0 = "吃了没，您吶?"
var z3 = "嗨！吃饱了溜溜弯儿。"
var z5 = "回头去给老太太请安！"
var l1 = "刚吃。"
var l2 = "您这，嘛去？"
var l4 = "有空家里坐坐啊。"

var liWriteLock sync.Mutex    // 李大爷的写锁
var zhangWriteLock sync.Mutex // 张大爷的写锁

type RequestResponse struct {
	Serial  uint32 // 序号
	Payload string // 内容
}
