package api

import "3-bin_manager/config"

type Api struct {
	cfg *config.Config
}

func NewApi(cfg *config.Config) *Api {
	return &Api{
		cfg: cfg,
	}
}
