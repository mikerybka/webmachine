package tests

import (
	"fmt"
	"io"
	"net/http"
)

const endpoint = "http://localhost:3000"

func GetIndexHTML() error {
	url := endpoint + "/example.org"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "hello"
	if string(b) != expected {
		return fmt.Errorf("tests.GetIndexHTML(): expected %q, got %q", expected, string(b))
	}
	return nil
}

func GetIndexJSON() error {
	url := endpoint + "/example.org/index.json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "{}"
	if string(b) != expected {
		return fmt.Errorf("tests.GetIndexJSON(): expected %q, got %q", expected, string(b))
	}
	return nil
}

func Get() error {
	url := endpoint + "/example.org/go"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "hello\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}

func Post() error {
	url := endpoint + "/example.org/go"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "hello\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}

func Put() error {
	url := endpoint + "/example.org/go"
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "hello\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}

func Delete() error {
	url := endpoint + "/example.org/go"
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "hello\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}

func Patch() error {
	url := endpoint + "/example.org/go"
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "hello\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}

func PathParams() error {
	url := endpoint + "/example.org/go/test"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "test\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}
func QueryParams() error {
	url := endpoint + "/example.org/go/query?q=123&b=true"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "q: 123\nb: true\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}
func RequestHeaders() error {
	return nil
}
func WildcardPathParam() error {
	url := endpoint + "/example.org/go/test/123/456/789"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	expected := "test\n123/456/789\n"
	if string(b) != expected {
		return fmt.Errorf("tests.Get(): expected %q, got %q", expected, string(b))
	}
	return nil
}
func RequestBody() error {
	return nil
}
func ResponseBody() error {
	return nil
}
func ResponseStatusCode() error {
	return nil
}
func ResponseHeaders() error {
	return nil
}
