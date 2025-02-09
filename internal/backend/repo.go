package backend

import "database/sql"

type BackendRepo interface {
	GetLogs() ([]Container, error)
	UpdateContainers(cts []Container) error
	UpdateContainer(ctr Container) error
}

type BackendDBRepository struct {
	dtb *sql.DB
}

func NewDBRepo(sdb *sql.DB) *BackendDBRepository {
	return &BackendDBRepository{dtb: sdb}
}
