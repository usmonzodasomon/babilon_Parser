package db

const create = `CREATE TABLE IF NOT EXISTS netflow (
    source_ip TEXT,
    destination_ip TEXT,
    packets INT,
    bytes INT,
    sport INT,
    dport INT,
    proto INT,
    account_id INT,
    tclass INT,
    date_time INT,
    nfsource_ip TEXT
);`

func up() {
	_, err := DB.Exec(create)
	if err != nil {
		panic(err)
	}
}
