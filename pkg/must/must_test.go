// SPDX-FileCopyrightText: Copyright The Lima Authors
// SPDX-License-Identifier: Apache-2.0

package must

import (
	"errors"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMustSuccess(t *testing.T) {
	expected := "hello"
	actual := Must(expected, nil)
	assert.Equal(t, expected, actual)

	expectedInt := 42
	actualInt := Must(expectedInt, nil)
	assert.Equal(t, expectedInt, actualInt)
}

func TestMustPanic(t *testing.T) {
	expectedErr := errors.New("test error")
	defer func() {
		r := recover()
		assert.Assert(t, r != nil, "Must should have panicked")
		assert.Equal(t, expectedErr, r)
	}()

	Must("hello", expectedErr)
}
