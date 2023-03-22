package polygon

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/pkg/network/polygon/rpc"
	"log"
)

const NetworkName = "polygon"

// PlateausValidationService implementation of PlateausValidation
type PlateausValidationService struct {
	rpc rpc.PlateausValidation
}

func NewPlateausValidationService(r rpc.PlateausValidation) *PlateausValidationService {
	return &PlateausValidationService{
		rpc: r,
	}
}

func (s PlateausValidationService) GetBalance() (int64, error) {
	bal, err := s.rpc.BalanceOf()

	if err != nil {
		log.Printf("could not rpc.BalanceOf: %s", err)
		return 0, err
	}

	return bal, nil
}
