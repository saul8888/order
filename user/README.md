
**POST**
/api/employee
```json
{
	"firstName": "test11",
	"lastName": "quispe",
	"rolesId": "chef",
	"locationId": "1234",
	"email":"test11@example.com",
	"password":"test1234",
	"status": "activo",
	"phoneNumber": "9445-98989",
	"addresses": "chorrillos"
}
```

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


