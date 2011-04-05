package main

import (
	"gntp"
	"flag"
	"fmt"
	"os"
)

const DEFAULT_NOTIFY_NAME = "default"

func main() {
	server := flag.String("s", "localhost:23053", "growl server(host:port)")
	appname := flag.String("a", "gntp-send", "application")
	noRegister := flag.Bool("nr", false, "no register")
	hashAlgorithm := flag.String("ha", "SHA256", "hash algorithm")
	encryptAlgorithm := flag.String("ea", "AES", "encrypt algorithm")
	password := flag.String("p", "", "password")
	event := flag.String("e", DEFAULT_NOTIFY_NAME, "event")
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "usage: gntp-send [options] title message [icon] [url]\n")
		flag.PrintDefaults()
		return
	}
	title := flag.Arg(0)
	message := flag.Arg(1)
	icon := flag.Arg(2)
	url := flag.Arg(3)

	client := gntp.NewClient()
	client.SetAppName(*appname)
	client.SetServer(*server)
	client.SetPassword(*password)
	client.SetHashAlgorithm(*hashAlgorithm)
	client.SetEncryptAlgorithm(*encryptAlgorithm)

	n := []gntp.Notification{{DEFAULT_NOTIFY_NAME, DEFAULT_NOTIFY_NAME, true}}
	if *event != DEFAULT_NOTIFY_NAME {
		n = append(n, gntp.Notification{*event, *event, true})
	}
	var err os.Error
	if *noRegister == false {
		err = client.Register(n)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.String())
			return
		}
	}
	err = client.Notify(*event, title, message, icon, url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.String())
		return
	}
}