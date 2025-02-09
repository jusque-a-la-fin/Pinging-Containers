package backend

import (
	"log"
)

func (repo *BackendDBRepository) GetLogs() ([]Container, error) {
	query := `SELECT prs.id, cns.ipv4, prs.ping_time, cns.success_ping_time
                  FROM containers cns
                  JOIN ping_results prs ON cns.ipv4 = prs.container_ipv4;`

	rows, err := repo.dtb.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cns := make([]Container, 0)
	for rows.Next() {
		ctr := Container{}
		if err := rows.Scan(&ctr.ID, &ctr.IPv4, &ctr.PingDuration, &ctr.SuccessPingTime); err != nil {
			log.Fatal(err)
		}
		cns = append(cns, ctr)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return cns, nil
}
