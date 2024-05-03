package service

import (
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	"errors"
	"math"
	"sort"
	"strconv"
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
	for i, _ := range questions {
		point = append(point,
			entity.TopicPoints{
				Id:         0,
				QuestionId: id,
				UserId:     questions[i].UserId,
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

func (q *QuestionService) GetAIRecommendation(id int32) ([]dto.TripDto, error) {
	var destinationDto []dto.TripDto

	var tripDto []dto.TripDto
	userPoints, total, err := q.pointRepository.GetAllByUserIDProducts(id)
	if total == 0 {
		zap.Logger.Error(err)
		return tripDto, err
	}
	if err != nil {
		zap.Logger.Error(err)
		return tripDto, err
	}

	trips, total, err := q.tripRepository.GetAllTrips()
	if total == 0 {
		zap.Logger.Error(err)
		return tripDto, err
	}
	if err != nil {
		zap.Logger.Error(err)
		return tripDto, err
	}

	destination, err := MatchTrip(userPoints, trips)
	if err != nil {
		zap.Logger.Error(err)
		return tripDto, err
	}
	idList := ""
	for i, _ := range destination {
		idList += strconv.Itoa(int(destination[i].Id)) + "-"
		destinationDto = append(destinationDto, dto.TripDto{
			Id:          destination[i].Id,
			Description: destination[i].Description,
			Name:        destination[i].Name,
			Status:      destination[i].Status,
			Lat:         destination[i].Lat,
			Lng:         destination[i].Lng,
		})
	}
	product := entity.Product{
		Id:        0,
		Name:      "AI Recommendation",
		Quantity:  1,
		Price:     1000,
		Status:    enum.ProductAvailable,
		UserId:    id,
		Trip:      idList,
		CreatedAt: time.Now(),
	}
	pro, err := q.productRepository.Create(product)
	if pro.Id == 0 {
		zap.Logger.Error(err)
		return tripDto, err
	}
	if err != nil {
		zap.Logger.Error(err)
		return tripDto, err
	}
	return destinationDto, nil
}
func MatchTrip(userPoints []entity.TopicPoints, trips []entity.Trip) ([]entity.Trip, error) {
	var destination []entity.Trip

	minDistance := math.Inf(1)
	perfectTripIndex := -1

	var distances []float64
	matchRates := make(map[int]float64)
	for i, trip := range trips {
		matchRate := math.Sqrt(math.Pow(float64(trip.NaturePoint-userPoints[1].Points), 2) +
			math.Pow(float64(trip.HistoricalPoint-userPoints[2].Points), 2) +
			math.Pow(float64(trip.CulturalPoint-userPoints[4].Points), 2) +
			math.Pow(float64(trip.RelaxPoint-userPoints[5].Points), 2) +
			math.Pow(float64(trip.AdventurePoint-userPoints[3].Points), 2))
		matchRates[i] = matchRate

		perfectTripIndex = i

	}

	if perfectTripIndex == -1 {
		return destination, errors.New("trip not found")
	}

	destination = append(destination, trips[perfectTripIndex])

	for i, trip := range trips {
		if distances[i] <= 100 && i != perfectTripIndex {
			dist := Distance(trips[i].Lat, trips[i].Lng, trip.Lat, trip.Lng)
			distances = append(distances, dist)
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
