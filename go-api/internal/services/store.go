package services

import (
	"errors"
	"go-api/internal/database"
	"go-api/internal/dto"
	"go-api/internal/models"
	"go-api/pkg"
	"math"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoreService struct{}

func NewStoreService() *StoreService {
	return &StoreService{}
}

func (s *StoreService) GetAllStores(query dto.PaginationQueryDTO) (pkg.PaginatedResponse, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	if query.SortBy == "" {
		query.SortBy = "created_at"
	}
	if query.SortOrder == "" {
		query.SortOrder = "DESC"
	}
	db := database.DB.Model(&models.Store{}).Where("deleted_at IS NULL")

	// Count
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return pkg.PaginatedResponse{}, err
	}

	// Pagination
	offset := (query.Page - 1) * query.PageSize
	totalPages := int(math.Ceil(float64(totalCount) / float64(query.PageSize)))

	var stores []models.Store
	if err := database.DB.Order(query.SortBy + " " + query.SortOrder).Limit(query.PageSize).Offset(offset).Find(&stores).Error; err != nil {
		return pkg.PaginatedResponse{}, err
	}
	return pkg.PaginatedResponse{
		Data: stores,
		Pagination: pkg.PaginationMeta{
			CurrentPage:  query.Page,
			PageSize:     query.PageSize,
			TotalRecords: totalCount,
			TotalPages:   totalPages,
		},
	}, nil
}

func (s *StoreService) GetStoreByID(id string) (models.Store, error) {
	var store models.Store
	if err := database.DB.Where("id = ? AND deleted_at IS NULL", id).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Store{}, errors.New("store not found")
		}
		return models.Store{}, err
	}
	return store, nil
}

func (s *StoreService) CreateStore(storeDTO dto.CreateStoreDTO) (models.Store, error) {
	store := models.Store{
		Name: storeDTO.Name,
	}

	if err := database.DB.Create(&store).Error; err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *StoreService) UpdateStore(id string, storeDTO dto.UpdateStoreDTO) (models.Store, error) {
	storeID, err := uuid.Parse(id)
	if err != nil {
		return models.Store{}, errors.New("invalid store ID")
	}

	var store models.Store
	if err := database.DB.Where("id = ? AND deleted_at IS NULL", storeID).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Store{}, errors.New("store not found")
		}
		return models.Store{}, err
	}

	updates := make(map[string]interface{})

	if storeDTO.Name != nil {
		updates["name"] = *storeDTO.Name
	}

	updates["updated_at"] = time.Now()

	if err := database.DB.Model(&store).Updates(updates).Error; err != nil {
		return models.Store{}, err
	}

	if err := database.DB.Where("id = ?", storeID).First(&store).Error; err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *StoreService) DeleteStore(id string) error {
	storeID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid store ID")
	}

	var store models.Store
	if err := database.DB.Where("id = ? AND deleted_at IS NULL", storeID).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("store not found")
		}
		return err
	}

	if err := database.DB.Model(&store).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
