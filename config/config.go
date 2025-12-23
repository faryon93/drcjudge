package config

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
	"os"
	"slices"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/faryon93/drcjudge/drc"
)

// ---------------------------------------------------------------------------------------
//  types
// ---------------------------------------------------------------------------------------

// Config represents the settings a judgement session should run with.
type Config struct {
	IgnoreApproved bool     `yaml:"ignore_approved"`
	IgnoreLayers   []string `yaml:"ignore_layers"`
	IgnoreCode     []string `yaml:"ignore_code"`
}

// ---------------------------------------------------------------------------------------
//  public functions
// ---------------------------------------------------------------------------------------

// Load loads a configuration file from the filesystem.
func Load(filePath string) (*Config, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	conf := Default()
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return nil, err
	}

	for i, code := range conf.IgnoreCode {
		conf.IgnoreCode[i] = strings.ToLower(code)
	}

	return conf, nil
}

// Default returns the default configuration.
func Default() *Config {
	return &Config{
		IgnoreApproved: true,
		IgnoreLayers:   []string{},
		IgnoreCode:     []string{},
	}
}

// ---------------------------------------------------------------------------------------
//  public members
// ---------------------------------------------------------------------------------------

func (c *Config) IsCodeIgnored(code int) bool {
	// match by code number
	if slices.Contains(c.IgnoreCode, strconv.Itoa(code)) {
		return true
	}

	// are there string substitutions for the given code?
	codeNames := drc.CodeNames[code]
	if codeNames == nil {
		// no substitution -> can only be ignored by number
		// which was already checked before
		return false
	}

	// is any of the substitution names ignored?
	for _, codeName := range codeNames {
		if slices.Contains(c.IgnoreCode, codeName) {
			return true
		}
	}

	return false
}
