package domain

import (
    "time"
)

//Market represents a model
type Market struct {
    ID string `json:"id"`
    CreatedAt time.Time `json:"dateCreated,omitempty"`
    UpdatedAt time.Time `json:"dateUpdated,omitempty"`
    Name string `json:"name"`
    Picture string `json:"picture"`
}