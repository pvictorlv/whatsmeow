package whatsmeow

import (
	"errors"
	"fmt"
	"testing"

	"go.mau.fi/libsignal/signalerror"
)

// Regressão: os cases vazios do switch não caíam no seguinte (Go não tem
// fall-through implícito) e esses erros saíam como "desconhecido".
func TestGetRetryReasonFromError(t *testing.T) {
	cases := map[error]int{
		signalerror.ErrBadMAC:                           RetryReasonSignalErrorBadMac,
		signalerror.ErrNoSessionForUser:                 RetryReasonSignalErrorNoSession,
		signalerror.ErrNoSenderKeyForUser:               RetryReasonSignalErrorNoSession,
		signalerror.ErrWrongMessageVersion:              RetryReasonSignalErrorInvalidMessage,
		signalerror.ErrOldMessageVersion:                RetryReasonSignalErrorInvalidMessage,
		signalerror.ErrUnknownMessageVersion:            RetryReasonSignalErrorInvalidMessage,
		signalerror.ErrIncompleteMessage:                RetryReasonSignalErrorInvalidMessage,
		signalerror.ErrInvalidSignature:                 RetryReasonSignalErrorInvalidSignature,
		signalerror.ErrSenderKeyStateVerificationFailed: RetryReasonSignalErrorInvalidSignature,
		signalerror.ErrNoSignedPreKey:                   RetryReasonSignalErrorInvalidKey,
		signalerror.ErrNoSenderKeyStateForID:            RetryReasonSignalErrorInvalidKeyId,
		signalerror.ErrTooFarIntoFuture:                 RetryReasonSignalErrorFutureMessage,
		errors.New("something else"):                    RetryReasonUnknownError,
	}
	for err, want := range cases {
		if got := getRetryReasonFromError(fmt.Errorf("wrapped: %w", err)); got != want {
			t.Errorf("%v: got %d, want %d", err, got, want)
		}
	}
}
