package backend

import "database/sql"

type BackendRepo interface {
	GetLogs() ([]Container, error)
	GetList() ([]string, error)
	GetContainer(ipv4 string) (*ContainerStat, error)
	UpdateContainer(ctr Container) error
}

type BackendDBRepository struct {
	dtb *sql.DB
}

func NewDBRepo(sdb *sql.DB) *BackendDBRepository {
	return &BackendDBRepository{dtb: sdb}
}
