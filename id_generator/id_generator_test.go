package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkIDs(t *testing.T, prev_id int64, id int64, prev_time int64, time int64, prev_machine_id int, machine_id int, prev_seqno int, seqno int) {
	assert.Greater(t, id, prev_id, "IDs are not increasing")
	assert.GreaterOrEqual(t, time, prev_time, "Timestamps are not increasing")
	assert.Equal(t, machine_id, prev_machine_id, "Machine IDs are not the same")
	if time == prev_time {
		assert.GreaterOrEqual(t, seqno, prev_seqno, "Seqeunce numbers are not increasing for the same timestamp")
	}
}

func TestIdStructure(t *testing.T) {
	generator := NewGenerator(0)
	id, timestamp, machine_id, seqno := generator.Generate()

	reserved_bit := id >> 63
	assert.Equal(t, reserved_bit, int64(0), "Reserved bit: incorrect value!")

	encoded_timestamp := (id >> 22) & 0x1FFFFFFFFFF
	timestamp_41_bits := timestamp & 0x1FFFFFFFFFF
	assert.Equal(t, encoded_timestamp, timestamp_41_bits, "Timestamp: incorrect value!")

	encoded_machine_id := (id >> 10) & 0xFFF
	assert.Equal(t, encoded_machine_id, int64(machine_id), "Machine ID: incorrect value!")

	encoded_seqno := id & 0x3FF
	assert.Equal(t, encoded_seqno, int64(seqno), "Seqeunce number: incorrect value!")
}

func TestSubsequentIdGeneration(t *testing.T) {
	generator := NewGenerator(0)

	prev_id, prev_time, prev_machine_id, prev_seqno := generator.Generate()
	id, time, machine_id, seqno := generator.Generate()
	checkIDs(t, prev_id, id, prev_time, time, prev_machine_id, machine_id, prev_seqno, seqno)
}

func TestMultipleSubsequentIdGeneration(t *testing.T) {
	generator := NewGenerator(0)

	prev_id, prev_time, prev_machine_id, prev_seqno := generator.Generate()

	for i := 0; i < 1000; i++ {
		id, time, machine_id, seqno := generator.Generate()

		checkIDs(t, prev_id, id, prev_time, time, prev_machine_id, machine_id, prev_seqno, seqno)

		prev_id, prev_time, prev_machine_id, prev_seqno = id, time, machine_id, seqno
	}
}
