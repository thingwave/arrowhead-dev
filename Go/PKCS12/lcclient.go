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
  //"errors"
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
  //resp, err := client.Get("https://127.0.0.1:8443/serviceregistry/echo")
  resp, err := client.Get(uri)
	if err != nil {
		//log.Fatal(err)
    return nil, -1, err
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatal(err)
    return nil, -1, err
	}

  return data, resp.StatusCode, nil
}

func main() {
  fmt.Println("Eclipse Arrowhead PKCS #12 certificate example\nⒸ ThingWave AB")

  flag.Parse()
  if *certFile == "" || *keyFile == "" || *cloudCaFile == "" {
    fmt.Printf("Missing arguments!\nUse --help to print help")
  }

  if *help == true {

    return
  }

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(*cloudCaFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

  fmt.Println("URI:", flag.Args()[0])
  uri := flag.Args()[0]
  data, statusCode, err := getRequest(client, uri)

  if err != nil {
    fmt.Printf("Couldn't connect!\n")
  } else {
    fmt.Printf("Response code: %v\n", int(statusCode))
	  fmt.Printf("Response data:\n%s\n", string(data))
  }
}
