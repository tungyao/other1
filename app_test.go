package main

import (
	"log"
	"testing"
)

func Test_compress(t *testing.T) {
	var s = "1kakskkkaskl12klds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kak1kakskkkaskl12klds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskdskds1kakskkkaskl12klds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskdskds1kakskkkaskl12klds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskdskds1kakskkkaskl12klds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskdskds1kakskkkaskl12klds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskdskdsskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskds1kakskkkaskl12kldskdskds"
	// var nk = 0
	log.Println(archCeil(int64(len(s)), 0x400))
	// for i := 0; i < len(s); i++ {
	// 	// udp, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IP("127.0.0.1"), Port: 8080})
	// 	// if err != nil {
	// 	// 	log.Fatalln(err)
	// 	// }
	// 	var data = make([]byte, 0x406)
	// 	if i != 0 && (i%0x400 == 0 || len(s)-nk*0x400 < 0x400) {
	// 		data = compress(len(s)/0x400, []byte(s[nk*0x400:i]))
	// 		log.Println(data)
	// 		nk += 1
	// 		// udp.WriteToUDP(data, &net.UDPAddr{IP: net.IP("127.0.0.1"), Port: 8080})
	// 	}
	// }
}
