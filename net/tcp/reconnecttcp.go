package tcp

import (
	"fmt"
	"github.com/utils-go/ngo/collections/concurrent"
	"net"
	"sync"
	"time"
)

// ReconnectTcp TCP长连接实现，支持自动重连
type ReconnectTcp struct {
	DstAddr       string                         // 目标地址
	reconnectChan chan struct{}                  // 重连信号通道
	closeChan     chan struct{}                  // 关闭信号通道
	RecvBuffer    *concurrent.ConcurrentListT[byte] // 接收缓冲区
	conn          net.Conn                       // TCP连接
	mu            sync.Mutex                     // 互斥锁，保证并发安全
}

// SendMsg 发送消息到目标地址
// 参数:
//   data: 要发送的数据
// 返回值:
//   error: 发送过程中的错误
func (t *ReconnectTcp) SendMsg(data []byte) error {
	t.mu.Lock()
	conn := t.conn
	t.mu.Unlock()

	if conn == nil {
		t.reconnect()
		return fmt.Errorf("%s not connect yet", t.DstAddr)
	}

	t.mu.Lock()
	_, err := conn.Write(data)
	t.mu.Unlock()

	if err != nil {
		t.reconnect()
		time.Sleep(time.Millisecond * 100)
		return err
	}
	return nil
}

// NewRTcpConnection 创建一个新的TCP长连接
// 参数:
//   addr: 目标地址，格式为"host:port"
// 返回值:
//   *ReconnectTcp: 新创建的TCP长连接
//   error: 创建过程中的错误
func NewRTcpConnection(addr string) (*ReconnectTcp, error) {
	if _, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		return nil, err
	}
	t := &ReconnectTcp{
		DstAddr:       addr,
		reconnectChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		RecvBuffer:    concurrent.NewListT[byte](),
	}
	go t.handleReconnect()
	go t.handleRead()
	return t, nil
}

// connect 建立TCP连接
func (t *ReconnectTcp) connect() {
	t.mu.Lock()
	oldConn := t.conn
	t.mu.Unlock()
	
	if oldConn != nil {
		oldConn.Close()
	}
	
	con, err := net.DialTimeout("tcp", t.DstAddr, time.Second*30)
	if err != nil {
		t.mu.Lock()
		t.conn = nil
		t.mu.Unlock()
		return
	}
	
	t.mu.Lock()
	t.conn = con
	t.mu.Unlock()
	
	fmt.Printf("connect to %s success\n", con.RemoteAddr())
}

// handleReconnect 处理重连逻辑
func (t *ReconnectTcp) handleReconnect() {
	for {
		select {
		case <-t.reconnectChan:
			t.connect()
		case <-t.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit reconnect goroutine\n", t.DstAddr)
			return
		}
	}
}

// reconnect 发送重连信号（非阻塞）
func (t *ReconnectTcp) reconnect() {
	select {
	case t.reconnectChan <- struct{}{}:
	default:
	}
}

// handleRead 处理读取逻辑
func (t *ReconnectTcp) handleRead() {
	buf := make([]byte, 1024)
	for {
		// 处理关闭
		select {
		case <-t.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit handleRead goroutine\n", t.DstAddr)
			return
		default:
		}
		if t.conn == nil {
			t.reconnect()
			time.Sleep(time.Millisecond * 200)
			continue
		}
		n, err := t.conn.Read(buf)
		if err == nil {
			t.RecvBuffer.AddRange(buf[:n])
		} else {
			t.reconnect()
			time.Sleep(time.Millisecond * 200)
		}
	}
}

// Close 关闭TCP连接
func (t *ReconnectTcp) Close() {
	close(t.closeChan)
}
