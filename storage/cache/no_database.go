// Copyright 2020 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cache

import "fmt"

var NoDatabaseError = fmt.Errorf("no database specified for cachestore")

type NoDatabase struct{}

func (NoDatabase) Close() error {
	return NoDatabaseError
}

func (NoDatabase) SetList(prefix, name string, items []string) error {
	return NoDatabaseError
}

func (NoDatabase) GetList(prefix, name string, begin int, end int) ([]string, error) {
	return nil, NoDatabaseError
}

func (NoDatabase) GetString(prefix, name string) (string, error) {
	return "", NoDatabaseError
}

func (NoDatabase) SetString(prefix, name string, val string) error {
	return NoDatabaseError
}

func (NoDatabase) GetInt(prefix, name string) (int, error) {
	return 0, NoDatabaseError
}

func (NoDatabase) SetInt(prefix, name string, val int) error {
	return NoDatabaseError
}