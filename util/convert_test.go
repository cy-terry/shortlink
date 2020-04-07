package util

import (
	"testing"
)

func TestHashURLConverter_ToShortURL(t *testing.T) {
	c := &HashURLConverter{}
	s := c.ToShortURL("http://39.99.146.109/polarion/#/project/AEM502PMTM/workitems/task?query=NOT%20HAS_VALUE%3Aresolution")
	t.Log(s)
}
