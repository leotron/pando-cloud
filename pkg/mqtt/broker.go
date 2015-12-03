package mqtt

import (
	"net"
)

type Broker struct {
	mgr *Manager
}

func NewBroker(p Provider) *Broker {
	// context
	mgr := NewManager(p)

	handler := &Broker{mgr: mgr}

	return handler
}

func (b *Broker) Handle(conn net.Conn) {
	b.mgr.NewConn(conn)
}

func (b *Broker) SendMessageToDevice(deviceid uint64, msgtype string, message []byte) error {
	msg := &Publish{}
	msg.Header.QosLevel = QosAtLeastOnce
	msg.TopicName = DeviceIdToClientId(deviceid) + "/" + msgtype
	msg.Payload = BytesPayload(message)
	return b.mgr.PublishMessage2Device(deviceid, msg)
}

func (b *Broker) GetToken(deviceid uint64) ([]byte, error) {
	return b.mgr.GetToken(deviceid)
}
