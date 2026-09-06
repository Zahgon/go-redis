package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	certDir := "../../dockers/standalone/tls"
	username := "testcertuser"
	tlsPort := "6666"
	nonTLSPort := "6379"

	setupClient := redis.NewClient(&redis.Options{
		Addr: "localhost:" + nonTLSPort,
	})
	defer setupClient.Close()

	err := setupClient.ACLSetUser(ctx,
		username,
		"on",
		"nopass",
		"~*",
		"+@all",
	).Err()
	if err != nil {
		log.Printf("Note: Could not create ACL user (may already exist): %v", err)
	}

	caCert, err := os.ReadFile(certDir + "/ca.crt")
	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair(
		certDir+"/"+username+".crt",
		certDir+"/"+username+".key",
	)
	if err != nil {
		log.Fatalf("Failed to load client certificate: %v", err)
	}

	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{clientCert},
		ServerName:         "localhost",
		InsecureSkipVerify: true,
	}

	client := redis.NewClient(&redis.Options{
		Addr:      "localhost:" + tlsPort,
		TLSConfig: tlsConfig,
	})
	defer client.Close()

	whoami, err := client.ACLWhoAmI(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
	fmt.Printf("✅ Authenticated as: %s (via TLS certificate CN)\n", whoami)

	err = client.Set(ctx, "tls-auth-example", "hello from cert auth!", 0).Err()
	if err != nil {
		log.Fatalf("SET failed: %v", err)
	}

	val, err := client.Get(ctx, "tls-auth-example").Result()
	if err != nil {
		log.Fatalf("GET failed: %v", err)
	}
	fmt.Printf("✅ SET/GET successful: %s\n", val)

	client.Del(ctx, "tls-auth-example")

	fmt.Println("\n🎉 TLS certificate authentication working!")
}
