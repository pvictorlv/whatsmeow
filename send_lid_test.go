package whatsmeow

import (
	"context"
	"testing"

	"go.mau.fi/whatsmeow/types"
)

// Every PN device already mapped: no usync round trip may happen. The client
// has no socket, so a query would panic or fail.
func TestFillMissingLIDsFromServerSkipsMappedDevices(t *testing.T) {
	pn := types.NewADJID("5511999999999", 0, 3)
	lid := types.NewADJID("123456789012345", 0, 3)
	lid.Server = types.HiddenUserServer
	mappings := map[types.JID]types.JID{pn: lid}

	(&Client{}).fillMissingLIDsFromServer(context.Background(), []types.JID{pn}, mappings)

	if mappings[pn] != lid {
		t.Fatalf("mapping changed: %v", mappings[pn])
	}
}

func TestFillMissingLIDsFromServerNoDevices(t *testing.T) {
	(&Client{}).fillMissingLIDsFromServer(context.Background(), nil, nil)
}
