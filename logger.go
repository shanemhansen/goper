package goper

import "io"

const CLR_R = "\x1b[31;1m"
const CLR_N = "\x1b[0m"

type ColourStream struct {
	W io.Writer
}

func (c ColourStream) Write(p []byte) (int, error) {
	m, _ := c.W.Write([]byte(CLR_R))
	n, err := c.W.Write(p)
	o, _ := c.W.Write([]byte(CLR_N))
	return m + n + o, err
}
