// Copyright 2015 The go-gfscore Authors
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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/gfscore/go-gfscore/metrics"
)

var (
	headerInMeter      = metrics.NewRegisteredMeter("gfs/downloader/headers/in", nil)
	headerReqTimer     = metrics.NewRegisteredTimer("gfs/downloader/headers/req", nil)
	headerDropMeter    = metrics.NewRegisteredMeter("gfs/downloader/headers/drop", nil)
	headerTimeoutMeter = metrics.NewRegisteredMeter("gfs/downloader/headers/timeout", nil)

	bodyInMeter      = metrics.NewRegisteredMeter("gfs/downloader/bodies/in", nil)
	bodyReqTimer     = metrics.NewRegisteredTimer("gfs/downloader/bodies/req", nil)
	bodyDropMeter    = metrics.NewRegisteredMeter("gfs/downloader/bodies/drop", nil)
	bodyTimeoutMeter = metrics.NewRegisteredMeter("gfs/downloader/bodies/timeout", nil)

	receiptInMeter      = metrics.NewRegisteredMeter("gfs/downloader/receipts/in", nil)
	receiptReqTimer     = metrics.NewRegisteredTimer("gfs/downloader/receipts/req", nil)
	receiptDropMeter    = metrics.NewRegisteredMeter("gfs/downloader/receipts/drop", nil)
	receiptTimeoutMeter = metrics.NewRegisteredMeter("gfs/downloader/receipts/timeout", nil)

	stateInMeter   = metrics.NewRegisteredMeter("gfs/downloader/states/in", nil)
	stateDropMeter = metrics.NewRegisteredMeter("gfs/downloader/states/drop", nil)
)
