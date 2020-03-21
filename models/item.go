package models

// ItemMeta holds different parameters which can be useful but not directly
// related to the actual content in the item
type ItemMeta struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt,omitempty"`
}

// Item represents what is stored under a given key
type Item struct {
	Meta    ItemMeta `json:"meta"`
	Content []byte   `json:"content"`
}
