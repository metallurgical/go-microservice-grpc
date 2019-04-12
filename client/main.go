package main

import (
	"io"
	"log"

	pb "github.com/metallurgical/go-microservice-grpc/api"
	ctx "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1989", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to grpc server: %v", err)
	}
	defer conn.Close()


	client := pb.NewContactClient(conn)

	contact := &pb.ContactInformation{
		Id: 1,
		Name: "Norlihazmey Ghazali",
		MobileNo: "60123456789",
		OfficeNo: "034543456",
		Email: "norlihazmey.ghazali@gmail.com",
		Address: []*pb.ContactInformation_Address{
			&pb.ContactInformation_Address{
				Line_1: "Lot 123, Block A",
				Line_2: "Jalan 123",
				Postcode: "68100",
				City: "Batu Caves",
				Country: "Selangor",
			},
			&pb.ContactInformation_Address{
				Line_1: "Lot 345, Block B",
				Line_2: "Jalan 345",
				Postcode: "765678",
				City: "Batu Caves",
				Country: "Selangor",
			},
		},
	}

	resp, err := client.CreateNewContact(ctx.Background(), contact)

	if !resp.Success {
		log.Printf("Creating new contact was not successfull")
	} else {
		log.Printf("Successfull create new contact with ID %v", resp.GetId())
	}

	search := pb.ContactSearch{Name: "Norlihazmey Ghazali"}

	stream, err := client.GetContacts(ctx.Background(), &search)
	if err != nil {
		log.Printf("Error get customer")
	}

	for {
		contact, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error searching for existing contact")
		}

		log.Printf("Found : %v", contact)
	}
}