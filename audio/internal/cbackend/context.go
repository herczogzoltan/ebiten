// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ebitenginecbackend || ebitencbackend
// +build ebitenginecbackend ebitencbackend

package cbackend

import (
	"io"

	"github.com/hajimehoshi/ebiten/v2/internal/cbackend"
)

type Context struct {
	sampleRate      int
	channelCount    int
	bitDepthInBytes int

	players *players
}

func NewContext(sampleRate, channelCount, bitDepthInBytes int) (*Context, chan struct{}, error) {
	c := &Context{
		sampleRate:      sampleRate,
		channelCount:    channelCount,
		bitDepthInBytes: bitDepthInBytes,
		players:         newPlayers(),
	}
	cbackend.OpenAudio(sampleRate, channelCount, c.players.read)
	ready := make(chan struct{})
	close(ready)
	return c, ready, nil
}

func (c *Context) NewPlayer(src io.Reader) *Player {
	return newPlayer(c, src)
}

func (c *Context) Suspend() error {
	// Do nothing so far.
	return nil
}

func (c *Context) Resume() error {
	// Do nothing so far.
	return nil
}

func (c *Context) Err() error {
	return nil
}

func (c *Context) defaultBufferSize() int {
	return c.sampleRate * c.channelCount * c.bitDepthInBytes / 2 // 0.5[s]
}
