package models

import "time"

type Product struct {
	Id         int32     `json:"id"`
	Name       string    `json:"name"`
	Value      float64   `json:"value"`
	CreateDate time.Time `json:"createDate"`
}