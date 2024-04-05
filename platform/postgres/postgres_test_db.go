package postgres

/*
import (
	"attempt4/internal/domain/dto"
	"attempt4/platform/zap"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

var config dto.Config
var db *gorm.DB

// Todo: struct ile taşı
func init() {
	zap.Logger.Info("Env dosyaları okunuyor")

	//Todo: os.getwd ile al
	viper.AddConfigPath("../../..")
	viper.SetConfigType("env")
	viper.SetConfigName("test.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println(err)
	}

	db, err = InitializeDatabase(config.DBTestUrl)
	if err != nil {
		log.Println(err)
	}
}

func GetTestDb() *gorm.DB {
	return db.Begin()
}
*/
