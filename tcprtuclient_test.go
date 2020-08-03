package modbus

import (
	"net"
	"testing"
	"time"
)

func TestTCPRTU(t *testing.T) {

	conn, err := net.Dial("tcp", "localhost:8001")
	handler := NewTCPRTUClientHandler(conn)
	handler.Timeout = 5 * time.Second
	handler.SlaveId = 3
	//handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	//// Connect manually so that multiple requests are handled in one connection session
	//err := handler.Connect()
	//if err != nil {
	//	panic(err)
	//}
	defer handler.Close()

	client := NewClient(handler)

	results, err := client.ReadHoldingRegisters(0, 100)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("results:", results)
	t.Log("len(results):", len(results))
}
