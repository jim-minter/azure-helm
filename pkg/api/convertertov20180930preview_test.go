package api

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// testContainerService and testOpenShiftCluster are defined in
// converterfromv20180930preview_test.go.

func TestConvertToV20180930preview(t *testing.T) {
	oc := ConvertToV20180930preview(testContainerService)
	if !reflect.DeepEqual(oc, testOpenShiftCluster) {
		t.Errorf("ConvertToV20180930preview returned unexpected result\n%s\n", spew.Sdump(oc))
	}
}
