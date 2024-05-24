package main

import (
	"encoding/json"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"go-exercise-vvila/internal/api"
	"go-exercise-vvila/internal/cache"
	"go-exercise-vvila/internal/kraken"
	"go-exercise-vvila/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockKrakenServer() *httptest.Server {
	expected := "{\"error\":[],\"result\":{\"XXBTZUSD\":{\"a\":[\"67501.50000\",\"1\",\"1.000\"],\"b\":[\"67501.40000\",\"1\",\"1.000\"],\"c\":[\"67500.20000\",\"0.00021525\"],\"v\":[\"3334.02749786\",\"3473.20902081\"],\"p\":[\"68626.01177\",\"68650.47747\"],\"t\":[32203,34397],\"l\":[\"66250.00000\",\"66250.00000\"],\"h\":[\"70500.00000\",\"70500.00000\"],\"o\":\"69144.10000\"}}}"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	return svr
}

func mockKrakenService(serverURL string) *service.KrakenService {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()
	redisClientMocked := cache.NewCacheClient(mr.Addr())
	krakenClient := kraken.KrakenClient{BaseURL: serverURL}
	return service.NewKrakenService(&krakenClient, redisClientMocked)
}

func TestGetLTPHandler(t *testing.T) {
	krakenServer := mockKrakenServer()
	defer krakenServer.Close()
	krakenMockURL := fmt.Sprintf("%s?q=", krakenServer.URL)
	krakenMockURL = fmt.Sprintf("%s%s", krakenMockURL, "%s")
	krakenService := mockKrakenService(krakenMockURL)
	handler := api.GetLTPHandler(krakenService)

	rr := httptest.NewRecorder()
	url := fmt.Sprintf("%s/api/v1/ltp", krakenServer.URL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if ltp, ok := response["ltp"].([]interface{}); ok {
		for _, item := range ltp {
			if ltpMap, ok := item.(map[string]interface{}); ok {
				if ltpMap["amount"] != "67500.20000" || ltpMap["pair"] != "XXBTZUSD" {
					t.Errorf("Response does not contain expected field 'ltp'")
				}
			}
		}
	} else {
		t.Errorf("Response does not contain expected field 'ltp'")
	}

}
