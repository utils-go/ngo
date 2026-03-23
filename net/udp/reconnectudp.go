package udp

import (
	"fmt"
	"github.com/utils-go/ngo/collections/concurrent"
	"net"
	"sync"
	"time"
)

/*
长连接：将连接存储起来，不主动关闭
*/

// ReconnectUdp UDP长连接实现，支持自动重连
type ReconnectUdp struct {
	DstAddr       *net.UDPAddr                  // 目标地址
	reconnectChan chan struct{}                 // 重连信号通道
	closeChan     chan struct{}                 // 关闭信号通道
	RecvBuffer    *concurrent.ConcurrentListT[[]byte] // 接收缓冲区
	conn          *net.UDPConn                  // UDP连接
	mu            sync.Mutex                    // 互斥锁，保证并发安全
}

// SendMsg 发送消息到目标地址
// 参数:
//   data: 要发送的数据
// 返回值:
//   error: 发送过程中的错误
func (u *ReconnectUdp) SendMsg(data []byte) error {
	if u.conn == nil {
		u.reconnect()
		time.Sleep(time.Millisecond * 100)
		return fmt.Errorf("%s not connect yet", u.DstAddr)
	}

	u.mu.Lock()
	_, err := u.conn.WriteToUDP(data, u.DstAddr)
	u.mu.Unlock()

	if err != nil {
		u.reconnect()
		time.Sleep(time.Millisecond * 100)
		return err
	}
	return nil
}

// NewRUdpConnection 创建一个新的UDP长连接
// 参数:
//   serverAddr: 目标地址，格式为"host:port"
// 返回值:
//   *ReconnectUdp: 新创建的UDP长连接
//   error: 创建过程中的错误
func NewRUdpConnection(serverAddr string) (*ReconnectUdp, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		return nil, err
	}
	u := &ReconnectUdp{
		DstAddr:       udpAddr,
		reconnectChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		RecvBuffer:    concurrent.NewListT[[]byte](),
	}
	go u.handleReconnect()
	go u.handRead()
	return u, nil
}

// connect 建立UDP连接
func (u *ReconnectUdp) connect() {
	if u.conn != nil {
		u.conn.Close()
	}

	con, err := net.ListenUDP("udp", nil)
	if err != nil {
		return
	}

	fmt.Printf("listen to [%s] success\n", con.LocalAddr())
	u.conn = con
}

// handleReconnect 处理重连逻辑
func (u *ReconnectUdp) handleReconnect() {
	for {
		select {
		case <-u.reconnectChan:
			u.connect()
		case <-u.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit reconnect goroutine\n", u.DstAddr)
			return
		}
	}
}

// reconnect 发送重连信号（非阻塞）
func (t *ReconnectUdp) reconnect() {
	select {
	case t.reconnectChan <- struct{}{}:
	default:
	}
}

// handRead 处理读取数据
func (u *ReconnectUdp) handRead() {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-u.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit reconnect goroutine\n", u.DstAddr)
			return
		default:
		}
		if u.conn == nil {
			u.reconnect()
			time.Sleep(time.Millisecond * 200)
			continue
		}

		n, addr, err := u.conn.ReadFromUDP(buffer)
		if err == nil {
			if addr.String() == u.DstAddr.String() {
				u.RecvBuffer.Add(buffer[:n])
			} else {
				fmt.Printf("recv from [%s] but listen on [%s]\n", addr.String(), u.DstAddr.String())
			}
		} else {
			u.reconnect()
			time.Sleep(time.Millisecond * 200)
		}
	}
}

// Close 关闭UDP连接
func (u *ReconnectUdp) Close() {
	close(u.closeChan)
}
