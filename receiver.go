package main

import (
    "log"
    "net"

    "gopkg.in/mgo.v2/bson"
)

const CONN_HOST = "localhost"
const CONN_PORT = "1337"
const CONN_TYPE = "tcp"

type Msg struct {
    Message string
}

func main() {
    ln, err := net.Listen(CONN_TYPE, CONN_HOST +":"+ CONN_PORT)
    if err != nil {
        log.Printf("shit - %s\n", err)
        return
    }
    log.Println("listning on :1337")
    conn, err := ln.Accept()
    if err != nil {
        log.Println("wtf")
        return
    }
    for {
        buf := make([]byte, 1024)
        reqLen, err := conn.Read(buf)
        if err != nil {
            continue
        }
        log.Println("received message of length ", reqLen)
        var msg_textplain Msg
        err2 := bson.Unmarshal(buf[:reqLen], &msg_textplain)
        if err2 != nil {
            log.Println("errored on decoding, ", err2)
        }
        log.Printf("Msg received: %v\n", msg_textplain)
        log.Printf("encoded: %v\n", buf[:reqLen])
    }
}
