
**POST**
/api/merchant
```json
{
	"name": "test11",
	"status": "activo",
	"languageCode":"Peru",
	"currency":"USD",
	"addresses": {
		"country":"Peru",
		"state":"Lima",
		"city":"chorrillos",
		"street":"street",
		"postalcode":"15066"
	}
}
```

**GET**
/api/merchant?Id=5ed339a7e7259a2e031138e5

**PUT**
/api/merchant?Id=5ed339a7e7259a2e031138e5
```json
{
	"firstName": "saul",
}
```

**GET**
/api/merchant/total
```json
{
	"name": "dany",
	"status": "activo",
	"currency":"USD",
	"addresses": {
		"country":"USA"
	}
}
```

**DELETE**
/api/merchant?Id=[ID]