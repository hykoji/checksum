package main

import (
	"crypto/md5"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s FILE\n", os.Args[0])
		os.Exit(-1)
	}

	for _, file := range args {
		src, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}
		st, err := src.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}
		if st.IsDir() {
			fmt.Fprintf(os.Stderr, "%s is a directory\n", file)
			continue
		}
		fmt.Println(file)

		sha1 := sha1.New()
		io.Copy(sha1, src)
		fmt.Printf("SHA1: %x\t\n", sha1.Sum(nil))

		src.Seek(0, 0)
		md5 := md5.New()
		io.Copy(md5, src)
		fmt.Printf("MD5 : %x\t\n", md5.Sum(nil))

		src.Close()
	}
}
