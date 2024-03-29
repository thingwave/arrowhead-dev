# Eclipse Arrowhead Local Cloud Test Client
This utility aplication is a Python3 based HTTP(S) client that performs a GET request to the submitted URI. This can for example be used to test the availability of Eclipse Arrowhead Core systems, or as a software base for developing more complete Eclipse Arrowhead systems and clients.

This tool can be used together the Eclipse Arrowhead [Reference implementation](https://www.github.com/eclipse-arrowhead/core-java-spring).
Note that this tool currently only supports insecure mode with HTTP, or secure mode with HTTPS using PEM-based certificates.

### To run the application
To run:
```
python3 lcclient.py --cloud=./certificates/testcloud2.pem --cert=./certificates/serviceregistry.pem --key=./certificates/serviceregistry.key https://127.0.0.1:8443/serviceregistry/echo
python3 lcclient.py --insecure http://127.0.0.1:8443/serviceregistry/echo
```

## Future work
1. JSON and Eclipse Arrowhead 4.4.0 data models.
2. More integrated features to interact with the Service registry, Orchestration, and Authorization systems.

