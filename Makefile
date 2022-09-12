db_up:
	docker run --rm --name pg-intro -p 5432:5432 -e POSTGRES_PASSWORD=Ohfaeyi7Zahz -e POSTGRES_DB=website postgres

