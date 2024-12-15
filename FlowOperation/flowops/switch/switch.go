package switchop

import (
	"fmt"
	"io"
	"log"
	"net"
)

// OpenFlowMessage interface ensures that any OpenFlow message can be serialized and sent.
type OpenFlowMessage interface {
	WriteTo(w io.Writer) (int64, error)
}

func SendToSwitch(switchID uint64, message OpenFlowMessage) error {
	log.Printf("SendToSwitch switchID=%d\n", switchID)
	switchPortBase := 6633
	switchAddress := fmt.Sprintf("127.0.0.1:%d", switchPortBase+int(switchID-1))

	conn, err := net.Dial("tcp", switchAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to switch %d at %s: %w", switchID, switchAddress, err)
	}
	defer conn.Close()

	// Write the OpenFlow message to the switch
	_, err = message.WriteTo(conn)
	if err != nil {
		return fmt.Errorf("failed to send message to switch %d: %w", switchID, err)
	}

	log.Printf("Message sent to switch %d at %s successfully", switchID, switchAddress)
	return nil
}
