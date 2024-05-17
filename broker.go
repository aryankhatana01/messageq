package main

type Config struct {
	ListenAddr        string
	StoreProducerFunc StoreProducerFunc
}

type Broker struct {
	storage Storer
}
