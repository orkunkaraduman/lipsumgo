package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var sentenceMu sync.Mutex
var sentenceBuf = bytes.NewBufferString("")
var sentenceIdx int

func getSentence() string {
	sentenceMu.Lock()
	defer sentenceMu.Unlock()
	var s string
	for {
		var err error
		s, err = sentenceBuf.ReadString('.')
		if err != nil {
			if err == io.EOF {
				sentenceBuf = bytes.NewBufferString(lipsum)
				sentenceIdx = 0
				continue
			}
			panic(err)
		}
		sentenceIdx++
		break
	}
	return fmt.Sprintf("%s #%d", strings.TrimSpace(s), sentenceIdx)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	snt := getSentence()
	h := w.Header()
	h["Content-Type"] = []string{"text/plain"}
	h["Cache-Control"] = []string{"no-cache, no-store"}
	var s string
	s += fmt.Sprintf("RequestURI: %s\n", r.RequestURI)
	s += fmt.Sprintf("RemoteAddr: %s\n", r.RemoteAddr)
	s += fmt.Sprintf("\n%s\n", snt)
	n, err := w.Write([]byte(s))
	s = ""
	if err != nil {
		s = err.Error()
	}
	log.Printf("%s [HTTP]\n%s %s %d %s", snt, r.RequestURI, r.RemoteAddr, n, s)
}

func main() {
	var interval uint
	flag.UintVar(&interval, "interval", 60, "interval in seconds")
	flag.UintVar(&interval, "n", 60, "interval in seconds")
	var addr string
	flag.StringVar(&addr, "addr", ":12345", "bind address")
	flag.StringVar(&addr, "a", ":12345", "bind address")
	flag.Parse()

	log.SetOutput(os.Stdout)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	server, listenErrChan := httpsrvInit(addr, nil, httpHandler)

	timer := time.NewTimer(0)

mainloop:
	for {
		select {
		case <-sigChan:
			break mainloop
		case listenErr := <-listenErrChan:
			if listenErr != nil && listenErr == http.ErrServerClosed {
				break mainloop
			}
			panic(listenErr)
		case <-timer.C:
			if interval == 0 {
				continue mainloop
			}
			timer.Reset(time.Duration(interval) * time.Second)
			log.Print(getSentence())
		}
	}

	if err := httpsrvClose(server, 30*time.Second); err != nil {
		log.Print(err)
	}
}
