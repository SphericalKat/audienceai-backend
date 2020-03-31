package entities

type Status struct {
	FilePath      string      `json:"file_path"`
	FileName      string      `json:"file_name" gorm:"primary_key"`
	Status        string      `json:"status"`
	NumFrames     int         `json:"num_frames"`
	EmotionScores [][]float64 `json:"emotion_scores"`
}
