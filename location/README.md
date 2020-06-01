
**POST**
/api/location
```json
{
	"name": "test11",
	"merchantId": [merchantID],
	"email": "test11@example.com",
	"addresses": {
		"country":"Peru",
		"state":"Lima",
		"city":"chorrillos",
		"street":"street",
		"postalcode":"15066"
	},
	"currency":"test11@example.com",
	"description":"test1234",
	"website":"test1234",
	"twitter":"test1234",
	"instagram":"test1234",
	"status": "activo",
	"phoneNumber": "9445-98989"
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


