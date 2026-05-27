package nosurf

// Masks/unmasks the given data *in place*
// with the given key
// Slices must be of the same length, or oneTimePad will panic
func oneTimePad(data, key []byte) { _ = "STUB: not implemented"; return }

func maskToken(data []byte) []byte { _ = "STUB: not implemented"; return nil }

// tokenLength*2 == len(enckey + token)

// the first half of the result is the OTP
// the second half is the masked token itself

// generate the random token

func unmaskToken(data []byte) []byte { _ = "STUB: not implemented"; return nil }
