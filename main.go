package main

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
//  imports
// ---------------------------------------------------------------------------------------

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/faryon93/drcjudge/config"
	"github.com/faryon93/drcjudge/drc"
)

// ---------------------------------------------------------------------------------------
//  constants
// ---------------------------------------------------------------------------------------

const (
	DefaultConfigurationFile = "drcjudge.yml"
)

// ---------------------------------------------------------------------------------------
//  application entry
// ---------------------------------------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: drcjudge <drc-result.json>")
		os.Exit(1)
	}

	conf, err := config.Load(DefaultConfigurationFile)
	if err != nil {
		log.Println("failed to load configuration file", DefaultConfigurationFile, ":", err)
		os.Exit(1)
	}

	drcResultPath := os.Args[1]
	drcResult, err := drc.LoadResult(drcResultPath)
	if err != nil {
		log.Println("failed to load result file", drcResultPath, ":", err)
		os.Exit(1)
	}
	log.Printf("loaded %d drc errors", len(drcResult.Errors))

	// verify there are no DRC errors
	errorCount := 0
	for _, drcError := range drcResult.Errors {
		if conf.IgnoreApproved && drcError.IsApproved() {
			continue
		}

		// ignore by layer name
		if slices.Contains(conf.IgnoreLayers, drcError.LayerName) {
			continue
		}

		// ignore by layer number
		if slices.Contains(conf.IgnoreLayers, strconv.Itoa(drcError.LayerId)) {
			continue
		}

		// ignore by error class-code
		if conf.IsCodeIgnored(drcError.Code) {
			continue
		}

		log.Printf("drc #%d: Layer=%s, Class=%s, X=%.03f, Y=%.03f",
			errorCount, drcError.LayerName, drcError.Description, drcError.Geometry.X, drcError.Geometry.Y)
		errorCount++
	}

	// if there were drc errors -> exit with failure code
	if errorCount > 0 {
		log.Println("🚫 drc failed")
		os.Exit(5)
	} else {
		log.Println("✅ drc succeeded")
	}
}
