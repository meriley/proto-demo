package main

import (
	"fmt"

	"./myproto"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func main() {

	person := myproto.Person{
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

	fmt.Println(person.String())
}
