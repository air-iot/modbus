package modbus

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestRTUTCP(t *testing.T) {

	handler := NewRTUTCPClientHandler("127.0.0.1:9000")
	handler.Timeout = 5 * time.Second
	handler.SlaveId = 1
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer handler.Close()

	client := NewClient(handler)

	for {
		results, err := client.ReadHoldingRegisters(100, 1)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("results:", results)
		t.Log("len(results):", len(results))
		//time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
	}

}
