# easy pubsub [![Go Reference](https://pkg.go.dev/badge/github.com/buraksecer/go-easy-pubsub.svg)](https://pkg.go.dev/github.com/buraksecer/go-easy-pubsub) [![Go Report Card](https://goreportcard.com/badge/github.com/buraksecer/go-easy-pubsub)](https://goreportcard.com/report/github.com/buraksecer/go-easy-pubsub) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/buraksecer/go-easy-pubsub) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/buraksecer/go-easy-pubsub)


### Installation

```
go get github.com/buraksecer/go-easy-pubsub v0.0.3
```


------------



### Example

Firstly, you must Init new topic operation.

```go
topic.Init(clientId)
```

You can create basicly a client

```go
topic.Create(topicName)
```
You can list the topic of the client

```go
topics := topic.Topics()

for _, v := range topics.Topics {
	fmt.Println(v)
}

```

After, you can create a new subscription

```go
topic.CreateSubscription(topicName, subName)
```

You can list the subscription of the client

```go
topic.Subscriptions(topicName)
```

After finally, let's go Publish first message :)

```go
topic.Publish(topicName, "Fly me to the moon..")
```

And exactly, you must pull the message -_-

```go
topic.Receive(subName)
```


------------


Google pub/sub service easy way to use
