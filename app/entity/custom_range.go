package entity

import "fmt"

type CustomRange struct {
	InputStr      string
	StartSec      *float64
	EndSec        *float64
	StartSegIndex *int64
	EndSegIndex   *int64
}

func (c *CustomRange) ToString() string {
	startSec := 0.0
	endSec := 0.0
	startSegIndex := int64(0)
	endSegIndex := int64(0)

	if c.StartSec != nil {
		startSec = *c.StartSec
	}
	if c.EndSec != nil {
		endSec = *c.EndSec
	}
	if c.StartSegIndex != nil {
		startSegIndex = *c.StartSegIndex
	}
	if c.EndSegIndex != nil {
		endSegIndex = *c.EndSegIndex
	}

	return fmt.Sprintf("StartSec: %.2f, EndSec: %.2f, StartSegIndex: %d, EndSegIndex: %d",
		startSec, endSec, startSegIndex, endSegIndex)
}
