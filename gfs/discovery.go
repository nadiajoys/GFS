// Copyright 2019 The go-gfscore Authors
// This file is part of the go-gfscore library.
//
// The go-gfscore library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-gfscore library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-gfscore library. If not, see <http://www.gnu.org/licenses/>.

package gfs

import (
	"github.com/gfscore/go-gfscore/core"
	"github.com/gfscore/go-gfscore/core/forkid"
	"github.com/gfscore/go-gfscore/p2p"
	"github.com/gfscore/go-gfscore/p2p/dnsdisc"
	"github.com/gfscore/go-gfscore/p2p/enode"
	"github.com/gfscore/go-gfscore/rlp"
)

// ethEntry is the "gfs" ENR entry which advertises gfs protocol
// on the discovery network.
type ethEntry struct {
	ForkID forkid.ID // Fork identifier per EIP-2124

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e ethEntry) ENRKey() string {
	return "gfs"
}

// startEthEntryUpdate starts the ENR updater loop.
func (gfs *Gfscore) startEthEntryUpdate(ln *enode.LocalNode) {
	var newHead = make(chan core.ChainHeadEvent, 10)
	sub := gfs.blockchain.SubscribeChainHeadEvent(newHead)

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case <-newHead:
				ln.Set(gfs.currentEthEntry())
			case <-sub.Err():
				// Would be nice to sync with gfs.Stop, but there is no
				// good way to do that.
				return
			}
		}
	}()
}

func (gfs *Gfscore) currentEthEntry() *ethEntry {
	return &ethEntry{ForkID: forkid.NewID(gfs.blockchain)}
}

// setupDiscovery creates the node discovery source for the gfs protocol.
func (gfs *Gfscore) setupDiscovery(cfg *p2p.Config) (enode.Iterator, error) {
	if cfg.NoDiscovery || len(gfs.config.DiscoveryURLs) == 0 {
		return nil, nil
	}
	client := dnsdisc.NewClient(dnsdisc.Config{})
	return client.NewIterator(gfs.config.DiscoveryURLs...)
}
