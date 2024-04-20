package service

import _models "github.com/MarkTBSS/080_Pagination/pkg/itemShop/models"

type ItemShopService interface {
	Listing(itemFilter *_models.ItemFilter) (*_models.ItemResult, error)
}
