# fcm
[![GoDoc](https://godoc.org/github.com/edganiukov/fcm?status.svg)](https://godoc.org/github.com/edganiukov/fcm)
[![Build Status](https://travis-ci.org/edganiukov/fcm.svg?branch=master)](https://travis-ci.org/edganiukov/fcm)
[![Go Report Card](https://goreportcard.com/badge/github.com/edganiukov/fcm)](https://goreportcard.com/report/github.com/edganiukov/fcm)

Golang client library for Firebase Cloud Messaging. Implemented only [HTTP client](https://firebase.google.com/docs/cloud-messaging/http-server-ref#downstream).

More information on [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging/)

### Getting Started
-------------------
To install fcm, use `go get`:

```bash
go get github.com/edganiukov/fcm
```
or `dep`, add to the `Gopkg.toml` file dependency:

```toml
[[constraint]]
  name = "github.com/edganiukov/fcm"
  version = "0.3.0"
```
and then run:

```bash
dep ensure
```

### Sample Usage
----------------
Here is a simple example illustrating how to use FCM library:
```go
package main

import (
	"github.com/edganiukov/fcm"
)

func main() {
	// Create the message to be sent.
	msg := &fcm.Message{
     		Token: "sample_device_token",
      		Data: map[string]interface{}{
         		"foo": "bar",
      		},
  	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("sample_api_key")
	if err != nil {
  		log.Fatal(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		/* ... */
	}
	/* ... */
}
```


#### TODO:
---------
- [ ] Retry only failed messages while multicast messaging.
