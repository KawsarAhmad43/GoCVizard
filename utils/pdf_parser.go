package utils

import (
    "bytes"
    "io/ioutil"
    "net/http"
)

// ExtractTextFromPDF extracts text from a given PDF file using UniPDF API.
func ExtractTextFromPDF(filePath, apiKey string) (string, error) {
    // Read the PDF file
    pdfData, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }

    // Create a POST request to the UniPDF API
    req, err := http.NewRequest("POST", "https://api.unidoc.io/v1/pdf/extract/text", bytes.NewBuffer(pdfData))
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/pdf")

    // Execute the request
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Read and return the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
