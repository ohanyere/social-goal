package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type envelope map[string]any

const maxBodyBytes int64 = 1_048_576 // 1MB

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	return enc.Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	_ = writeJSON(w, status, envelope{
		"error": message,
	})
}

func readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode the first JSON value
	if err := dec.Decode(dst); err != nil {
		return err
	}

	// Ensure there is ONLY one JSON value in the body
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must contain only one JSON value")
	}

	return nil
}