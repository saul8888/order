
**POST**
/api/catalog
```json
{
	"name": "water",
	"category": "to drink",
	"description":"to drink",
	"locationId": "5ed33e1b71751c59d9c1af7f",
	"options1": "small",
	"marcas": [
		{
			"name":"cielo",
			"SKU":"SAQO",
			"unit":"botle",
			"price":5,
			"currency":"S/.",
			"stock": {
				"description":"receive",
				"inStock":32,
				"alertStock":10
			}
		},
		{
			"name":"san mateo",
			"SKU":"JAQO",
			"unit":"botle",
			"price":6,
			"currency":"S/.",
			"stock": {
				"description":"receive",
				"inStock":25,
				"alertStock":8
			}
		}
	]
}
```

**GET**
/api/location?Id=5ecd8dd1b4117f07a32d8882

**PUT**
/api/location?locationId=5ecbf37c29bb3ef231594f75
```json
{
	"firstName": "saul",
	"lastName": "quispe",
	"email":"saul@example.com",
	"phoneNumber": "9445-98989",
	"addresses": "chorrillos"
}
```

**GET**
/api/location/total
```json
{
	"limit": 2,
	"offset": 0
}
```

**DELETE**
/api/location?Id=5ecbf37c29bb3ef231594f75