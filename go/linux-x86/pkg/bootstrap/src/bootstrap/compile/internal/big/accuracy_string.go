// Do not edit. Bootstrap copy of /usr/local/google/buildbot/src/android/build-tools/out/obj/go/src/cmd/compile/internal/big/accuracy_string.go

//line /usr/local/google/buildbot/src/android/build-tools/out/obj/go/src/cmd/compile/internal/big/accuracy_string.go:1
// generated by stringer -type=Accuracy; DO NOT EDIT

package big

import "fmt"

const _Accuracy_name = "BelowExactAbove"

var _Accuracy_index = [...]uint8{0, 5, 10, 15}

func (i Accuracy) String() string {
	i -= -1
	if i < 0 || i+1 >= Accuracy(len(_Accuracy_index)) {
		return fmt.Sprintf("Accuracy(%d)", i+-1)
	}
	return _Accuracy_name[_Accuracy_index[i]:_Accuracy_index[i+1]]
}
