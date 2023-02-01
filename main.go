package main

import (
	"fmt"

	pb "github.com/torkelrogstad/proto-enum-example/gen/go/example/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	knownMarshaled, err := proto.Marshal(&pb.Foo{
		MyEnum: pb.MyEnum(1),
	})
	if err != nil {
		panic(err)
	}

	unknownMarshaled, err := proto.Marshal(&pb.Foo{
		MyEnum: pb.MyEnum(100),
	})
	if err != nil {
		panic(err)
	}

	var foo pb.Foo
	if err := proto.Unmarshal(knownMarshaled, &foo); err != nil {
		panic(err)
	}

	fmt.Printf("known: %s\n", foo.String())

	if err := proto.Unmarshal(unknownMarshaled, &foo); err != nil {
		panic(err)
	}

	fmt.Printf("unknown: %s\n", foo.String())
}
