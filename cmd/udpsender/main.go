package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
    "log"
)

func main() {
    addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
    if err != nil {
        fmt.Errorf("Could not resolve address: %e", err)
        return
    }

    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        fmt.Errorf("could not make udp connection: %v", conn)
        return
    }
    defer func() {
        fmt.Println("Clossing connection ...")
        conn.Close()
    }()

    // Read from stdinput
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        conn.Write([]byte(line))
    }
}
