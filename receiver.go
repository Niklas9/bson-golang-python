package main

import (
    "log"
    "net"
    "os"
    "strconv"

    "gopkg.in/mgo.v2/bson"
)


const CONN_HOST = "localhost"
const CONN_PORT = 1337
const CONN_TYPE = "unix" // or tcp, make this yml configurable
const CONN_SOCK_FILE = "/tmp/whatever.sock"

type Message struct {
    Text string
    IsSent bool
}

func handleClient(conn net.Conn) {
    buf := make([]byte, 1024)
    reqLen, err := conn.Read(buf)
    if err != nil {
        log.Printf("unable to read from connection, %s", err)
    }
    log.Println("received message of length ", reqLen)
    var msg Message
    err2 := bson.Unmarshal(buf[:reqLen], &msg)
    if err2 != nil {
        log.Println("errored on decoding, ", err2)
    }
    log.Printf("Msg received: %+v\n", msg)
    log.Printf("encoded: %+v\n", buf[:reqLen])
    if msg.IsSent {
        log.Printf("OMG OMG OMG ! It's true!\n")
    } else {
        log.Printf("WTF WTF WTF\n")
    }
}

func clientConns(ln net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    i := 0
    go func() {
        for {
            client, err := ln.Accept()
            if client == nil {
                log.Printf("couldn't accept: %v\n", err)
            }
            i++
            log.Printf("%d: %v <-> %v\n", i, client.LocalAddr(),
                       client.RemoteAddr())
            ch <- client
        }
    }()
    return ch
}

func main() {
    var ln net.Listener
    var err error
    if CONN_TYPE == "unix" {
        ln, err = net.Listen(CONN_TYPE, CONN_SOCK_FILE)
        defer os.Remove(CONN_SOCK_FILE)
    } else {
        ln, err = net.Listen(CONN_TYPE, CONN_HOST +":"+ strconv.Itoa(CONN_PORT))
    }
    if err != nil {
        log.Printf("unable to create %s socket - %s\n", CONN_TYPE, err)
        return
    }
    defer ln.Close()
    if CONN_TYPE == "unix" {
        log.Printf("listening on %s\n", CONN_SOCK_FILE)
    } else {
        log.Println("listning on :%d", CONN_PORT)
    }
    conns := clientConns(ln)
    for {
        go handleClient(<-conns)
    }
}
