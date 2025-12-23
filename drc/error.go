package drc

import "github.com/faryon93/drcjudge/helper"

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
//  constants
// ---------------------------------------------------------------------------------------

const (
	CodeKeepout  = 7
	CodeOverlap  = 19
	CodeAirwire  = 21
	CodeWireStub = 23

	StateNormal   = 0
	StateApproved = 2
)

var (
	CodeNames = map[int][]string{
		CodeKeepout: helper.StringMutate("keep", "out"),
		CodeOverlap: []string{
			"overlap",
		},
		CodeAirwire:  helper.StringMutate("air", "wire"),
		CodeWireStub: helper.StringMutate("wire", "stub"),
	}
)

// ---------------------------------------------------------------------------------------
//  types
// ---------------------------------------------------------------------------------------

// Error represents a single DRC error.
// For possible properties see: UL_ERROR (https://help.autodesk.com/view/fusion360/ENU/?guid=ECD-ULP-ERROR)
type Error struct {
	LayerId     int      `json:"layer_id"`
	LayerName   string   `json:"layer_name"`
	Code        int      `json:"code"`
	Description string   `json:"description"`
	State       int      `json:"state"`
	Geometry    Geometry `json:"geometry"`
}

// ---------------------------------------------------------------------------------------
//  public members
// ---------------------------------------------------------------------------------------

// IsApproved returns true if this error has been approved.
func (e *Error) IsApproved() bool {
	return e.State == StateApproved
}
