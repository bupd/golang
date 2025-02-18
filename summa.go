package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/go-containerregistry/pkg/crane"
)

// Result represents the structure of the data inside artifacts.json
type Result struct {
	Registry  string     `json:"registry"`
	Artifacts []Artifact `json:"artifacts"`
}

// Artifact represents an individual artifact
type Artifact struct {
	Repository string   `json:"repository"`
	Tag        []string `json:"tag"`
	Labels     []string `json:"labels"`
	Type       string   `json:"type"`
	Digest     string   `json:"digest"`
	Deleted    bool     `json:"deleted"`
}

func main() {
	// Image reference, change it to your image
	imgRef := "localhost:80/satellite/group1/state:latest"

	// Pull the image
	img, err := crane.Pull(imgRef)
	if err != nil {
		log.Fatalf("Failed to pull image: %v", err)
	}

	// Extract the artifacts.json from the image
	content, err := crane.Export(img)
	if err != nil {
		log.Fatalf("Failed to export image content: %v", err)
	}

	// Parse artifacts.json from the content
	var result Result
	err = json.Unmarshal(content, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal artifacts.json: %v", err)
	}

	// Output the result
	fmt.Printf("Registry: %s\n", result.Registry)
	for _, artifact := range result.Artifacts {
		fmt.Printf("Artifact: %s, Digest: %s\n", artifact.Repository, artifact.Digest)
	}
}
