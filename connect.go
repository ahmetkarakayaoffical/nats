package openuem_nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func ConnectWithNATS(servers, clientCert, clientKey, caCert string) (*nats.Conn, error) {
	c, err := nats.Connect(
		servers,
		nats.RootCAs(caCert),
		nats.ClientCert(clientCert, clientKey),
		nats.MaxReconnects(-1),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Println("[INFO]: ... reconnected to the message broker")
		}),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			if err != nil {
				log.Printf("[INFO]: ... disconnected from message broker due to: %s, will attempt reconnect", err.Error())
			}
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("[INFO]: ... connection closed. Reason: %q\n", nc.LastError())
		}),
	)

	if err != nil {
		return nil, err
	}

	log.Println("[INFO] ... connection established with NATS server")
	return c, nil
}
