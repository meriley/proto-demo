package main

import (
	"fmt"

	"./myproto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func main() {

	/**
	 * A defined Person defined by our protobuf
	 */
	person := &myproto.Person{
		Prefix: myproto.Person_MR,
		First:  &wrappers.StringValue{Value: "John"},
		Last:   &wrappers.StringValue{Value: "Smith"},
		Suffix: myproto.Person_SR,
		Location: &myproto.Location{
			Address1: &wrappers.StringValue{Value: "837 S. Ash St."},
			City:     &wrappers.StringValue{Value: "Scottsdale"},
			State:    &wrappers.StringValue{Value: "Arizona"},
		},
		Contact: &myproto.Contact{
			Phone: []*myproto.Contact_PhoneNumber{
				&myproto.Contact_PhoneNumber{
					Type:   myproto.Contact_CELL,
					Number: &wrappers.StringValue{Value: "(480) 123-4567"},
				},
			},
			Email: []*wrappers.StringValue{
				&wrappers.StringValue{Value: "jsmith@gmail.com"},
			},
		},
	}

	/**
	 * Our protobuf person written to console as a string.
	 */
	fmt.Println(person.String())

	/**
	 *	Serialize our Person to a stream of bytes
	 */

	serialized, marshalErr := proto.Marshal(person)

	if marshalErr != nil {
		fmt.Println(marshalErr)
		return
	}

	fmt.Print("This is out data serialized and ready to send: ")
	fmt.Println(serialized)

	/**
	 * Deserialize our person byte stream back into a person.
	 */
	newPerson := myproto.Person{}
	unmarshalErr := proto.Unmarshal(serialized, &newPerson)

	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return
	}

	fmt.Print("Our person has returned to person form!")
	fmt.Print(newPerson.String())

}
