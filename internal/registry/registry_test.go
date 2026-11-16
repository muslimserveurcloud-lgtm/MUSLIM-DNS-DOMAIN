package registry

import "testing"

func TestRegisterDomain(t *testing.T) {
	r := New()
	if err := r.AddExtension("msc"); err != nil {
		t.Fatal(err)
	}
	if err := r.RegisterDomain("example.msc", "client-1"); err != nil {
		t.Fatal(err)
	}
	if _, ok := r.GetDomain("example.msc"); !ok {
		t.Fatal("domain not found")
	}
	if err := r.RegisterDomain("example.msc", "client-2"); err == nil {
		t.Fatal("duplicate domain accepted")
	}
}
