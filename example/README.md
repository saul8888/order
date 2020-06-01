
**POST**
/api/employee
```json
{
	"object": "test11",
	"locationId": "quispe",
	"email":"test11@example.com",
	"status": "activo",
	"addresses": {
		"country":"Peru",
		"city": "Lima",
		"postalCode": "15066"
	}
}
```
employee_id: 5ed0238bf7070da9a85cb950
location_id: 5ecd8ea5b4117f07a32d8883
**GET**
/api/employee?Id=5ecbf37c29bb3ef231594f75

**PUT**
/api/employee?employeeId=5ecbf37c29bb3ef231594f75
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
/api/employee/total
```json
{
	"limit": "1",
	"offset": "0"
}
```

**DELETE**
/api/employee?Id=5ecbf37c29bb3ef231594f75


