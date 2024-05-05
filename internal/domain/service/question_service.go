package service

import (
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type QuestionService struct {
	userRepository     repository.UserRepository
	questionRepository repository.QuestionRepository
	pointRepository    repository.PointRepository
	productRepository  repository.ProductRepository
	tripRepository     repository.TripRepository
}

func NewQuestionService(
	userRepository repository.UserRepository,
	questionRepository repository.QuestionRepository,
	pointRepository repository.PointRepository,
	productRepository repository.ProductRepository,
	tripRepository repository.TripRepository) QuestionService {

	q := QuestionService{
		userRepository,
		questionRepository,
		pointRepository,
		productRepository,
		tripRepository,
	}
	return q
}
func (q *QuestionService) GetQuestions() ([]dto.QuestionDto, error) {
	var questionsDto []dto.QuestionDto
	var questions []entity.Question
	var totalNumber int64
	var err error

	questions, totalNumber, err = q.questionRepository.GetAllProducts()
	if totalNumber == 0 {
		return questionsDto, errors.New("no questions found")
	}
	if err != nil {
		zap.Logger.Error(err)
		return questionsDto, err
	}

	for i, _ := range questions {
		questionsDto = append(questionsDto,
			dto.QuestionDto{
				Id:          questions[i].Id,
				Description: questions[i].Description,
				Topic:       questions[i].Topic,
			})
	}

	return questionsDto, nil
}
func (q *QuestionService) CalculatePoints(id int32, questions []dto.TopicPointsDto) error {
	var point []entity.TopicPoints
	tempPoints, err := q.pointRepository.GetById(id)
	if tempPoints.Id != 0 {
		return errors.New("Önceden çözülmüş")
	}
	if err != nil {
		return err
	}

	for i, _ := range questions {
		point = append(point,
			entity.TopicPoints{
				Id:         0,
				QuestionId: questions[i].QuestionId,
				UserId:     id,
				Points:     questions[i].Points,
				Status:     enum.UserActiveStatus,
			})
	}
	for i, _ := range point {
		e, err := q.pointRepository.Create(point[i])
		if e.Id == 0 || err != nil {
			zap.Logger.Error(err)
			return err
		}
	}
	return nil
}
func (q *QuestionService) GetAIRecommendation(id int32) (dto.ProductTripDto, error) {
	var destinationDto []dto.TripDto
	var result dto.ProductTripDto
	userPoints, total, err := q.pointRepository.GetAllByUserIDProducts(id)
	if total == 0 {
		zap.Logger.Error(err)
		return result, err
	}
	if err != nil {
		zap.Logger.Error(err)
		return result, err
	}

	trips, total, err := q.tripRepository.GetAllTrips()
	if total == 0 {
		zap.Logger.Error(err)
		return result, err
	}
	if err != nil {
		zap.Logger.Error(err)
		return result, err
	}

	destination, err := MatchTrip(userPoints, trips)
	if err != nil {
		zap.Logger.Error(err)
		return result, err
	}
	var idStrings []string
	for i, _ := range destination {
		idStrings = append(idStrings, strconv.FormatInt(int64(destination[i].Id), 10))
		destinationDto = append(destinationDto, dto.TripDto{
			Id:          destination[i].Id,
			Description: destination[i].Description,
			Name:        destination[i].Name,
			Status:      destination[i].Status,
			Lat:         destination[i].Lat,
			Lng:         destination[i].Lng,
		})
	}
	idListString := strings.Join(idStrings, "-")

	name := "AIRecommendation" + generateUniqueString()
	product := entity.Product{
		Id:        0,
		Name:      name,
		Quantity:  1,
		Price:     1000,
		Status:    enum.ProductAvailable,
		UserId:    id,
		Trip:      idListString,
		CreatedAt: time.Now(),
	}
	pro, err := q.productRepository.Create(product)
	if pro.Id == 0 {
		zap.Logger.Error(err)
		return result, err
	}
	if err != nil {
		zap.Logger.Error(err)
		return result, err
	}

	idStringSlice := strings.Split(pro.Trip, "-")

	for _, idStr := range idStringSlice {
		tripID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Error converting ID string to int:", err)
			return result, err
		}
		tripByParsedId, err := q.tripRepository.GetById(int32(tripID))
		if err != nil {
			zap.Logger.Error(err)
			return result, err
		}

		tempTripDto := dto.TripDto{
			Id:          tripByParsedId.Id,
			Description: tripByParsedId.Description,
			Name:        tripByParsedId.Name,
			Status:      tripByParsedId.Status,
			Lat:         tripByParsedId.Lat,
			Lng:         tripByParsedId.Lng,
		}
		result.Trip = append(result.Trip, tempTripDto)
	}
	result.Name = pro.Name
	result.Price = pro.Price
	result.Quantity = pro.Quantity
	return result, nil
}
func MatchTrip(userPoints []entity.TopicPoints, trips []entity.Trip) ([]entity.Trip, error) {
	var destination []entity.Trip

	minDistance := math.Inf(1)
	perfectTripIndex := -1

	matchRates := make(map[int]float64)
	var NaturePoint float64
	var HistoricalPoint float64
	var CulturalPoint float64
	var RelaxPoint float64
	var AdventurePoint float64

	for i, trip := range trips {
		NaturePoint = float64(trip.NaturePoint - (userPoints[0].Points + userPoints[1].Points))
		HistoricalPoint = float64(trip.NaturePoint - (userPoints[2].Points + userPoints[3].Points))
		CulturalPoint = float64(trip.NaturePoint - (userPoints[4].Points + userPoints[5].Points))
		RelaxPoint = float64(trip.NaturePoint - (userPoints[6].Points + userPoints[7].Points))
		AdventurePoint = float64(trip.NaturePoint - (userPoints[8].Points + userPoints[9].Points))
		matchRate := math.Sqrt(
			math.Pow(NaturePoint, 2) +
				math.Pow(HistoricalPoint, 2) +
				math.Pow(CulturalPoint, 2) +
				math.Pow(RelaxPoint, 2) +
				math.Pow(AdventurePoint, 2))
		matchRates[i] = matchRate
	}
	perfectTripIndex = 0
	for i, matchRate := range matchRates {
		if matchRates[perfectTripIndex] > matchRate {
			perfectTripIndex = i
		}
	}

	if perfectTripIndex == -1 {
		return destination, errors.New("trip not found")
	}

	destination = append(destination, trips[perfectTripIndex])

	for i, trip := range trips {
		dist := Distance(trips[i].Lat, trips[i].Lng, destination[0].Lat, destination[0].Lng)
		if dist <= 100 && i != perfectTripIndex && matchRates[i] > 1 {
			destination = append(destination, trip)
			if dist < minDistance {
				minDistance = dist
			}
		}
	}

	sort.Slice(destination, func(i, j int) bool {
		return matchRates[i] < matchRates[j]
	})

	return destination, nil
}
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	const radius = 6371 // Earth radius in kilometers
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)
	lat1 = lat1 * (math.Pi / 180)
	lat2 = lat2 * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return radius * c
}
func generateUniqueString() string {
	timestamp := time.Now().Unix()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000)

	uniqueString := fmt.Sprintf("%d%d", timestamp%1000000, randomNumber)

	return uniqueString
}
