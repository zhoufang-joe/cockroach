// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package kvevent

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/util/randutil"
	"github.com/stretchr/testify/assert"
)

func TestBufferEntryQueue(t *testing.T) {
	rand, _ := randutil.NewTestRand()

	q := bufferEventChunkQueue{}

	// Add one event and remove it
	assert.True(t, q.empty())
	q.enqueue(Event{})
	assert.False(t, q.empty())
	_, ok := q.dequeue()
	assert.True(t, ok)
	assert.True(t, q.empty())

	// Fill 5 chunks and then pop each one, ensuring empty() returns the correct
	// value each time.
	eventCount := bufferEventChunkArrSize * 5
	for i := 0; i < eventCount; i++ {
		q.enqueue(Event{})
	}
	for {
		assert.Equal(t, eventCount <= 0, q.empty())
		_, ok = q.dequeue()
		if !ok {
			assert.True(t, q.empty())
			break
		} else {
			eventCount--
		}
	}
	assert.Equal(t, 0, eventCount)
	q.enqueue(Event{})
	assert.False(t, q.empty())
	q.dequeue()
	assert.True(t, q.empty())

	// Add events to fill 5 chunks and assert they are consumed in fifo order.
	eventCount = bufferEventChunkArrSize * 5
	lastPop := -1
	lastPush := -1

	for eventCount > 0 {
		op := rand.Intn(2)
		if op == 0 {
			q.enqueue(Event{approxSize: lastPush + 1})
			lastPush++
		} else {
			e, ok := q.dequeue()
			if !ok {
				assert.Equal(t, lastPop, lastPush)
				assert.True(t, q.empty())
			} else {
				assert.Equal(t, e.approxSize, lastPop+1)
				lastPop++
				eventCount--
			}
		}
	}

	// Verify that purging works.
	eventCount = bufferEventChunkArrSize * 2.5
	for eventCount > 0 {
		q.enqueue(Event{approxSize: lastPush + 1})
		eventCount--
	}
	q.purge()
	assert.True(t, q.empty())
}
