package utils

import (
	"fmt"
	"hmall/application/address/api/internal/model"
	"hmall/application/address/api/internal/types"
)

func AddressPO_to_AddresDTO(po *model.AddressPO) *types.QueryAddressesDTO {
	return &types.QueryAddressesDTO{
		Id:        po.Id,
		City:      po.City,
		IsDefault: po.IsDefault,
		Mobile:    po.Mobile,
		Notes:     po.Notes,
		Province:  po.Province,
		Street:    po.Street,
		Town:      po.Town,
		Contact:   po.Contact,
	}
}

func CacheIds(id int) string {
	return fmt.Sprintf("%s#%d", types.CacheAddressKey, id)
}
