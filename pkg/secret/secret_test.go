package secret

import "testing"

func TestCreation(t *testing.T) {
	store := NewInMemory("testKey")

	if store.encryptionKey != "testKey" {
		t.Error("secret store was not created correctly")
	}
}

func TestGet(t *testing.T) {
	store := setupStore()
	val, err := store.Get("key1")
	if err != nil {
		t.Errorf("failed to retrieve value '%s' from the kv store", "key1")
	}

	val2, err := store.Get("key2")
	if err != nil {
		t.Errorf("failed to retrieve value '%s' from the kv store", "key2")
	}

	if val != "value1" {
		t.Errorf("unexpected value retrieved from kv store - got '%s', expected '%s'", val, "value1")
	}

	if val2 != "value2" {
		t.Errorf("unexpected value retrieved from kv store - got '%s', expected '%s'", val, "value2")
	}
}

func TestGetNonExisting(t *testing.T) {
	store := setupStore()
	_, err := store.Get("err1")
	if err == nil {
		t.Errorf("no error thrown when trying to read non-existent key '%s'", "err1")
	}
}

func TestSet(t *testing.T) {
	store := setupStore()

	err := store.Set("key3", "testValue")
	if err != nil {
		t.Error("error while setting value")
	}

	val := store.keyValues["key3"]
	if val == "" {
		t.Errorf("failed to insert the correct value - got '%s'", val)
	}

	if len(val) != 50 {
		t.Errorf("failed to insert the correct value - got length %v, expected length %v", len(val), 50)
	}
}

func TestSetOverwrite(t *testing.T) {
	store := setupStore()
	originalVal := store.keyValues["key1"]

	store.Set("key1", "overwritten")

	val := store.keyValues["key1"]
	if val == originalVal {
		t.Errorf("failed to overwrite key %s - old value still present", "key1")
	}
}

func setupStore() Secret {
	store := Secret{
		encryptionKey: "test",
		keyValues:     make(map[string]string),
	}
	store.keyValues["key1"] = "2113e8af2d434c2157539077a53000f90199a9c2f1d8"
	store.keyValues["key2"] = "86b365b215fb05760b6f7ab6eaa2670f9488967ebb1e"
	return store
}
