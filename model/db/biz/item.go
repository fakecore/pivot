package biz

import (
	"time"
)

type Item struct {
	CreatedAt    time.Time `json:"createdAt"`            // 创建时间
	UpdatedAt    time.Time `json:"updatedAt"`            // 更新时间
	Id           uint      `json:"id" gorm:"primarykey"` // 主键ID
	Name         string    `json:"name" gorm:"comment:商品名字"`
	BrandId      uint      `json:"brandId" gorm:"comment:品牌ID"`
	FirstTypeId  uint      `json:"firstTypeId" gorm:"comment:第一大类"`
	SecondTypeId uint      `json:"secondTypeId" gorm:"comment:第二大类"`
	ImgUrl       string    `json:"imgUrl" gorm:"comment:图片地址"`
	Price        float64   `json:"price" gorm:"comment:商品价格"`
}
