package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "github.com/metallurgical/go-microservice-grpc/api"
	"google.golang.org/grpc"
	"strings"
)

var _ = fmt.Printf

// Server side struct
type Server struct {
	contactList []*pb.ContactInformation
}

// Server handler to create new contact.
func (server *Server) CreateNewContact(context context.Context, contact *pb.ContactInformation) (*pb.ContactResponse, error) {
	server.contactList = append(server.contactList, contact)

	return &pb.ContactResponse{Id: contact.Id, Success: true}, nil
}

// Server handler to filter contact
func (server *Server) GetContacts(search *pb.ContactSearch, stream pb.Contact_GetContactsServer) error {

	for _, contact := range server.contactList {
		if search.Name != "" {
			if !strings.Contains(contact.Name, search.Name) {
				continue
			} else {
				// Send stream response to client if found.
				if err := stream.Send(contact); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func main() {
	conn, err := net.Listen("tcp", ":1989")
	if err != nil {
		log.Fatalf("Could not list to port 1989: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterContactServer(grpcServer, &Server{})
	grpcServer.Serve(conn)
}

