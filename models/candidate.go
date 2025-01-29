package models

type Candidate struct {
    Name       string
    Skills     []string
    Matched    []string
    Score      int
    CVFilePath string
}
