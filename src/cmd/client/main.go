package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"fmt"

	pb "github.com/stretchkennedy/reverse-index/gen"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

const (
	address = "localhost:20000"
)

func main() {
	// init
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewIndexClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a file to the index",
			Action: func(c *cli.Context) error {
				filename := c.Args().Get(0)
				content, err := ioutil.ReadFile(filename)
				if err != nil {
					log.Fatalf("could not read file: %v", err)
				}
				_, err = client.AddDocument(ctx, &pb.AddDocumentRequest{
					Name: filename,
					Fields: map[string]string{
						"content": string(content),
						"filename": filename,
					},
				})
				if err != nil {
					log.Fatalf("could not add: %v", err)
				}
				log.Printf("added %s", filename)
				return nil
			},
		},
		{
			Name:  "query",
			Usage: "query the index",
			Action: func(c *cli.Context) error {
				r, err := client.QueryDocuments(ctx, &pb.QueryDocumentsRequest{
					Phrase: strings.Join(c.Args(), " "),
				})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				for docName, doc := range r.Documents {
					fmt.Printf("%s\n", docName)
					for fieldName, field := range doc.Fields {
						fmt.Printf("  .%s: %v\n", fieldName, field.Offsets)
					}
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
