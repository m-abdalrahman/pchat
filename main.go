package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func init() {
	//flags
	flag.StringVar(&Username, "username", "", "set username, example 'pchat -username \"user\"'")
	flag.StringVar(&PrivateKey, "privatekey", "", "set your private key")
}

func main() {
	flag.Parse()
	CheckArgs()

	ctx, c, id := Wconfig()
	fID := Receive(ctx, c, id)

	var pubKey, msg string
	reader := bufio.NewReader(os.Stdin)

	prKey, err := c.PrivateKey(ctx, id)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("**********************************************************\n")
	fmt.Println("Your Private Key:", hexutil.Encode(prKey))

	puKey, err := c.PublicKey(ctx, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Your Public Key:", hexutil.Encode(puKey))
	fmt.Print("**********************************************************\n\n")

	fmt.Print("Enter your partner public key: ")
	pubKey, _ = reader.ReadString('\n')

	// receive message
	receivedMsg := make(chan string)
	go func(chan string) {
		for {
			m, err := c.FilterMessages(ctx, fID)
			if err != nil {
				log.Fatalln(err)
			}

			if len(m) > 0 {
				receivedMsg <- fmt.Sprint(string(m[0].Padding)+": ", string(m[0].Payload))
			}
		}
	}(receivedMsg)

	// send message
	go func() {
		for {
			msg, _ = reader.ReadString('\n')
			msg = strings.TrimSpace(msg)

			if msg != "" {
				Post(ctx, c, Username, msg, strings.TrimSpace(pubKey))
				msg = ""
			}
		}
	}()

	for {
		select {
		case m := <-receivedMsg:
			fmt.Println(m)
		default:
		}
	}
}
