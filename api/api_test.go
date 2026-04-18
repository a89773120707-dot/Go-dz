package api

import (
	"3-bin_manager/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestClient(serverUrl string) *Client {
	return &Client{
		httpClient: &http.Client{},
		cfg:        &config.Config{Key: "api-key"},
		baseUrl:    serverUrl,
	}
}

func TestClient_CreateBin_Success(t *testing.T) {
	//test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("ожидался метод POST, получен %s", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"metadata": {"id":"test-bin-id-123"}}`))
	}))
	defer server.Close()

	//newclient
	client := newTestClient(server.URL)

	//data map[string]{interface}
	data := map[string]interface{}{"key": "val"}

	binID, err := client.CreateBin(data)
	if err != nil {
		t.Fatalf("ожидалась nil ошибка, получена %v", err)
	}

	if binID != "test-bin-id-123" {
		t.Errorf("ожидался ID 'test-bin-id-123', получен %s", binID)
	}

}

func TestGetBin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("ожидался метод GET, получен %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"record":{"login": "user", "password": "secret"}}`))
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	record, err := client.GetBin("test-bin-id")

	if err != nil {
		t.Fatalf("ожидалась nil ошибка, получена %v", err)
	}
	if record["login"] != "user" {
		t.Errorf("ожидался login='user', получен %v", record["login"])
	}

}

func TestUpdateBin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("ожидался метод PUT, получен %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"metadata":{"id":"update-bin-id"}}`))
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	data := map[string]interface{}{"key": "val"}

	binID, err := client.UpdateBin("test-binID", data)

	if err != nil {
		t.Fatalf("ожидалась nil ошибка, получена %v", err)
	}

	if binID != "update-bin-id" {
		t.Errorf("ожидался ID 'updated-bin-id', получен %s", binID)
	}
}

func TestDeletBin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("ожидался метод DELETE, получен %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	clietn := newTestClient(server.URL)

	err := clietn.DeleteBin("test-delete-bin")
	if err != nil {
		t.Fatalf("ожидалась nil ошибка, получена %v", err)
	}
}
