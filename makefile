.PHONY: psql

psql:
	@psql -h ${POSTGRES_HOSTNAME} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_DB}
