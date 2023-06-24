package vocabulary

type WORD_TYPE int

type Word struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	LanguageID int    `json:"language_id"`
	Type       int    `json:"type"`
	CreatedAt  string `json:"created_at"`
	Language   Language
}

type Language struct {
	ID        int    `json:"id"`
	ISO       string `json:"iso"`
	CreatedAt string `json:"created_at"`
}

type Translation struct {
	ID            int `json:"id"`
	WordID        int `json:"word_id"`
	TranslationID int `json:"translation_id"`
}

const (
	Base WORD_TYPE = 0
	User WORD_TYPE = 1
)
