package main

import (
    "fmt"
    "log"
    "errors"
    "strings"
    "io"
    "net"
)

func main() {
    /*file, err := os.Open("messages.txt")
    if err != nil {
        log.Fatalf("Could not open file: %v", err)
    }
    defer file.Close()*/
    listener, err := net.Listen("tcp", ":42069")
    if err != nil {
        log.Fatalf("Could not create listener: %v", err)
        return
    }
    defer listener.Close()
    for {
        fmt.Println("Waiting for connection")
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Connection Accepted")
        linesCh := getLinesChannel(conn)
        for line := range linesCh {
            fmt.Printf("%s\n", line)
        }
        fmt.Println("Connection has been closed")
    }
    return
}

func getLinesChannel(f io.ReadCloser) <-chan string {
    linesCh := make(chan string)
    go func () {
        buff := make([]byte, 8)
        var currentLine string
        for {
            num, err := f.Read(buff)
            if err != nil {
                if currentLine != "" {
                    //fmt.Printf("read: %s\n", currentLine)
                    linesCh <- currentLine
                    currentLine = ""
                }
                if errors.Is(err, io.EOF) {
                    close(linesCh)
                    break
                }

                log.Fatalf("Could not read bytes: %v", err)
                return
            }
            // Split buff at new line characters
            lines := strings.Split(string(buff[:num]), "\n")
            for i := 0; i < len(lines) -1; i++ {
                //fmt.Printf("read: %s%s\n", currentLine, lines[i])
                linesCh <- currentLine + lines[i]
                currentLine = ""
            }
            currentLine += lines[len(lines)-1]
        }
    }()
    return linesCh
}
