package main

import (
    "path/filepath"
    "bufio"
    "os"
    "log"
    "strings"
    "fmt"
)

func main() {
    fps, err := filepath.Glob("*.py")
    if err != nil {
        log.Fatal("filepath.Glob", err)
    }
    w, err := os.Create("q100.txt")
    if err != nil {
        log.Fatal("OpenFile", err)
    }
    defer w.Close()
    for _, fp := range fps {
        r, err := os.OpenFile(fp, os.O_RDONLY, os.ModePerm)
        if err != nil {
            log.Fatal("OpenFile", err)
        }        
        defer r.Close()
        rb := bufio.NewReader(r)
        for i:=0; i<2; {
            bytes, err := rb.ReadBytes('\n')
            if err != nil {
                break
            }
            if strings.HasPrefix(string(bytes), "2.") {
                break
            }
            if strings.HasPrefix(string(bytes), "'''") {
                i += 1
                continue
            }
            fmt.Printf("bytes %s\n", bytes)
            w.Write(bytes)
        }
        w.WriteString("\r\n")
    }       
}
