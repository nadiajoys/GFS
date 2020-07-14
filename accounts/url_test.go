// Copyright 2018 The go-gfscore Authors
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

package accounts

import (
	"testing"
)

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://gfscore.org")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "gfscore.org" {
		t.Errorf("expected: %v, got: %v", "gfscore.org", url.Path)
	}

	_, err = parseURL("gfscore.org")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "gfscore.org"}
	if url.String() != "https://gfscore.org" {
		t.Errorf("expected: %v, got: %v", "https://gfscore.org", url.String())
	}

	url = URL{Scheme: "", Path: "gfscore.org"}
	if url.String() != "gfscore.org" {
		t.Errorf("expected: %v, got: %v", "gfscore.org", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "gfscore.org"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://gfscore.org\"" {
		t.Errorf("expected: %v, got: %v", "\"https://gfscore.org\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://gfscore.org\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "gfscore.org" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "gfscore.org"}, URL{"https", "gfscore.org"}, 0},
		{URL{"http", "gfscore.org"}, URL{"https", "gfscore.org"}, -1},
		{URL{"https", "gfscore.org/a"}, URL{"https", "gfscore.org"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "gfscore.org"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
