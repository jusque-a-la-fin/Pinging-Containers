package backend

import (
	"fmt"
)

func (repo *BackendDBRepository) GetContainer(ipv4 string) (*ContainerStat, error) {
	query := `SELECT ping_time FROM ping_results WHERE container_ipv4 = $1;`
	rows, err := repo.dtb.Query(query, &ipv4)
	if err != nil {
		return nil, fmt.Errorf("error while getting container info: %v", err)
	}
	defer rows.Close()

	cts := &ContainerStat{}
	pingDurations := make([]string, 0)
	for rows.Next() {
		var pingDuration string
		if err := rows.Scan(&pingDuration); err != nil {
			return nil, fmt.Errorf("error returned from method `Scan`, package `sql`: %v", err)
		}
		pingDurations = append(pingDurations, pingDuration)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating over rows returned by query: %v", err)
	}

	var successPingTime string
	query = `SELECT success_ping_time FROM containers WHERE ipv4 = $1;`
	err = repo.dtb.QueryRow(query, ipv4).Scan(&successPingTime)
	if err != nil {
		return nil, fmt.Errorf("error while getting successful ping time: %v", err)
	}

	cts.PingDurations = pingDurations
	cts.SuccessPingTime = successPingTime
	return cts, nil
}
