package repository

import (
	_entities "github.com/MarkTBSS/080_Pagination/entities"
	_models "github.com/MarkTBSS/080_Pagination/pkg/itemShop/models"
)

type ItemShopRepository interface {
	Listing(itemFilter *_models.ItemFilter) ([]*_entities.Item, error)
	Counting(itemFilterDto *_models.ItemFilter) (int64, error)
}
