create_db:
	PGPASSWORD=postgres createdb -h localhost -U postgres -e url_shorter
migration:
	PGPASSWORD=postgres  psql -h localhost -U postgres -d url_shorter -f migrations/init.sql