package main

import (
	"testing"
)

func TestVersion_isHigherThan(t *testing.T) {
	tests := map[string]struct{
		v1   version
		v2   version
		want bool
	}{
		"patch v1 higher": {
			v1:   version{major: 1, minor: 1, patch: 2},
			v2:   version{major: 1, minor: 1, patch: 1},
			want: true,
		},
		"patch v2 higher": {
			v1:   version{major: 2, minor: 2, patch: 1},
			v2:   version{major: 2, minor: 2, patch: 2},
			want: false,
		},
		"minor v1 higher": {
			v1:   version{major: 3, minor: 3, patch: 3},
			v2:   version{major: 3, minor: 2, patch: 5},
			want: true,
		},
		"minor v2 higher": {
			v1:   version{major: 4, minor: 3, patch: 3},
			v2:   version{major: 4, minor: 4, patch: 1},
			want: false,
		},
		"major v1 higher": {
			v1:   version{major: 6, minor: 3, patch: 1},
			v2:   version{major: 5, minor: 5, patch: 5},
			want: true,
		},
		"major v2 higher": {
			v1:   version{major: 6, minor: 6, patch: 6},
			v2:   version{major: 7, minor: 3, patch: 1},
			want: false,
		},
		"same v1 higher": {
			v1:   version{major: 7, minor: 7, patch: 7},
			v2:   version{major: 7, minor: 7, patch: 7},
			want: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(st *testing.T) {
			got := test.v1.isHigherThan(test.v2)
			if got != test.want {
				st.Errorf("expected v1 %v isHigherThan v2 %v to be %v, got %v", test.v1.Patch(), test.v2.Patch(), test.want, got)
			}
		})
	}
}
