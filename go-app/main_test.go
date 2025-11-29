package main
import ("net/http"; "net/http/httptest"; "testing")
func TestHandler(t *testing.T) {
	tests := []struct { name, path, want string }{
		{"Root", "/", "Welcome to root web"},
		{"Sub", "/abc", "you access the abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.path, nil)
			rr := httptest.NewRecorder()
			handler(rr, req)
			if rr.Body.String() != tt.want { t.Errorf("got %v want %v", rr.Body.String(), tt.want) }
		})
	}
}