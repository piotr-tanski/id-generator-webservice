package generator

import "time"

const generator_epoch = 1735689600000 // Unix timestamp (in ms) for Wed Jan 01 2025 00:00:00 GMT+0000

type IdGenerator struct {
	machine_id int
	timestamp  int64
	seqno      int
}

func NewGenerator(machine_id int) IdGenerator {
	return IdGenerator{machine_id: machine_id, timestamp: 0, seqno: 1}
}

func (generator *IdGenerator) Generate() (int64, int64, int, int) {
	t := time.Now().UnixMilli() - generator_epoch // Time since generator_epoch
	if t == generator.timestamp {
		// Increase the sequence number if multiple IDs with the same unix timestamp.
		generator.seqno += 1
	} else {
		generator.timestamp = t
		generator.seqno = 1
	}

	// The ID occupies 64 bits.
	// 1 reserved bit (0) - 41 bits for timestamp (in ms) - 12 bits for machine id - 10 bits for sequence number
	// 	41 bits for timestamp in ms mean ~70 years,
	// 	12 bits for machine id mean ~4,000 machines,
	// 	10 bits for sequence number mean 1024 IDs per millisecond per machine.

	id := int64(0)
	id |= (int64(generator.seqno) & 0x3FF)
	id |= ((int64(generator.machine_id) & 0xFFF) << 10)
	id |= ((int64(generator.timestamp) & 0x1FFFFFFFFFF) << 22)
	return id, t, generator.machine_id, generator.seqno
}
