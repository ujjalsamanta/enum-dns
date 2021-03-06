// Copyright 2016 Hadrien Kohl hadrien.kohl@gmail.com
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

package enum

import "testing"

func Test_OverlapsWith(t *testing.T) {

	r := NumberRange{Lower: 400000000000000, Upper: 500000000000000}

	tt := []struct {
		r   NumberRange
		exp bool
	}{
		{NumberRange{Lower: 400000000000000, Upper: 500000000000000}, true},
		{NumberRange{Lower: 399999999999999, Upper: 500000000000000}, true},
		{NumberRange{Lower: 450000000000000, Upper: 500000000000000}, true},
		{NumberRange{Lower: 400000000000000, Upper: 500000000000001}, true},
		{NumberRange{Lower: 400000000000000, Upper: 450000000000000}, true},
		{NumberRange{Lower: 399999999999999, Upper: 500000000000001}, true},

		{NumberRange{Lower: 500000000000001, Upper: 500000000000002}, false},
		{NumberRange{Lower: 399999999999997, Upper: 399999999999998}, false},
	}

	for _, v := range tt {
		if r.OverlapWith(v.r) != v.exp {
			t.Errorf("[%d:%d].OverlapWith([%d:%d]) returned %d, expected %d",
				r.Lower, r.Upper,
				v.r.Lower, v.r.Upper,
				r.OverlapWith(v.r), v.exp,
			)
		}
	}

}

func Test_Starts(t *testing.T) {

	r := NumberRange{Lower: 400000000000000, Upper: 500000000000000}

	tt := []struct {
		r   NumberRange
		exp bool
	}{
		// Inside
		{NumberRange{Lower: 400000000000000, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000001, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000000, Upper: 499999999999999}, true},
		{NumberRange{Lower: 400000000000001, Upper: 499999999999999}, false},

		// Overlapping
		{NumberRange{Lower: 399999999999999, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000000, Upper: 500000000000001}, false},
		{NumberRange{Lower: 399999999999999, Upper: 500000000000001}, false},

		// Outside
		{NumberRange{Lower: 500000000000001, Upper: 500000000000002}, false},
		{NumberRange{Lower: 399999999999997, Upper: 399999999999998}, false},
	}

	for _, v := range tt {
		if r.Starts(v.r) != v.exp {
			t.Errorf("[%d:%d].Starts([%d:%d]) returned %t, expected %t",
				r.Lower, r.Upper,
				v.r.Lower, v.r.Upper,
				r.Starts(v.r), v.exp,
			)
		}
	}
}

func Test_Finishes(t *testing.T) {

	r := NumberRange{Lower: 400000000000000, Upper: 500000000000000}

	tt := []struct {
		r   NumberRange
		exp bool
	}{
		// Inside
		{NumberRange{Lower: 400000000000000, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000001, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000000, Upper: 499999999999999}, false},
		{NumberRange{Lower: 400000000000001, Upper: 499999999999999}, false},

		// Overlapping
		{NumberRange{Lower: 399999999999999, Upper: 500000000000000}, true},
		{NumberRange{Lower: 400000000000000, Upper: 500000000000001}, false},
		{NumberRange{Lower: 399999999999999, Upper: 500000000000001}, false},

		// Outside
		{NumberRange{Lower: 500000000000001, Upper: 500000000000002}, false},
		{NumberRange{Lower: 399999999999997, Upper: 399999999999998}, false},
	}

	for _, v := range tt {
		if r.Finishes(v.r) != v.exp {
			t.Errorf("[%d:%d].Starts([%d:%d]) returned %t, expected %t",
				r.Lower, r.Upper,
				v.r.Lower, v.r.Upper,
				r.Finishes(v.r), v.exp,
			)
		}
	}
}

func Test_Equals(t *testing.T) {

	r := NumberRange{Lower: 400000000000000, Upper: 500000000000000}

	tt := []struct {
		r   NumberRange
		exp bool
	}{
		// Inside
		{NumberRange{Lower: 400000000000000, Upper: 500000000000000}, true},
		{NumberRange{Lower: 400000000000001, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000000, Upper: 499999999999999}, false},
		{NumberRange{Lower: 400000000000001, Upper: 499999999999999}, false},

		// Overlapping
		{NumberRange{Lower: 399999999999999, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000000, Upper: 500000000000001}, false},
		{NumberRange{Lower: 399999999999999, Upper: 500000000000001}, false},

		// Outside
		{NumberRange{Lower: 500000000000001, Upper: 500000000000002}, false},
		{NumberRange{Lower: 399999999999997, Upper: 399999999999998}, false},
	}

	for _, v := range tt {
		if r.Equals(v.r) != v.exp {
			t.Errorf("[%d:%d].Starts([%d:%d]) returned %t, expected %t",
				r.Lower, r.Upper,
				v.r.Lower, v.r.Upper,
				r.Equals(v.r), v.exp,
			)
		}
	}
}

func Test_Contains(t *testing.T) {

	r := NumberRange{Lower: 400000000000000, Upper: 500000000000000}

	tt := []struct {
		r   NumberRange
		exp bool
	}{
		// Inside
		{NumberRange{Lower: 400000000000000, Upper: 500000000000000}, true},
		{NumberRange{Lower: 400000000000001, Upper: 500000000000000}, true},
		{NumberRange{Lower: 400000000000000, Upper: 499999999999999}, true},
		{NumberRange{Lower: 400000000000001, Upper: 499999999999999}, true},

		// Overlapping
		{NumberRange{Lower: 399999999999999, Upper: 500000000000000}, false},
		{NumberRange{Lower: 400000000000000, Upper: 500000000000001}, false},
		{NumberRange{Lower: 399999999999999, Upper: 500000000000001}, false},

		// Outside
		{NumberRange{Lower: 500000000000001, Upper: 500000000000002}, false},
		{NumberRange{Lower: 399999999999997, Upper: 399999999999998}, false},
	}

	for _, v := range tt {
		if r.Contains(v.r) != v.exp {
			t.Errorf("[%d:%d].Contains([%d:%d]) returned %t, expected %t",
				r.Lower, r.Upper,
				v.r.Lower, v.r.Upper,
				r.Contains(v.r), v.exp,
			)
		}
	}

}
