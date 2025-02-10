package backend

import (
	"fmt"
)

func (repo *BackendDBRepository) GetLogs() ([]Container, error) {
	query := `SELECT prs.id, cns.ipv4, prs.ping_time, cns.success_ping_time
                  FROM containers cns
                  JOIN ping_results prs ON cns.ipv4 = prs.container_ipv4;`

	rows, err := repo.dtb.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error while getting containers info: %v", err)
	}
	defer rows.Close()

	cns := make([]Container, 0)
	for rows.Next() {
		ctr := Container{}
		if err := rows.Scan(&ctr.ID, &ctr.IPv4, &ctr.PingDuration, &ctr.SuccessPingTime); err != nil {
			return nil, fmt.Errorf("error returned from method `Scan`, package `sql`: %v", err)
		}
		cns = append(cns, ctr)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating over rows returned by query: %v", err)
	}
	return cns, nil
}
