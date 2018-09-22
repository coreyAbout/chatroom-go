package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Params struct {
	Room, User, Input string
}

type Rect struct{}

// to store all chat history
var m = make(map[string][]string)
var tM = make(map[string][]int64)

//First receive parameter. Second send back to client parameter, must be a pointer
//return value error is needed
func (r *Rect) Area(p Params, ret *string) error {
	chat := p.Room + p.User + " " + p.Input

	//all chat history
	m[p.Room] = append(m[p.Room], chat)
	//time stamp for all chat
	tM[p.Room] = append(tM[p.Room], time.Now().Unix())

	// chatlog of current room
	roomChat := ""

	// delete chat history longer than 7 days: here 300 seconds
	if len(m[p.Room]) > 1 {
		t0 := tM[p.Room][len(tM[p.Room])-2]
		t1 := time.Now().Unix()
		if t1-t0 > 300 {
			delete(m, p.Room)
			delete(tM, p.Room)
		}
	}

	// send chat history
	for _, each := range m[p.Room] {
		if each[0:4] == p.Room {
			roomChat += each
		}
	}

	*ret = roomChat
	fmt.Println(m)
	fmt.Println(tM)

	return nil
}

//for listing current chatrooms, not using
func (r *Rect) Perimeter(p string, ret *string) error {
	roomlist := ""

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	for _, eachroom := range keys {
		roomlist += eachroom + "   "
	}

	*ret = roomlist
	return nil
}

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	rect := new(Rect)
	//rpc register
	rpc.Register(rect)
	//acquire tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:6002")
	chkError(err)
	//listen
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError(err2)

	//infinate loop to receive and send
	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}

		//goroutine 单独处理rpc连接请求
		go rpc.ServeConn(conn)
	}
}

//modified from: https://www.oudahe.com/p/41463/
