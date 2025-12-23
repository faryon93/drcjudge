package drc

import (
	"encoding/json"
	"os"
)

// drcjudge
// Copyright (C) 2025 Maximilian Pachl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// ---------------------------------------------------------------------------------------
//  types
// ---------------------------------------------------------------------------------------

type Result struct {
	Errors []*Error `json:"errors"`
}

// ---------------------------------------------------------------------------------------
//  public functions
// ---------------------------------------------------------------------------------------

// LoadResult loads a DRC result from the filesystem.
func LoadResult(filePath string) (*Result, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var drcResult Result
	return &drcResult, json.Unmarshal(buf, &drcResult)
}
