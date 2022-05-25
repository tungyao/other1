package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

// 消息体结构
// header 1bit | id 2bit | length 3bit | payload 1024 + 6
// header 0 无后续,>0 有后续 , 后续的 标号代表 数据体序号 exp: 1,2,3
// id 消息体的id
// length 当前消息体的长度
// payload 实际消息体
var globe = make(map[string][]string)

func compress(n int, str []byte) []byte {
	var l = len(str)
	var b = make([]byte, l+6)
	b[0] = uint8(n)
	b[1] = 0x0
	b[2] = 0x1
	b[3] = uint8(l & 0xf)
	b[4] = uint8((l >> 4) & 0xf)
	b[5] = uint8((l >> 8) & 0xf)
	for i := 6; i < len(b); i++ {
		b[i] = str[i-6]
	}
	return b
}

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{Port: 8080})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		dt40f := make([]byte, 0x406)
		_, src, err := udp.ReadFromUDP(dt40f)
		var length int32 = int32(dt40f[4]) & 0xf
		length += (int32(dt40f[5]) & 0xf) << 4
		length += (int32(dt40f[6]) & 0xf) << 8

		var data string
		if dt40f[0] == 0 {
			data = string(dt40f[6 : 6+length])
		} else {
			point := int((dt40f[1] & 0xf) + ((dt40f[2] & 0xf) << 4))
			key := src.String() + strconv.Itoa(point)
			// create cache
			if point == 0 {
				data = strings.Join(globe[key], "")
				delete(globe, key)
			} else {
				if globe[key] == nil {
					globe[key] = make([]string, int(dt40f[1]))
				}
				globe[key][int(dt40f[1])] = string(dt40f[6 : 6+length])
			}
		}
		if data != "" {
			log.Println(data)
		}
		if err == nil {
			udp.WriteToUDP([]byte("hello world"), src)
		}
	}

}
