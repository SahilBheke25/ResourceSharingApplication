package app

import (
	"database/sql"

	"github.com/SahilBheke25/quick-farm-backend/internal/app/equipment"
	"github.com/SahilBheke25/quick-farm-backend/internal/app/rental"
	"github.com/SahilBheke25/quick-farm-backend/internal/app/user"
	"github.com/SahilBheke25/quick-farm-backend/internal/repository"
)

type Dependencies struct {
	equipmentHandler equipment.Handler
	userHandler      user.Handler
	rentalHandler    rental.Handler
}

func InitializeDependencies(db *sql.DB) *Dependencies {

	equipmentRepo := repository.NewEquipmentStore(db)
	equipmentService := equipment.NewService(equipmentRepo)
	equipmentHandler := equipment.NewHandler(equipmentService)

	userRepo := repository.NewUserStorer(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	rentalRepo := repository.NewRentalStore(db)
	rentalService := rental.NewService(rentalRepo, equipmentService)
	rentalHandler := rental.NewHandler(rentalService)
	
	return &Dependencies{
		equipmentHandler: equipmentHandler,
		userHandler: userHandler,
		rentalHandler: rentalHandler,
	}
}
