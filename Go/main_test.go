package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetBooks(t *testing.T) {
	expected := `{
    "error": "not authorized",
    "message": "not authorized"
}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
	req.Header.Add("Authorization", "initestyasir")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]
	fmt.Println(val)
	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "text/plain; charset=utf-8" {
		t.Fatalf("Expected \"text/plain; charset=utf-8\", got %s", val[0])
	}
}

func TestCreateBooks(t *testing.T) {
	expected := `{
    "error": "not authorized",
    "message": "not authorized"
}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	defer ts.Close()

	data := url.Values{}
	data.Set("title", "the great best")
	data.Set("author", "nama")
	data.Set("isbn", "ISBN-12312")
	data.Set("publish_date", "2024-04-04")

	client := &http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/books", ts.URL), strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "adminadmin")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
func TestGetSingleBooks(t *testing.T) {
	expected := `{
    "error": "not authorized",
    "message": "not authorized"
}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/books/1", ts.URL), nil)
	req.Header.Add("Authorization", "adminadmin")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]
	fmt.Println(val)
	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "text/plain; charset=utf-8" {
		t.Fatalf("Expected \"text/plain; charset=utf-8\", got %s", val[0])
	}
}
func TestUpdateBooks(t *testing.T) {
	expected := `{
    "error": "not authorized",
    "message": "not authorized"
}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	defer ts.Close()

	data := url.Values{}
	data.Set("title", "the great best")
	data.Set("author", "nama")
	data.Set("isbn", "ISBN-12312")
	data.Set("publish_date", "2024-04-04")

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/books/1", ts.URL), strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "adminadmin")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}

func TestDeleteBooks(t *testing.T) {
	expected := `{
    "error": "not authorized",
    "message": "not authorized"
}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/books/1", ts.URL), nil)
	req.Header.Add("Authorization", "adminadmin")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
