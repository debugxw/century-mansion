package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

var Cbr = CommunicationBatchRead{}

type CommunicationBatchRead struct {
}

// 发送缓冲
func (c *CommunicationBatchRead) writeLoop(conn *net.TCPConn, chanel chan []byte) {
	for b := range chanel {
		_, err := conn.Write(b)
		if err != nil {
			fmt.Println("write err:", err)
			return
		}
	}
}

// 序列化RequestResponse，并发送
// 序列化后的结构如下：
//
//	长度  4字节
//	Serial 4字节
//	PayLoad 变长
func (c *CommunicationBatchRead) writeTo(r *RequestResponse, chanel chan []byte) {
	payloadBytes := []byte(r.Payload)
	length := uint32(len(payloadBytes)) + 4
	packages := make([]byte, length+4)

	binary.BigEndian.PutUint32(packages, length)
	binary.BigEndian.PutUint32(packages[4:8], r.Serial)
	copy(packages[8:], payloadBytes)

	chanel <- packages
}

// 接收数据，反序列化成RequestResponse
func (c *CommunicationBatchRead) readFrom(conn *net.TCPConn, recvBuf []byte, recvIndex int) ([]*RequestResponse, int, error) {
	retResponses := make([]*RequestResponse, 0)
	n, err := conn.Read(recvBuf[recvIndex:cap(recvBuf)])
	if err != nil {
		return nil, n + recvIndex, fmt.Errorf("读数据故障：%s", err.Error())
	}

	index := 0
	for {
		ret := &RequestResponse{}
		if index+8 <= n+recvIndex {
			length := int(binary.BigEndian.Uint32(recvBuf[index : index+4]))
			ret.Serial = binary.BigEndian.Uint32(recvBuf[index+4 : index+8])

			if index+8+length-4 <= n+recvIndex {
				ret.Payload = string(recvBuf[index+8 : index+8+length-4])
				index += length + 4
				retResponses = append(retResponses, ret)
			} else {
				break
			}
		} else {
			break
		}
	}

	if n+recvIndex-index > 0 {
		copy(recvBuf[:n+recvIndex-index], recvBuf[index:n+recvIndex])
	}

	return retResponses, n + recvIndex - index, nil
}

// 张大爷的耳朵
func (c *CommunicationBatchRead) zhangDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup, chanel chan []byte) {
	defer wg.Done()
	receiveBuf := make([]byte, 0, 1024)
	receiveIndex := 0
	for zRecCount < total*3 {
		responses, nextReceiveIndex, err := c.readFrom(conn, receiveBuf, receiveIndex)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		receiveIndex = nextReceiveIndex
		for _, r := range responses {
			if r.Payload == l2 { // 如果收到：您这，嘛去？
				go c.writeTo(&RequestResponse{r.Serial, z3}, chanel) // 回复：嗨！吃饱了溜溜弯儿。
			} else if r.Payload == l4 { // 如果收到：有空家里坐坐啊。
				go c.writeTo(&RequestResponse{r.Serial, z5}, chanel) // 回复：回头去给老太太请安！
			} else if r.Payload == l1 { // 如果收到：刚吃。
				// 不用回复
			} else {
				fmt.Println("张大爷听不懂：" + r.Payload)
				break
			}
			zRecCount++
		}
	}
}

// 张大爷的嘴
func (c *CommunicationBatchRead) zhangDaYeSay(chanel chan []byte) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		c.writeTo(&RequestResponse{nextSerial, z0}, chanel)
		nextSerial++
	}
}

// 李大爷的耳朵，实现是和张大爷类似的
func (c *CommunicationBatchRead) liDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup, chanel chan []byte) {
	defer wg.Done()
	recBuf := make([]byte, 0, 1024)
	recIndex := 0
	for lRecCount < total*3 {
		responses, nextRecIndex, err := c.readFrom(conn, recBuf, recIndex)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		recIndex = nextRecIndex
		for _, r := range responses {
			if r.Payload == z0 { // 如果收到：吃了没，您吶?
				c.writeTo(&RequestResponse{r.Serial, l1}, chanel) // 回复：刚吃。
			} else if r.Payload == z3 {
				// do nothing
			} else if r.Payload == z5 {
				// do nothing
			} else {
				fmt.Println("李大爷听不懂：" + r.Payload)
				break
			}
			lRecCount++
		}
	}
}

// 李大爷的嘴
func (c *CommunicationBatchRead) liDaYeSay(chanel chan []byte) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		c.writeTo(&RequestResponse{nextSerial, l2}, chanel)
		nextSerial++
		c.writeTo(&RequestResponse{nextSerial, l4}, chanel)
		nextSerial++
	}
}

func (c *CommunicationBatchRead) startServer(wg *sync.WaitGroup) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("张大爷在胡同口等着 ...")
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("碰见一个李大爷:" + conn.RemoteAddr().String())
		chanel := make(chan []byte, 100)
		go c.writeLoop(conn, chanel)
		go c.zhangDaYeListen(conn, wg, chanel)
		go c.zhangDaYeSay(chanel)
	}

}

func (c *CommunicationBatchRead) startClient(wg *sync.WaitGroup) *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	channel := make(chan []byte, 100)
	go c.writeLoop(conn, channel)
	go c.liDaYeListen(conn, wg, channel)
	go c.liDaYeSay(channel)
	return conn
}
