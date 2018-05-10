package main

import (
	"encoding/binary"
	"flag"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	flag.Parse()
	in := flag.Arg(0)

	if bson.IsObjectIdHex(in) {
		o := bson.ObjectIdHex(in)
		fmt.Printf("Inspecting ObjectID \"%s\"\n", in)
		fmt.Printf("\tTime: %s\n", o.Time())
		fmt.Printf("\tPid: %d\n", o.Pid())
		m, _ := binary.Uvarint(o.Machine())
		fmt.Printf("\tMachine: %d\n", m)
		fmt.Printf("\tCounter: %d\n", o.Counter())
	} else {
		fmt.Printf("\"%s\" is not a valid BSON ObjectID\n", in)
	}
	fmt.Println("done")
}
