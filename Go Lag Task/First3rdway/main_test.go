package main

import (
	"net/http"
	"testing"
)

func TestWelcome_name_JSON(t *testing.T) {
	resp, err := http.Get("http://localhost:1000/persons")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestWelcome_name_JSON2(t *testing.T) {
	resp, err := http.Get("http://localhost:1000/health-check")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

// func TestCreate(t *testing.T) {
// 	person := &Person{
// 		Id:        "1",
// 		Firstname: "Issac",
// 		Lastname:  "N"
// 	}
// 	jsonPerson, _ := json.Marshal(person)
// 	request, _ := http.NewRequest("Get", "/persons", bytes.NewBuffer(jsonPerson))
// 	response := httptest.NewRecorder()
// 	Router().ServeHTTP(response, request)
// 	assert.Equal(t, 200, response.Code, "OK response is expected")
// }

// func testPersons(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/person?Name=Thomas", nil)
// 	w := httptest.NewRecorder()
// 	Persons(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 	}
// 	if string(data) != "" {
// 		return
// 	}
// }
