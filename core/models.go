package core

// ----- `Project` model
type Project struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Sections  []Section `gorm:"foreignKey:ProjectID; constraint:OnDelete:CASCADE;"`
	CreatedAt int64     `gorm:"autoUpdateTime:nano"`
	UpdatedAt int64     `gorm:"autoCreateTime:nano"`
}

// ----- `Section` model
type Section struct {
	ID        string `gorm:"primaryKey"`
	Heading   string `gorm:"not null"`
	ProjectID string `gorm:"not null"`
	Tasks     []Task `gorm:"foreignKey:SectionID; constraint:OnDelete:CASCADE;"`
	CreatedAt int64  `gorm:"autoUpdateTime:nano"`
	UpdatedAt int64  `gorm:"autoCreateTime:nano"`
}

// ----- `Task` model
type Task struct {
	ID          string `gorm:"primaryKey"`
	Heading     string `gorm:"not null"`
	Description string
	Priority    uint      `gorm:"not null; default:4"`
	Done        bool      `gorm:"not null; default:false"`
	SectionID   string    `gorm:"not null"`
	SubTasks    []SubTask `gorm:"foreignKey:TaskID; constraint:OnDelete:CASCADE;"`
	CreatedAt   int64     `gorm:"autoUpdateTime:nano"`
	UpdatedAt   int64     `gorm:"autoCreateTime:nano"`
}

// ----- `SubTask` model
type SubTask struct {
	ID          string `gorm:"primaryKey"`
	Heading     string `gorm:"not null"`
	Description string
	Priority    uint   `gorm:"not null; default:4"`
	Done        bool   `gorm:"not null; default:false"`
	TaskID      string `gorm:"not null"`
	CreatedAt   int64  `gorm:"autoUpdateTime:nano"`
	UpdatedAt   int64  `gorm:"autoCreateTime:nano"`
}
