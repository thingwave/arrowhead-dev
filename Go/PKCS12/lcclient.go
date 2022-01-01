/********************************************************************************
 * Copyright (c) 2021 ThingWave AB
 *
 * This program and the accompanying materials are made available under the
 * terms of the Eclipse Public License 2.0 which is available at
 * http://www.eclipse.org/legal/epl-2.0.
 *
 * SPDX-License-Identifier: EPL-2.0
 ********************************************************************************/
package main

import (
  "os"
  "fmt"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	certFile = flag.String("cert", "", "A PEM encoded System certificate file.")
	keyFile  = flag.String("key", "", "A PEM encoded System private key file.")
	cloudCaFile = flag.String("cloud", "", "A PEM encoced Local cloud CA certificate file.")
	help = flag.Bool("help", false, "To display a help text.")
)


func getRequest(client *http.Client, uri string) ([]byte, int, error) {
  resp, err := client.Get(uri)
	if err != nil {
    return nil, -1, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    return nil, -1, err
	}

  return data, resp.StatusCode, nil
}


func main() {
  fmt.Println("Eclipse Arrowhead Local cloud HTTPS client tool\nⒸ ThingWave AB")

  flag.Parse()
  if *help == true {
    fmt.Println("Usage example:\nlcclient --cloud=./certificates/testcloud2.pem --cert=./certificates/serviceregistry.pem --key=./certificates/serviceregistry.key <uri>")
    os.Exit(0)
  }
  if *certFile == "" || *keyFile == "" || *cloudCaFile == "" {
    fmt.Println("Missing arguments!\nUse --help to print help\n")
    os.Exit(-1)
  }

	// load client certificate
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// load local cloud CA certificate
	caCert, err := ioutil.ReadFile(*cloudCaFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
    InsecureSkipVerify: false,
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

  uri := flag.Args()[0]
  data, statusCode, err := getRequest(client, uri)

  if err != nil {
    fmt.Printf("Couldn't connect to %s!\n", uri)
    os.Exit(1)
  } else {
    fmt.Printf("Response code: %v\n", int(statusCode))
	  fmt.Printf("Response data:\n%s\n", string(data))
  }

  os.Exit(0)
}
