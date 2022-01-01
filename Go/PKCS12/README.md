# Eclipse Arrowhead Local Cloud Test Client
This folder shows how to use PKCS#12 certificates in Golang with Eclipse Arrowhead.
This mini project contains a Golang based HTTPS client that performs a GET request to the submitted URI. This can for example be used to test the availability of Eclipse Arrowhead Core systems, or and a software base for developing more complete Eclipse Arrowhead systems and clients.

This tool can be used together the Eclipse Arrowhead [Reference implementation](https://www.github.com/eclipse-arrowhead/core-java-spring).
Note that this tool currently only supports HTTPS with PEM-based certificates.

### To build the project
To compile, run:
```
go build lcclient.go
```

### To run the application
To run:
```
./lcclient --cloud=./certificates/testcloud2.pem --cert=./certificates/serviceregistry.pem --key=./certificates/serviceregistry.key https://127.0.0.1:8443/serviceregistry/echo
```

## Future work
1. HTTP only (no certificates)
2. PKCS #12 certificates
3. More integrated features to interact with the Service registry, Orchestation and Authorization systems

