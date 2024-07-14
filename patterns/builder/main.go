package main

import (
	"fmt"
)

// HTTPRequest представляет HTTP-запрос с различными параметрами
type HTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

// HTTPRequestBuilder предоставляет методы для настройки параметров HTTP-запроса
type HTTPRequestBuilder struct {
	method  string
	url     string
	headers map[string]string
	body    string
}

// NewHTTPRequestBuilder создает новый экземпляр HTTPRequestBuilder
func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		headers: make(map[string]string),
	}
}

// SetMethod устанавливает метод HTTP-запроса
func (b *HTTPRequestBuilder) SetMethod(method string) *HTTPRequestBuilder {
	b.method = method
	return b
}

// SetURL устанавливает URL HTTP-запроса
func (b *HTTPRequestBuilder) SetURL(url string) *HTTPRequestBuilder {
	b.url = url
	return b
}

// AddHeader добавляет заголовок к HTTP-запросу
func (b *HTTPRequestBuilder) AddHeader(key, value string) *HTTPRequestBuilder {
	b.headers[key] = value
	return b
}

// SetBody устанавливает тело HTTP-запроса
func (b *HTTPRequestBuilder) SetBody(body string) *HTTPRequestBuilder {
	b.body = body
	return b
}

// Build создает объект HTTPRequest на основе настроек
func (b *HTTPRequestBuilder) Build() *HTTPRequest {
	return &HTTPRequest{
		Method:  b.method,
		URL:     b.url,
		Headers: b.headers,
		Body:    b.body,
	}
}

func main() {
	// Пример использования строителя для создания HTTP-запроса
	request := NewHTTPRequestBuilder().
		SetMethod("POST").
		SetURL("https://api.example.com/data").
		AddHeader("Content-Type", "application/json").
		SetBody(`{"key":"value"}`).
		Build()

	fmt.Printf("Method: %s\n", request.Method)
	fmt.Printf("URL: %s\n", request.URL)
	fmt.Printf("Headers: %s\n", request.Headers)
	fmt.Printf("Body: %s\n", request.Body)
}
