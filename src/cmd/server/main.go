package main

import (
	"context"
	"log"
	"net"

	pb "github.com/stretchkennedy/reverse-index/gen"
	"github.com/stretchkennedy/reverse-index/src/index"
	"google.golang.org/grpc"
)

const (
	port = ":20000"
)

type server struct {
	index *index.Index
}

func newServer() *server {
	return &server{
		index: index.NewIndex(),
	}
}

func (s *server) AddDocument(
	ctx context.Context, req *pb.AddDocumentRequest,
) (*pb.AddDocumentResponse, error) {
	log.Printf("AddDocument - %v", req)
	s.index.AddDocument(req.Name, req.Fields)
	return &pb.AddDocumentResponse{}, nil
}

func (s *server) QueryDocuments(
	ctx context.Context, req *pb.QueryDocumentsRequest,
) (*pb.QueryDocumentsResponse, error) {
	log.Printf("QueryDocuments - %v", req)
	locs := s.index.Query(req.Phrase)
	response := pb.QueryDocumentsResponse{
		Documents: map[string]*pb.QueryDocumentsResponse_Document{},
	}
	for _, loc := range locs {
		documentName := s.index.DocumentName(loc)
		fieldName := s.index.FieldName(loc)
		if response.Documents[documentName] == nil {
			response.Documents[documentName] = &pb.QueryDocumentsResponse_Document{
				Fields: map[string]*pb.QueryDocumentsResponse_Field{},
			}
		}
		document := response.Documents[documentName]
		if document.Fields[fieldName] == nil {
			document.Fields[fieldName] =
				&pb.QueryDocumentsResponse_Field{}
		}
		field := document.Fields[fieldName]
		field.Offsets = append(field.Offsets, loc.Offset())
	}
	return &response, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterIndexServer(s, newServer())
	log.Printf("listening on %s", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
