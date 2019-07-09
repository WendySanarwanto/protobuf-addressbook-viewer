package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/wendysanarwanto/protobuf-addressbook"
)

func main() {
	// Expecting a CLI Argument when running this program, to be the filename to save/load
	if len(os.Args) != 2 {
		log.Fatalf("[ERROR] Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fileName := os.Args[1]

	// Read the specified address book
	binContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("[ERROR] '%s' file does not exist. \n", fileName)
		} else {
			log.Fatalln("[ERROR] Error reading file: ", err)
		}
		return
	}

	fmt.Println("Displaying your address book ...")
	// Start unmarshall
	addressBook := &pb.AddressBook{}
	err = proto.Unmarshal(binContent, addressBook)
	if err != nil {
		log.Fatalln("[ERROR] Failed to parse address book: ", err)
		return
	}

	// Iterate People in addrssBook
	for _, address := range addressBook.People {
		fmt.Println("---------------------")
		fmt.Println("Name:\t\t", address.Name);
		fmt.Println("Email:\t\t", address.Email);
		fmt.Println("Phone(s):")
		for _, phone := range address.Phones {
			fmt.Println("\t* Number:", phone.Number, "; Type:", phone.Type)
		}
		fmt.Println("---------------------")
	}
}