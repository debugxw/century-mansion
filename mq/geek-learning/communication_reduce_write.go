package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
)

var Crw = CommunicationReduceWrite{}

type CommunicationReduceWrite struct {
}

// 发送缓冲
func (c *CommunicationReduceWrite) writeLoop(conn *net.TCPConn, chanel chan []byte) {
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
func (c *CommunicationReduceWrite) writeTo(r *RequestResponse, chanel chan []byte) {
	payloadBytes := []byte(r.Payload)
	length := uint32(len(payloadBytes)) + 4
	packages := make([]byte, length+4)

	binary.BigEndian.PutUint32(packages, length)
	binary.BigEndian.PutUint32(packages[4:8], r.Serial)
	copy(packages[8:], payloadBytes)

	chanel <- packages
}

// 接收数据，反序列化成RequestResponse
func (c *CommunicationReduceWrite) readFrom(conn *net.TCPConn) (*RequestResponse, error) {
	ret := &RequestResponse{}
	buf := make([]byte, 4)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf("读长度故障：%s", err.Error())
	}
	length := binary.BigEndian.Uint32(buf)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf("读Serial故障：%s", err.Error())
	}
	ret.Serial = binary.BigEndian.Uint32(buf)
	payloadBytes := make([]byte, length-4)
	if _, err := io.ReadFull(conn, payloadBytes); err != nil {
		return nil, fmt.Errorf("读Payload故障：%s", err.Error())
	}
	ret.Payload = string(payloadBytes)
	return ret, nil
}

// 张大爷的耳朵
func (c *CommunicationReduceWrite) zhangDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup, chanel chan []byte) {
	defer wg.Done()
	for zRecCount < total*3 {
		r, err := c.readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// fmt.Println("张大爷收到：" + r.Payload)
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

// 张大爷的嘴
func (c *CommunicationReduceWrite) zhangDaYeSay(chanel chan []byte) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		c.writeTo(&RequestResponse{nextSerial, z0}, chanel)
		nextSerial++
	}
}

// 李大爷的耳朵，实现是和张大爷类似的
func (c *CommunicationReduceWrite) liDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup, chanel chan []byte) {
	defer wg.Done()
	for lRecCount < total*3 {
		r, err := c.readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// fmt.Println("李大爷收到：" + r.Payload)
		if r.Payload == z0 { // 如果收到：吃了没，您吶?
			go c.writeTo(&RequestResponse{r.Serial, l1}, chanel) // 回复：刚吃。
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

// 李大爷的嘴
func (c *CommunicationReduceWrite) liDaYeSay(chanel chan []byte) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		c.writeTo(&RequestResponse{nextSerial, l2}, chanel)
		nextSerial++
		c.writeTo(&RequestResponse{nextSerial, l4}, chanel)
		nextSerial++
	}
}

func (c *CommunicationReduceWrite) startServer(wg *sync.WaitGroup) {
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

func (c *CommunicationReduceWrite) startClient(wg *sync.WaitGroup) *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	channel := make(chan []byte, 100)
	go c.writeLoop(conn, channel)
	go c.liDaYeListen(conn, wg, channel)
	go c.liDaYeSay(channel)
	return conn
}
