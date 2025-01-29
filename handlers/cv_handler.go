package handlers

import (
	"fmt"
	"go-cv-matcher/models"
	"go-cv-matcher/utils"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ShowUploadForm renders the upload form
func ShowUploadForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/upload.html"))
	tmpl.Execute(w, nil)
}

// UploadCV handles the CV upload and skill matching
func UploadCV(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("cv")
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Save the file
		dir := "uploads"
		os.MkdirAll(dir, os.ModePerm)
		filePath := filepath.Join(dir, header.Filename)
		outFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = outFile.ReadFrom(file)
		if err != nil {
			http.Error(w, "Error writing to file", http.StatusInternalServerError)
			return
		}

		// Extract text from the uploaded PDF
		text, err := utils.ExtractTextFromPDF(filePath, "601d1f9feb925bbbe89e7b7d4605f963649b99fc5c7346398331b59c89a71a5a")
		if err != nil {
			http.Error(w, "Error extracting text from PDF", http.StatusInternalServerError)
			return
		}

		// Example position skills (replace this with dynamic data if needed)
		position := struct {
			Skills []string
		}{
			Skills: []string{"php", "javascript", "laravel"},
		}

		// Check for matched skills in the extracted text
		matchedSkills := []string{}
		for _, skill := range position.Skills {
			if strings.Contains(strings.ToLower(text), strings.ToLower(skill)) {
				matchedSkills = append(matchedSkills, skill)
			}
		}

		// Calculate score
		score := 0
		if len(position.Skills) > 0 {
			score = len(matchedSkills) * 10 / len(position.Skills)
		}

		// Create a candidate object
		candidate := models.Candidate{
			Name:    header.Filename,
			Skills:  position.Skills,
			Matched: matchedSkills,
			Score:   score,
		}

		// Respond with candidate details
		fmt.Fprintf(w, "Candidate: %s\nMatched Skills: %v\nScore: %d/10", candidate.Name, candidate.Matched, candidate.Score)
	}
}
