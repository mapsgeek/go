package main

import "github.com/mapsgeek/tolling/types"

type DataProducer interface {
	ProduceData(types.OBUData) error
}
