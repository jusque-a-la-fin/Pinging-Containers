package backend

import (
	"database/sql"
	"fmt"
)

func (repo *BackendDBRepository) UpdateContainer(ctr Container) error {
	err := AddContainer(repo.dtb, ctr)
	if err != nil {
		return fmt.Errorf("error returned from function `AddContainer`, package `backend`: %v", err)
	}
	return nil
}

func UpdatePingSuccessTime(dtb *sql.DB, ctr Container) error {
	query := `
		     UPDATE containers
		     SET success_ping_time = $1
		     WHERE ipv4 = $2;`

	_, err := dtb.Exec(query, ctr.SuccessPingTime, ctr.IPv4)
	if err != nil {
		return fmt.Errorf("error while updating time of a successful ping: %v", err)
	}

	err = AddPingTime(dtb, ctr)
	if err != nil {
		return fmt.Errorf("error returned from function `AddPingTime`, package `backend`: %v", err)
	}
	return nil
}

func AddContainer(dtb *sql.DB, ctr Container) error {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM containers WHERE ipv4 = $1)`
	err := dtb.QueryRow(query, ctr.IPv4).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error while checking if a container exists: %v", err)
	}

	if exists && ctr.IsSuccess {
		UpdatePingSuccessTime(dtb, ctr)
	} else {
		query = `INSERT INTO containers (ipv4, success_ping_time) VALUES ($1, $2);`
		_, err = dtb.Exec(query, ctr.IPv4, ctr.SuccessPingTime)
		if err != nil {
			return fmt.Errorf("error while adding container info: %v", err)
		}
	}
	AddPingTime(dtb, ctr)
	return nil
}

func AddPingTime(dtb *sql.DB, ctr Container) error {
	query := `INSERT INTO ping_results (ping_time, container_ipv4) VALUES ($1, $2);`
	_, err := dtb.Exec(query, ctr.PingDuration, ctr.IPv4)
	if err != nil {
		return fmt.Errorf("error while adding time of the ping: %v", err)
	}
	return nil
}
