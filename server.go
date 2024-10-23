package openuem_nats

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type MessageServer struct {
	Host           string
	Port           string
	CAcert         *x509.Certificate
	ClientCertPath string
	ClientKeyPath  string
	Connection     *nats.Conn
}

func New(host, port, clientCertPath, clientKeyPath string, caCert *x509.Certificate) *MessageServer {
	ms := MessageServer{
		Host:           host,
		Port:           port,
		CAcert:         caCert,
		ClientCertPath: clientCertPath,
		ClientKeyPath:  clientKeyPath,
	}
	return &ms
}

func (ms *MessageServer) Connect() error {
	// Load our client cert
	cert, err := tls.LoadX509KeyPair(ms.ClientCertPath, ms.ClientKeyPath)
	if err != nil {
		return err
	}

	// Create cert pool for our CA
	certPool := x509.NewCertPool()
	certPool.AddCert(ms.CAcert)

	// Prepare NATS tls config
	config := &tls.Config{
		ServerName:   ms.Host,
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
	}

	// Prepare URL
	url := fmt.Sprintf("tls://%s:%s", ms.Host, ms.Port)

	// Connect
	c, err := nats.Connect(
		url,
		nats.Secure(config),
		nats.MaxReconnects(-1),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Println("[INFO]: reconnected to the message broker")
		}),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			if err != nil {
				log.Printf("[INFO]: disconnected from message broker due to: %s, will attempt reconnect", err.Error())
			}
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("[INFO]: connection closed. Reason: %q\n", nc.LastError())
		}),
	)

	if err != nil {
		return err
	}

	log.Println("[INFO]: connection established with NATS server")
	ms.Connection = c
	return nil
}

func (ms *MessageServer) Close() error {
	ms.Connection.Close()
	return nil
}
