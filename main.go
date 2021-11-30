package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"relafy/prisma/db"

	"github.com/speps/go-hashids/v2"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func hash(salt string) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 7
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434})
	//d, _ := h.DecodeWithError(e)

	return e
}

func run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	url := os.Args[1]
	short := hash(url)

	// create a bit
	createdBit, err := client.Bit.CreateOne(
		db.Bit.URL.Set(url),
		db.Bit.Short.Set(short),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdBit, "", "  ")
	fmt.Printf("created bit: %s\n", result)

	// find a single bit
	post, err := client.Bit.FindUnique(
		db.Bit.ID.Equals(createdBit.ID),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ = json.MarshalIndent(post, "", "  ")
	fmt.Printf("bit: %s\n", result)

	return nil
}
