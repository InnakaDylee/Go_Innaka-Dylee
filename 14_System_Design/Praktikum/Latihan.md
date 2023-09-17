# Latihan
### 1. Gambarkan desain ERD dari sistem pembelian tiket kereta api!
* More Detail [ERD](https://viewer.diagrams.net/?page-id=W5dUf4NcdmSJ9-Wtqa15&highlight=0000ff&edit=_blank&layers=1&nav=1&page-id=W5dUf4NcdmSJ9-Wtqa15&viewbox=%7B%22x%22%3A-100%2C%22y%22%3A-160%2C%22width%22%3A1621%2C%22height%22%3A831%2C%22border%22%3A100%7D#G1fenU1VWJol3SQcvGcw1HRr7-94XDyGfo)

----

### 2. Gambarkan use case diagram dari sistem pembelian tiket kereta api!
* More Detail [Usecase_Diagram](https://viewer.diagrams.net/?page-id=54cnupEtWff-8IgcIQ6t&highlight=0000ff&edit=_blank&layers=1&nav=1&page-id=54cnupEtWff-8IgcIQ6t#G1fenU1VWJol3SQcvGcw1HRr7-94XDyGfo)

----

### 3. Terdapat sebuah query pada SQL yaitu SELECT * FROM users; Dengan tujuan yang sama, tuliskan dalam bentuk perintah:
* Redis
	```
	SCAN 0 MATCH users:*
	```
* Neo4j
	```
	MATCH (u:users)
	RETURN u
	```
* Cassandra
	```
	SELECT * FROM users;
	```