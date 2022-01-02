/********************************************************************************
 * Copyright (c) 2022 ThingWave AB
 *
 * This program and the accompanying materials are made available under the
 * terms of the Eclipse Public License 2.0 which is available at
 * http://www.eclipse.org/legal/epl-2.0.
 *
 * SPDX-License-Identifier: EPL-2.0
 ********************************************************************************/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//senml "github.com/thingwave/senml-go"
)

var (
	certFile    = flag.String("cert", "", "A PEM encoded System certificate file.")
	keyFile     = flag.String("key", "", "A PEM encoded System private key file.")
	cloudCaFile = flag.String("cloud", "", "A PEM encoced Local cloud CA certificate file.")
	insecure    = flag.Bool("insecure", false, "To disable TLS.")
	help        = flag.Bool("help", false, "To display a help text.")
	op          = flag.String("op", "", "Execute a command to an Eclipse Arrowhead Core system.")
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

//func postRequest(client *http.Client, uri string, payload string) ([]byte, int, error) {

//}

func loadPEMCertificates(cloudfile string, certfile string, keyfile string) (tls.Certificate, *x509.CertPool, error) {
	// load client certificate
	cert, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		return tls.Certificate{}, nil, err
	}

	// load local cloud CA certificate
	caCert, err := ioutil.ReadFile(cloudfile)
	if err != nil {
		return cert, nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return cert, caCertPool, nil
}

func main() {
	//var msg senml.SenMLMessage
	//fmt.Printf("SenML for the DataManager %v\n", msg)
	client := &http.Client{}
	fmt.Println("Eclipse Arrowhead Local cloud HTTPS client tool\nCopyright 2022 ThingWave AB")

	flag.Parse()
	if *help == true {
		fmt.Println("Usage example:\nlcclient --cloud=./certificates/testcloud2.pem --cert=./certificates/serviceregistry.pem --key=./certificates/serviceregistry.key <uri>")
		os.Exit(0)
	}
	if *insecure == false && (*certFile == "" || *keyFile == "" || *cloudCaFile == "") {
		fmt.Println("Missing arguments!\nUse --help to print help")
		os.Exit(-1)
	}

	if *insecure == false {
		cert, caCertPool, err := loadPEMCertificates(*cloudCaFile, *certFile, *keyFile)
		if err != nil {
			fmt.Println("Could not load certificates, exiting...")
			os.Exit(-1)
		}

		// create TLS config
		tlsConfig := &tls.Config{
			InsecureSkipVerify: false,
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		client = &http.Client{Transport: transport}
	}

	// perform the request
	uri := flag.Args()[0]
	var data []byte = nil
	var statusCode int
	var err error
	if *op == "" {
		data, statusCode, err = getRequest(client, uri)
	} else if *op == "listServices" {
		//data, statusCode, err = postRequest(client, uri, "")
	} else {
		fmt.Printf("Illegal operation, aborting...\n")
		os.Exit(-1)
	}

	// print response
	if err != nil {
		fmt.Printf("Couldn't connect to %s!\n", uri)
		os.Exit(1)
	} else {
		fmt.Printf("Response code: %v\n", int(statusCode))
		fmt.Printf("Response data:\n%s\n", string(data))
		os.Exit(0)
	}

}
