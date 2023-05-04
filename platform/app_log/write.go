package app_log

import (
	"attempt4/internal"
	"attempt4/platform/app_log/seed"
	"attempt4/platform/zap"
	"gorm.io/gorm"
	"time"
)

type ApplicationLog struct {
	Id           int32     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserId       int32     `gorm:"column:user_id"`
	LogType      string    `gorm:"column:type"`
	Content      string    `gorm:"column:content"`
	RelatedTable string    `gorm:"column:related_table"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}
type ApplicationLogDto struct {
	UserId       int32
	LogType      string
	Content      string
	RelatedTable string
	CreatedAt    time.Time
}
type ApplicationLogRepository struct {
	db *gorm.DB
}

func NewApplicationLogRepository(db *gorm.DB) ApplicationLogRepository {
	a := ApplicationLogRepository{db}
	return a
}

func (a *ApplicationLogRepository) Create(log ApplicationLog) (ApplicationLog, error) {
	if err := a.db.Create(&log).Error; err != nil {
		return log, internal.DBNotCreated
	}
	return log, nil
}

func (a *ApplicationLogRepository) GetById(id int32) (ApplicationLog, error) {
	var log ApplicationLog
	if err := a.db.Model(&log).Where("id=?", id).Scan(&log).Error; err != nil {
		return log, internal.DBNotFound
	}
	return log, nil
}

func (a *ApplicationLogRepository) GetByUserId(id int32) (ApplicationLog, error) {
	var log ApplicationLog
	if err := a.db.Model(&log).Where("user_id=?", id).Scan(&log).Error; err != nil {
		return log, internal.DBNotFound
	}
	return log, nil
}

func InitializeAppLogDatabase(db *gorm.DB) error {

	err := db.AutoMigrate(&ApplicationLog{})
	if err != nil {
		return err
	}
	seed.ApplicationLogSeed(db)
	return nil
}

type ApplicationLogService struct {
	ApplicationLogRepository ApplicationLogRepository
}

func NewApplicationLogService(applicationLogRepository ApplicationLogRepository) ApplicationLogService {
	a := ApplicationLogService{
		applicationLogRepository,
	}
	return a
}

func (a *ApplicationLogService) AddLog(content ApplicationLogDto) {
	log, err := a.ApplicationLogRepository.Create(
		ApplicationLog{
			Id:           0,
			UserId:       content.UserId,
			LogType:      content.LogType,
			Content:      content.Content,
			RelatedTable: content.RelatedTable,
			CreatedAt:    content.CreatedAt,
		})
	if err != nil {
		zap.Logger.Warn(content.Content)
	}
	if log.Id == 0 {
		zap.Logger.Warn(content.Content)
	}

}
