package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var base_dst_dir = "indexed"

func handle_line(lineChan chan string) {
	for {
		select {
		case line := <-lineChan:

			dst := []byte(line)
			if len(dst) < 4 {
				return
			}
			dst_dir := fmt.Sprintf("%v/%v/%v/%v", base_dst_dir, dst[0], dst[1], dst[2])
			os.MkdirAll(dst_dir, 0755)

			f, err := os.OpenFile(fmt.Sprintf("%v/list", dst_dir), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			if _, err := f.WriteString(fmt.Sprintf("%s\n", dst)); err != nil {
				log.Println(err)
			}
			f.Close()
		}
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	const threads = 16
	var channels [threads]chan string
	for i := range channels {
		channels[i] = make(chan string)
		go handle_line(channels[i])
	}

	var cnt uint64
	for scanner.Scan() {
		cnt++
		channels[cnt%threads] <- scanner.Text()
	}

}
