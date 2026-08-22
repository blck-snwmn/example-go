package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strconv"
	"testing"
	"testing/synctest"
	"time"
	"uuid"
)

func TestSeqFold(t *testing.T) {
	got := Values(1, 2, 3, 4).Fold(0, func(sum, value int) int {
		return sum + value
	})

	if want := 10; got != want {
		t.Errorf("Fold() = %d, want %d", got, want)
	}
}

func TestSeqMap(t *testing.T) {
	got := Values(1, 2, 3).Map(strconv.Itoa).Collect()
	want := []string{"1", "2", "3"}

	if !slices.Equal(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestSeqFilter(t *testing.T) {
	got := Values(1, 2, 3, 4).Filter(func(value int) bool {
		return value%2 == 0
	}).Collect()
	want := []int{2, 4}

	if !slices.Equal(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}

func TestSeqFlatMap(t *testing.T) {
	got := Values(1, 2, 3).FlatMap(func(value int) Seq[int] {
		return Values(value, value*10)
	}).Collect()
	want := []int{1, 10, 2, 20, 3, 30}

	if !slices.Equal(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestSeqIsLazy(t *testing.T) {
	converted := 0
	seq := Values(1, 2, 3).Map(func(value int) int {
		converted++
		return value * 2
	})

	if converted != 0 {
		t.Fatalf("Map() processed %d values before the sequence was consumed", converted)
	}

	seq.Collect()
	if converted != 3 {
		t.Errorf("Map() processed %d values, want 3", converted)
	}
}

func TestJSONV2RoundTrip(t *testing.T) {
	want := Person{Name: "Alice"}
	encoded, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	var got Person
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got != want {
		t.Errorf("Unmarshal() = %+v, want %+v", got, want)
	}
}

func TestJSONV2RejectsDuplicateNamesByDefault(t *testing.T) {
	var person Person
	err := json.Unmarshal([]byte(`{"name":"Alice","name":"Bob"}`), &person)
	if err == nil {
		t.Fatal("Unmarshal() succeeded with duplicate object names")
	}
}

func TestJSONV2AllowsDuplicateNamesWithOption(t *testing.T) {
	var person Person
	err := json.Unmarshal(
		[]byte(`{"name":"Alice","name":"Bob"}`),
		&person,
		jsontext.AllowDuplicateNames(true),
	)
	if err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got, want := person.Name, "Bob"; got != want {
		t.Errorf("Name = %q, want %q", got, want)
	}
}

func TestJSONV2RejectsInvalidUTF8(t *testing.T) {
	input := []byte{'"', 0xff, '"'}
	var value string
	if err := json.Unmarshal(input, &value); err == nil {
		t.Fatal("Unmarshal() succeeded with invalid UTF-8")
	}
}

func TestJSONV2RejectsUnknownMembersWithOption(t *testing.T) {
	var person Person
	err := json.Unmarshal(
		[]byte(`{"name":"Alice","age":27}`),
		&person,
		json.RejectUnknownMembers(true),
	)
	if err == nil {
		t.Fatal("Unmarshal() succeeded with an unknown object member")
	}
}

func TestJSONV2FieldNamesAreCaseSensitiveByDefault(t *testing.T) {
	input := []byte(`{"NAME":"Alice"}`)

	var caseSensitive Person
	if err := json.Unmarshal(input, &caseSensitive); err != nil {
		t.Fatalf("case-sensitive Unmarshal() error = %v", err)
	}
	if caseSensitive.Name != "" {
		t.Errorf("case-sensitive Name = %q, want empty", caseSensitive.Name)
	}

	var caseInsensitive Person
	if err := json.Unmarshal(input, &caseInsensitive, json.MatchCaseInsensitiveNames(true)); err != nil {
		t.Fatalf("case-insensitive Unmarshal() error = %v", err)
	}
	if got, want := caseInsensitive.Name, "Alice"; got != want {
		t.Errorf("case-insensitive Name = %q, want %q", got, want)
	}
}

func TestUUIDV7RoundTrip(t *testing.T) {
	id := uuid.NewV7()
	got, err := uuid.Parse(id.String())
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if got != id {
		t.Errorf("Parse() = %s, want %s", got, id)
	}
}

func TestNewTestServerWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const delay = time.Hour
		// The in-memory network lets synctest advance its fake clock while
		// the client and server are waiting on each other.
		server := httptest.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(delay)
			if _, err := fmt.Fprint(w, "done"); err != nil {
				t.Errorf("Fprint() error = %v", err)
			}
		}))

		started := time.Now()
		response, err := server.Client().Get("http://example.com/")
		if err != nil {
			t.Fatalf("Get() error = %v", err)
		}
		defer func() {
			if err := response.Body.Close(); err != nil {
				t.Errorf("response body Close() error = %v", err)
			}
		}()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatalf("ReadAll() error = %v", err)
		}
		if got, want := string(body), "done"; got != want {
			t.Errorf("response body = %q, want %q", got, want)
		}
		if got := time.Since(started); got != delay {
			t.Errorf("elapsed time = %v, want %v", got, delay)
		}
	})
}
