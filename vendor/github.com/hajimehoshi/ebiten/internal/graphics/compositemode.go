// Copyright 2018 The Ebiten Authors
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

package graphics

import (
	"fmt"
)

type CompositeMode int

const (
	CompositeModeSourceOver CompositeMode = iota // This value must be 0 (= initial value)
	CompositeModeClear
	CompositeModeCopy
	CompositeModeDestination
	CompositeModeDestinationOver
	CompositeModeSourceIn
	CompositeModeDestinationIn
	CompositeModeSourceOut
	CompositeModeDestinationOut
	CompositeModeSourceAtop
	CompositeModeDestinationAtop
	CompositeModeXor
	CompositeModeLighter
	CompositeModeUnknown

	CompositeModeMax = CompositeModeLighter
)

type Operation int

const (
	Zero Operation = iota
	One
	SrcAlpha
	DstAlpha
	OneMinusSrcAlpha
	OneMinusDstAlpha
)

func (c CompositeMode) Operations() (src Operation, dst Operation) {
	switch c {
	case CompositeModeSourceOver:
		return One, OneMinusSrcAlpha
	case CompositeModeClear:
		return Zero, Zero
	case CompositeModeCopy:
		return One, Zero
	case CompositeModeDestination:
		return Zero, One
	case CompositeModeDestinationOver:
		return OneMinusDstAlpha, One
	case CompositeModeSourceIn:
		return DstAlpha, Zero
	case CompositeModeDestinationIn:
		return Zero, SrcAlpha
	case CompositeModeSourceOut:
		return OneMinusDstAlpha, Zero
	case CompositeModeDestinationOut:
		return Zero, OneMinusSrcAlpha
	case CompositeModeSourceAtop:
		return DstAlpha, OneMinusSrcAlpha
	case CompositeModeDestinationAtop:
		return OneMinusDstAlpha, SrcAlpha
	case CompositeModeXor:
		return OneMinusDstAlpha, OneMinusSrcAlpha
	case CompositeModeLighter:
		return One, One
	default:
		panic(fmt.Sprintf("graphics: invalid composite mode: %d", c))
	}
}
