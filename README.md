# PubSub testing

This repo shows how to create an interfaces for the Google pubsub package.

Generating mocks from interfaces using `mockgen`:

```
mockgen --build_flags=--mod=mod -source=subscriber/subscriber.go -destination=mocks/subscriber.go
```
