package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	fmt.Println("Example 1: TLS with InsecureSkipVerify")
	client1 := redis.NewClient(&redis.Options{
		Addr: "localhost:6666",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	defer client1.Close()

	if err := client1.Ping(ctx).Err(); err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	} else {
		fmt.Println("✅ Connected successfully with InsecureSkipVerify")
	}

	fmt.Println("\nExample 2: TLS with CA certificate verification")

	caCert, err := os.ReadFile("path/to/ca.crt")
	if err != nil {
		fmt.Printf("Note: CA cert not found (this is expected in this example): %v\n", err)
	} else {
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		client2 := redis.NewClient(&redis.Options{
			Addr: "localhost:6666",
			TLSConfig: &tls.Config{
				RootCAs:    caCertPool,
				ServerName: "localhost",
			},
		})
		defer client2.Close()

		if err := client2.Ping(ctx).Err(); err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
		} else {
			fmt.Println("✅ Connected successfully with CA verification")
		}
	}

	fmt.Println("\nExample 3: TLS with client certificate (mutual TLS)")

	caCert, err = os.ReadFile("path/to/ca.crt")
	if err != nil {
		fmt.Printf("Note: CA cert not found (this is expected in this example): %v\n", err)
	} else {
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		cert, err := tls.LoadX509KeyPair("path/to/client.crt", "path/to/client.key")
		if err != nil {
			fmt.Printf("Note: Client cert not found (this is expected in this example): %v\n", err)
		} else {
			client3 := redis.NewClient(&redis.Options{
				Addr: "localhost:6666",
				TLSConfig: &tls.Config{
					RootCAs:      caCertPool,
					Certificates: []tls.Certificate{cert},
					ServerName:   "localhost",
				},
			})
			defer client3.Close()

			if err := client3.Ping(ctx).Err(); err != nil {
				fmt.Printf("Failed to connect: %v\n", err)
			} else {
				fmt.Println("✅ Connected successfully with client certificate")
			}
		}
	}

	fmt.Println("\nExample 4: Using rediss:// URL scheme")

	opt, err := redis.ParseURL("rediss://localhost:6666")
	if err != nil {
		fmt.Printf("Failed to parse URL: %v\n", err)
		return
	}

	opt.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	client4 := redis.NewClient(opt)
	defer client4.Close()

	if err := client4.Ping(ctx).Err(); err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	} else {
		fmt.Println("✅ Connected successfully using rediss:// URL")
	}

	fmt.Println("\nExample 5: TLS with certificate-based authentication")
	fmt.Println("Note: This requires Redis 6.2+ with tls-auth-clients-user CN configuration")
	fmt.Println("The certificate's CN (Common Name) field will be used as the Redis username")
	fmt.Println("See tls_cert_auth_test.go for a complete working example")
}
