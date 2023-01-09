package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/model"
	"github.com/mrizalr/mini-project-evermos/utils"
)

type ProvinceHandler struct{}

func NewProvinceHandler(r fiber.Router) {
	handler := ProvinceHandler{}
	r.Get("/provcity/listprovincies", handler.GetProvinces)
	r.Get("/provcity/listcities/:prov_id", handler.GetCities)
	r.Get("/provcity/detailprovince/:prov_id", handler.GetProvince)
	r.Get("/provcity/detailcity/:city_id", handler.GetCity)
}

func (h *ProvinceHandler) GetProvinces(c *fiber.Ctx) error {
	return h.GetData(c, "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json", []model.Province{})
}

func (h *ProvinceHandler) GetCities(c *fiber.Ctx) error {
	prov_id := c.Params("prov_id")
	return h.GetData(c, fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", prov_id), []model.City{})
}

func (h *ProvinceHandler) GetProvince(c *fiber.Ctx) error {
	prov_id := c.Params("prov_id")
	return h.GetData(c, fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/province/%s.json", prov_id), model.Province{})
}

func (h *ProvinceHandler) GetCity(c *fiber.Ctx) error {
	city_id := c.Params("city_id")
	return h.GetData(c, fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regency/%s.json", city_id), model.City{})
}

func (h *ProvinceHandler) GetData(c *fiber.Ctx, url string, result interface{}) error {
	errs := []string{}

	err := utils.GetRegionData(url, &result)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return c.JSON(model.Response{
			Status:  true,
			Message: "Failed to GET data",
			Errors:  errs,
			Data:    nil,
		})
	}

	return c.JSON(model.Response{
		Status:  true,
		Message: "Succeed to GET data",
		Errors:  nil,
		Data:    result,
	})
}
