package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	shh "github.com/ethereum/go-ethereum/whisper/shhclient"
	whisper "github.com/ethereum/go-ethereum/whisper/whisperv6"
)

// Wconfig it will retrun context, client configuration, and id
func Wconfig(url, privatekey string) (ctx context.Context, c *shh.Client, id string) {
	ctx = context.TODO()

	c, err := shh.Dial(url)
	if err != nil {
		log.Println(err)
	}

	// if private key set
	if privatekey == "" {
		id, _ = c.NewKeyPair(ctx)
	} else {
		prKey, _ := hexutil.Decode(privatekey)
		id, err = c.AddPrivateKey(ctx, prKey)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return
}

// Post ...
func Post(ctx context.Context, c *shh.Client, username string, msg string, toPubKey string) {
	_, err := c.Post(ctx, whisper.NewMessage{
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
		Padding:   []byte(username),
		Payload:   []byte(msg),
		PublicKey: hexutil.MustDecode(toPubKey)})
	if err != nil {
		log.Println(err)
	}
}

// Receive ...
func Receive(ctx context.Context, c *shh.Client, id string) (filterID string) {
	filterID, err := c.NewMessageFilter(ctx, whisper.Criteria{PrivateKeyID: id})
	if err != nil {
		log.Println(err)
	}
	return
}
