// Package core holds all the data structs for streamd.
package core

import (
	"time"

	"github.com/openspock/ds/base"
	"github.com/openspock/ds/queues"
)

const (
	hashSize  = 64
	keySize   = 512
	blockSize = 4 * 1024 * 1024
	qSize     = 4 * 1024
)

var eventQ *queues.ConcurrentEnqueueFifo

func init() {
	eventQ = queues.MakeFifo(qSize)
}

// Event represents a unique, identifiable and temporal event.
//
// The time dimension of this event is governed by when this event
// was received for processing.Event
//
// Each Event also stores an array of hashes this record can be linked with.
//
// The maximum size of an event cannot be larger than the maximum size
// of a Block. Event larger than the maximum Block size will have to be
// handled externally.
type Event struct {
	Size       uint32         // Size of an event in bytes.
	Salt       [8]byte        // A random salt
	When       time.Time      // When this was received for processing.
	Key        [keySize]byte  // Key of the owner of this record.
	Hash       [hashSize]byte // A hash of the data
	PrevHash   [hashSize]byte // A bash of the immediately previous record.
	LinkedHash []base.Pair    // array of pairs of weighted hashes to link records.
	Data       []byte         // The actual data
}

// Block is a temporally ordered series of events. Blocks are usually
// of a fixed size of 4 MB. The motivation behind this to make it easier
// to serializes a Block of Event on disk. Think of Block as a buffer, which
// when full, is serialized to disk.
//
// A Block may be linked to several other blocks as the underlying Event
// entries could be linked to Event entires in other Block s. This relationship
// is governed by the underlying Event and not by the Block itself.
type Block struct {
	Size        uint32    // Size of the block, in bytes. Should not exceed 4 MB.
	RecordCount uint64    // A count of Event.
	Records     []Event   // A series of Event.
	When        time.Time // time when this block was serialized.
}

// Offer enables to offer an event to the event queue. The event queue will accept
// the event successfully or throw an error, if full.
func Offer(e Event) error {
	if err := eventQ.Enqueue(e); err != nil {
		return err
	}
	return nil
}


