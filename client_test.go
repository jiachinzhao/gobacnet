package gobacnet

import (
	"fmt"
	"github.com/jiachinzhao/gobacnet/types"
	"testing"
)

func TestGetLocalAddress(t *testing.T){
	getLocalAddress()
	t.Log("yes")
}
func TestNewClient2(t *testing.T) {
	client, err := NewClient2(DefaultPort)
	if err != nil{
		t.Fatalf("%+v\n", err)
	}
	resp, err := client.WhoIs(1, types.MaxInstance)
	if err != nil{
		t.Fatalf("%+v\n", err)
	}
	fmt.Printf("resp: %v\n", resp)
}
