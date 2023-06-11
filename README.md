## Golang-services-grpc

- Go-gin
- grpc
- consul
- mongodb
- docker

> Provider service generates some random huge info, receiver accepts it and structurizes it into MongoDB, Visualizer will show info about communications (date and time of text generation, date and time when it's structurized in db, difference in ms, symbols quantity, name of provider and receiver).
> Communication between provider and receiver is realized with grpc. Every service should be in docker container. Between provider and receiver there is a Consul, which tells provider to which one receiver to address.

Stress-test: 8 providers, 3 receivers, 1 visualizer, 1 MongoDB, 1 Consul
