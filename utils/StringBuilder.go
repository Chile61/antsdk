package utils

import (
	"bytes"
	"fmt"
)

// StringBuilder StringBuilder
type StringBuilder struct {
	buf bytes.Buffer
}

// NewStringBuilder NewStringBuilder
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{buf: bytes.Buffer{}}
}

// Append Append
func (c *StringBuilder) Append(obj interface{}) *StringBuilder {
	c.buf.WriteString(fmt.Sprintf("%v", obj))
	return c
}

// ToString ToString
func (c *StringBuilder) ToString() string {
	return c.buf.String()
}
