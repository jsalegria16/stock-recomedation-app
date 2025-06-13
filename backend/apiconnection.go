package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// URL de la API a consultar
const apiUrl = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"

// Token de autenticación Bearer (¡no exponer en producción!)
const token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MzEsImVtYWlsIjoianNhbGVncmlhQHVuaWNhdWNhLmVkdS5jbyIsImV4cCI6MTc0OTc5MjYzMiwiaWQiOiIwIiwicGFzc3dvcmQiOiInIE9SICcxJyA9ICcxIn0.uzwLDf_c5150ZnvydQpl7fu5Rlhj1Ftl0Ve9cjrQHAY"

// fetchData realiza una petición GET a la API y retorna los datos como un map.
// Si page no es vacío, se envía como query param "next_page".
func fetchData(page string) (map[string]interface{}, error) {
	client := resty.New() // Crea un nuevo cliente HTTP

	// Prepara la petición con headers necesarios
	req := client.R().
		SetHeader("Authorization", token).
		SetHeader("Content-Type", "application/json")

	// Si se especifica una página, agrega el parámetro de paginación
	if page != "" {
		req.SetQueryParam("next_page", page)
	}

	// Realiza la petición GET
	resp, err := req.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	// Decodifica la respuesta JSON en un map
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	return result, err
}

func main() {
	// Llama a fetchData sin paginación (primera página)
	data, err := fetchData("")
	if err != nil {
		log.Fatalf("Error al obtener datos de la API: %v", err)
	}

	// Imprime el resultado como JSON formateado
	jsonPretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error al formatear el resultado: %v", err)
	}
	fmt.Println(string(jsonPretty))
}
