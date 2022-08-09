package connection

import "testing"


func TestConnection(t *testing.T){
	db := FetchConnection()
	if db == nil {
		t.Error("FetchConnection() failed")
	}
}