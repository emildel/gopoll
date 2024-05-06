package main

import "github.com/r3labs/sse/v2"

type StreamManager struct {
	StreamSlice []*sse.Stream
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		StreamSlice: make([]*sse.Stream, 0),
	}
}

func (s *StreamManager) AddStreamToManager(stream *sse.Stream) {
	s.StreamSlice = append(s.StreamSlice, stream)
}
