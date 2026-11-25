//go:build darwin

package keychain

import "testing"

func TestErrorMessage(t *testing.T) {
	for _, tc := range []struct {
		err  Error
		want string
	}{
		{ErrorItemNotFound, "The specified item could not be found in the keychain. (-25300)"},
		{ErrorReadOnly, "Read-only error. (-25292)"},
		{ErrorUserCanceled, "User canceled the operation. (-128)"},
		{Error(1), "Keychain Error. (1)"},
	} {
		if got := tc.err.Error(); got != tc.want {
			t.Errorf("Error(%d).Error() = %q, want %q", int(tc.err), got, tc.want)
		}
	}
}
