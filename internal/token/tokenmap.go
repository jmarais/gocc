//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package token

type TokenMap struct {
	IdMap   map[string]int
	TypeMap []string
	CharMap map[string]string
}

func NewTokenMap(symbols []string) *TokenMap {
	tm := &TokenMap{
		IdMap:   make(map[string]int),
		TypeMap: make([]string, len(symbols)),
		CharMap: nil,
	}

	for i, sym := range symbols {
		tm.IdMap[sym] = i
		tm.TypeMap[i] = sym
	}
	return tm
}

func (this *TokenMap) AddCharacterMap(c map[string]string) {
	this.CharMap = c
}
