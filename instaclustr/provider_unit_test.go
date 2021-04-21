package instaclustr

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestProvider(t *testing.T) {
    if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}
