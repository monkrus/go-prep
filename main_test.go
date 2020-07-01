package main

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestWrongType(t *testing.T) {
	result := nil
	expected := error
	if result != expected {
		t.Errorf("WrongType() test returned an unexpected result: got %v want %v", result, expected)
	}
}


func TestInit(t *testing.T) {
	result := error
	expected := nil
	if result != expected {
		t.Errorf("init() test returned an unexpected result: got %v want %v", result, expected)
	}
}

func TestHelloLink(t *testing.T) {
	result := error
	expected := nil
	if result != expected {
		t.Errorf("helloLink() test returned an unexpected result: got %v want %v", result, expected)
	}
}
func TestHttp(t *testing.T) {
	//
	  router := mux.NewRouter() 
	  {
      io.WriteString(w, "{ /"hello": /"expected response/"}")
	  }
  
	  router := httptest.NewRequest("GET", "http://localhost:8081/hello", nil)
	  w := httptest.NewRecorder()
	  router(w, req)
  
	  resp := w.Result()
	  body, _ := ioutil.ReadAll(resp.Body)
	  fmt.Println(resp.StatusCode)
	  fmt.Println(resp.Header.Get("Content-Type"))
	  fmt.Println(string(body))
  }

/*
We are testing the following functions"
-init
-helloLink
-wrongType
- also testing HTTP reponse request

*/
/*
OPTION 2

type AddResult struct {
x        int
y        int
expected int
}

var addResults =[]AddResult {
{1,2,3}, }

func TestAdd(t *testing.T) {
	for _, test :=range addResults {
		result :=Add(test.x, test.y)
		if (result !=test.expected){
			t.Fatal("Expected result not given")
		}
	}
}
*/

/* TEST SUITE

func TestExampleTestSuite(t *testing.T) {
suite.Run(t, new(ExampleTestSuite))
}

func TestWrongType(t *testing.T) {

	assert.Equal(t)

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object.Value)

	}

}
*/
/*
HTTP Testing

func TestHttp(t *testing.T) {
  //
    handler := func(w http.ResponseWriter, r *http.Request) {
    // here we write our expected response, in this case, we return a
    // JSON string which is typical when dealing with REST APIs
        io.WriteString(w, "{ \"status\": \"expected service response\"}")
    }

    req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
    w := httptest.NewRecorder()
    handler(w, req)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(resp.StatusCode)
    fmt.Println(resp.Header.Get("Content-Type"))
    fmt.Println(string(body))
}

*/
