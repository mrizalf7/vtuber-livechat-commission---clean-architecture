# vtuber-livechat-commission app
This is a restful api golang project using docker,jwt,rdbms with 4 tables,postgres database,login,register.

# Documentation
## health 
to check current server is alive:

<b>GET</b>
```
localhost:4000
```
  
Response (Status: 200)
```
Welcome
```

## register
Registering a new user

<b>POST</b>
```
localhost:4000/api/auth/register
```

Request Body
```
{
    "username":"dahell",
    "email":"dahell@yahoo.com",
    "password":"123"
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 3,
        "username": "dahell",
        "email": "dahell@yahoo.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMyIsImV4cCI6MTcxNjU2OTUzNiwiaWF0IjoxNjg0OTQ3MTM3LCJpc3MiOiJ5ZGhud2IifQ.IVQtMbZO0cmKrEUU82t_UN-QfYQrLIpaVgHxLej_TJA"
    }
}
```


## login
Authenticate by email and password

<b>POST</b>
```
localhost:4000/api/auth/login
```

Request body
```
{
    "email":"dahell@yahoo.com",
    "password":"123"
}
```

Response Success (Status: 200)
```
"status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 3,
        "username": "dahell",
        "email": "dahell@yahoo.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMyIsImV4cCI6MTcxNjU2OTY2NSwiaWF0IjoxNjg0OTQ3MjY1LCJpc3MiOiJ5ZGhud2IifQ.Kx70VI_vKBLh8uwVn4YRqlWsdohxvX4U1-OOYzfjnvo"
    }
```
## Get All Registered Users

<b>GET</b>
```
localhost:4000/api/user/registered
```

Response success (status: 200)
```
{
    "status": true,
    "message": "Sucess Get Users",
    "errors": null,
    "data": [
        {
            "id": 1,
            "username": "Rijal",
            "email": "rijal@yahoo.com"
        },
        {
            "id": 2,
            "username": "rizal",
            "email": "rizal@yahoo.com"
        },
        {
            "id": 3,
            "username": "dahell",
            "email": "dahell@yahoo.com"
        }
    ]
}
```

## Get Profile By Token
Get current info from logged user

<b>GET</b>
```
localhost:4000/api/user/profile
```

Headers
```
Authorization: yourToken
```

Response success (status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 3,
        "username": "dahell",
        "email": "dahell@yahoo.com"
    }
}
```
## Update profile
Update user data who logged in

<b>PUT</b>
```
localhost:4000/api/user/profile
```

Headers
```
Authorization: yourToken
```

Request body
```
{
    "username" :"hellda",
    "email" :"hellda@yahoo.com"
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 3,
        "username": "hellda",
        "email": "hellda@yahoo.com"
    }
}
```


<b>=============================================</b>
## Create a project

<b>POST</b>
```
localhost:4000/api/project
```

Headers
```
Authorization : yourToken
```

Request body
```
{
"name" : "Project LiveChat Etsy",
"description" : "Blue Theme Design"
}


```
Response success (Status: 201)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Project LiveChat Etsy",
        "description": "Blue Theme Design",
        "created_at": "2023-05-25T10:34:30.14795Z",
        "updated_at": "2023-05-25T10:34:30.14795Z"
    }
}
```
## Get projects

<b>GET</b>
```
localhost:4000/api/project
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": [
        {
            "id": 1,
            "name": "Project LiveChat Etsy",
            "description": "Blue Theme Design",
            "created_at": "2023-05-25T10:34:30.14795Z",
            "updated_at": "2023-05-25T10:34:30.14795Z"
        },
        {
            "id": 2,
            "name": "Project LiveChat Etsy(1)",
            "description": "Blue Theme Design(1)",
            "created_at": "2023-05-25T10:38:02.181392Z",
            "updated_at": "2023-05-25T10:38:02.181392Z"
        }
    ]
}
```

## Find a project by id
Find project by id

<b>GET</b>
```
localhost:4000/api/project/{id}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 2,
        "name": "Project LiveChat Etsy(1)",
        "description": "Blue Theme Design(1)",
        "created_at": "2023-05-25T10:38:02.181392Z",
        "updated_at": "2023-05-25T10:38:02.181392Z"
    }
}
```

## Update project

<b>PUT</b>
```
localhost:4000/api/project/{id}
```


Headers

```
Authorization : yourToken
```

Request body
```
{   
     "name": "Project LiveChat Rizal",
     "description": "Brown Theme Design"
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 2,
        "name": "Project LiveChat Rizal",
        "description": "Brown Theme Design",
        "created_at": "2023-05-25T10:38:02.181392Z",
        "updated_at": "2023-05-25T10:59:25.957269Z"
    }
}
```

## Delete project

<b>DELETE</b>
```
localhost:4000/api/project/{id}
```
Headers
```
Authorization: yourToken
```
Response success (Status: 200)
```
{
    "status": true,
    "message": "Deleted",
    "errors": null,
    "data": {
        "id": 2,
        "name": "Project LiveChat Rizal",
        "description": "Brown Theme Design",
        "created_at": "2023-05-25T10:38:02.181392Z",
        "updated_at": "2023-05-25T10:59:25.957269Z"
    }
}
```

## Create Price List

<b>POST</b>
```
localhost:4000/api/pricelist
```

Headers
```
Authorization : yourToken
```

Request body 
```
{
    "category" :"Premium",
    "description":"Cool design with premium price",
    "price_idr":"300K",
    "price_usd":"20$",
    "project_id" :1
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "price_list_id": 1,
        "category": "Premium",
        "description": "Cool design with premium price",
        "price_idr": "300K",
        "price_usd": "20$",
        "project_id": 1,
        "Project": {
            "id": 1,
            "name": "Project LiveChat Etsy",
            "description": "Blue Theme Design",
            "created_at": "2023-05-25T10:34:30.14795Z",
            "updated_at": "2023-05-25T10:34:30.14795Z"
        },
        "created_at": "2023-05-25T11:54:11.30405+07:00",
        "updated_at": "2023-05-25T11:54:11.30405+07:00"
    }
}
```

## Get Price List

<b>GET<b>

```
localhost:4000/api/pricelist
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": [
        {
            "price_list_id": 1,
            "category": "Premium",
            "description": "Cool design with premium price",
            "price_idr": "300K",
            "price_usd": "20$",
            "project_id": 1,
            "Project": {
                "id": 1,
                "name": "Project LiveChat Etsy",
                "description": "Blue Theme Design",
                "created_at": "2023-05-25T10:34:30.14795Z",
                "updated_at": "2023-05-25T10:34:30.14795Z"
            },
            "created_at": "2023-05-25T11:54:11.30405+07:00",
            "updated_at": "2023-05-25T11:54:11.30405+07:00"
        },
        {
            "price_list_id": 4,
            "category": "Norma1l",
            "description": "Nor1mal design with normal price",
            "price_idr": "150K1",
            "price_usd": "10$1",
            "project_id": 3,
            "Project": {
                "id": 3,
                "name": "Project LiveChat Rizal",
                "description": "Brown Theme Design",
                "created_at": "2023-05-25T12:06:43.942178Z",
                "updated_at": "2023-05-25T12:06:43.942178Z"
            },
            "created_at": "2023-05-25T12:17:22.214225+07:00",
            "updated_at": "2023-05-25T12:17:22.214225+07:00"
        },
        {
            "price_list_id": 5,
            "category": "Normaasdasd1l",
            "description": "Nor1mal designasdasd with normal price",
            "price_idr": "150adasK1",
            "price_usd": "10$asdasd1",
            "project_id": 3,
            "Project": {
                "id": 3,
                "name": "Project LiveChat Rizal",
                "description": "Brown Theme Design",
                "created_at": "2023-05-25T12:06:43.942178Z",
                "updated_at": "2023-05-25T12:06:43.942178Z"
            },
            "created_at": "2023-05-25T12:18:34.354326+07:00",
            "updated_at": "2023-05-25T12:18:34.354326+07:00"
        }
    ]
}
```

## Find Price List By Id

<b>GET<b>

```
localhost:4000/api/pricelist/{id}
```

Response success (Status :200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "price_list_id": 1,
        "category": "Premium",
        "description": "Cool design with premium price",
        "price_idr": "300K",
        "price_usd": "20$",
        "project_id": 1,
        "Project": {
            "id": 1,
            "name": "Project LiveChat Etsy",
            "description": "Blue Theme Design",
            "created_at": "2023-05-25T10:34:30.14795Z",
            "updated_at": "2023-05-25T10:34:30.14795Z"
        },
        "created_at": "2023-05-25T11:54:11.30405+07:00",
        "updated_at": "2023-05-25T11:54:11.30405+07:00"
    }
}
```
## Update Price List

<b>PUT<b>

```
localhost:4000/api/pricelist/{id}
```

Headers
```
Authorization : yourToken
```

Request body 
```
{
    "category" :"Normaasdasd1l",
    "description":"Nor1mal designasdasd with normal price",
    "price_idr":"150adasK1",
    "price_usd":"10$asdasd1",
    "project_id" :3
}
```
Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "price_list_id": 1,
        "category": "Normaasdasd1l",
        "description": "Nor1mal designasdasd with normal price",
        "price_idr": "150adasK1",
        "price_usd": "10$asdasd1",
        "project_id": 3,
        "Project": {
            "id": 3,
            "name": "Project LiveChat Rizal",
            "description": "Brown Theme Design",
            "created_at": "2023-05-25T12:06:43.942178Z",
            "updated_at": "2023-05-25T12:06:43.942178Z"
        },
        "created_at": "2023-05-25T11:54:11.30405+07:00",
        "updated_at": "2023-05-25T13:10:11.763445+07:00"
    }
}
```


## Delete Price List
<b>DELETE<b>
```
localhost:4000/api/pricelist/{id}
```
Headers 
```
Authorization : yourToken
```
Response success (Status: 200)
```
{
    "status": true,
    "message": "Deleted",
    "errors": null,
    "data": {
        "price_list_id": 5,
        "category": "Normaasdasd1l",
        "description": "Nor1mal designasdasd with normal price",
        "price_idr": "150adasK1",
        "price_usd": "10$asdasd1",
        "project_id": 3,
        "Project": {
            "id": 3,
            "name": "Project LiveChat Rizal",
            "description": "Brown Theme Design",
            "created_at": "2023-05-25T12:06:43.942178Z",
            "updated_at": "2023-05-25T12:06:43.942178Z"
        },
        "created_at": "2023-05-25T12:18:34.354326+07:00",
        "updated_at": "2023-05-25T12:18:34.354326+07:00"
    }
}
```

## Create Commission
<b>POST<b>
```
localhost:4000/api/commission
```

Headers
```
Authorization : yourToken
```
Request body 
```
{
    "name":"Etsy",
    "twitter_profile_url":"twitter/profile",
    "profile_pictModel" :"www.poto.png",
    "live_chat_picture" :"www.foto.png",
    "youtube_url" :"youurl/1",
    "status" : "pending",
    "project_id" :4,
    "price_list_id":1
}
```

Response success (Status : 201)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Etsy",
        "twitter_profile_url": "twitter/profile",
        "profile_pictModel": "www.poto.png",
        "live_chat_picture": "www.foto.png",
        "youtube_url": "youurl/1",
        "status": "pending",
        "price_list_id": 1,
        "PriceList": {
            "id": 1,
            "category": "Normaasdasd1l",
            "description": "Nor1mal designasdasd with normal price",
            "price_idr": "150adasK1",
            "price_usd": "10$asdasd1",
            "project_id": 3,
            "Project": {
                "id": 3,
                "name": "Project LiveChat Rizal",
                "description": "Brown Theme Design",
                "created_at": "2023-05-25T12:06:43.942178Z",
                "updated_at": "2023-05-25T12:06:43.942178Z"
            },
            "created_at": "2023-05-25T11:54:11.30405+07:00",
            "updated_at": "2023-05-25T13:10:11.763445+07:00"
        },
        "project_id": 4,
        "Project": {
            "id": 4,
            "name": "Project LiveChat Rizal(1)",
            "description": "Brown Theme Design(1)",
            "created_at": "2023-05-25T12:15:44.640946Z",
            "updated_at": "2023-05-25T12:15:44.640946Z"
        },
        "created_at": "2023-05-25T17:59:12.763307+07:00",
        "updated_at": "2023-05-25T17:59:12.763307+07:00"
    }
}
```

## Get Commissions
<b>GET<b>

```
localhost:4000/api/commission
```

Response success (Status: 200)
```
{
            "id": 15,
            "name": "Etsy",
            "twitter_profile_url": "twitter/profile",
            "profile_pictModel": "www.poto.png",
            "live_chat_picture": "www.foto.png",
            "youtube_url": "youurl/1",
            "status": "pending",
            "price_list_id": 1,
            "PriceList": {
                "id": 1,
                "category": "Normaasdasd1l",
                "description": "Nor1mal designasdasd with normal price",
                "price_idr": "150adasK1",
                "price_usd": "10$asdasd1",
                "project_id": 3,
                "Project": {
                    "id": 3,
                    "name": "Project LiveChat Rizal",
                    "description": "Brown Theme Design",
                    "created_at": "2023-05-25T12:06:43.942178Z",
                    "updated_at": "2023-05-25T12:06:43.942178Z"
                },
                "created_at": "2023-05-25T11:54:11.30405+07:00",
                "updated_at": "2023-05-25T13:10:11.763445+07:00"
            },
            "project_id": 4,
            "Project": {
                "id": 4,
                "name": "Project LiveChat Rizal(1)",
                "description": "Brown Theme Design(1)",
                "created_at": "2023-05-25T12:15:44.640946Z",
                "updated_at": "2023-05-25T12:15:44.640946Z"
            },
            "created_at": "2023-05-25T17:20:55.112433+07:00",
            "updated_at": "2023-05-25T17:20:55.112433+07:00"
        },
        {
            "id": 16,
            "name": "Etsy",
            "twitter_profile_url": "twitter/profile",
            "profile_pictModel": "www.poto.png",
            "live_chat_picture": "www.foto.png",
            "youtube_url": "youurl/1",
            "status": "pending",
            "price_list_id": 1,
            "PriceList": {
                "id": 1,
                "category": "Normaasdasd1l",
                "description": "Nor1mal designasdasd with normal price",
                "price_idr": "150adasK1",
                "price_usd": "10$asdasd1",
                "project_id": 3,
                "Project": {
                    "id": 3,
                    "name": "Project LiveChat Rizal",
                    "description": "Brown Theme Design",
                    "created_at": "2023-05-25T12:06:43.942178Z",
                    "updated_at": "2023-05-25T12:06:43.942178Z"
                },
                "created_at": "2023-05-25T11:54:11.30405+07:00",
                "updated_at": "2023-05-25T13:10:11.763445+07:00"
            },
            "project_id": 4,
            "Project": {
                "id": 4,
                "name": "Project LiveChat Rizal(1)",
                "description": "Brown Theme Design(1)",
                "created_at": "2023-05-25T12:15:44.640946Z",
                "updated_at": "2023-05-25T12:15:44.640946Z"
            },
            "created_at": "2023-05-25T17:59:12.763307+07:00",
            "updated_at": "2023-05-25T17:59:12.763307+07:00"
        }

```
## Find Commission By Id
<b>GET</b>

```
localhost:4000/api/commission/{id}
```
Response success (Status: 200)

```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 16,
        "name": "Etsy",
        "twitter_profile_url": "twitter/profile",
        "profile_pictModel": "www.poto.png",
        "live_chat_picture": "www.foto.png",
        "youtube_url": "youurl/1",
        "status": "pending",
        "price_list_id": 1,
        "PriceList": {
            "id": 1,
            "category": "Normaasdasd1l",
            "description": "Nor1mal designasdasd with normal price",
            "price_idr": "150adasK1",
            "price_usd": "10$asdasd1",
            "project_id": 3,
            "Project": {
                "id": 3,
                "name": "Project LiveChat Rizal",
                "description": "Brown Theme Design",
                "created_at": "2023-05-25T12:06:43.942178Z",
                "updated_at": "2023-05-25T12:06:43.942178Z"
            },
            "created_at": "2023-05-25T11:54:11.30405+07:00",
            "updated_at": "2023-05-25T13:10:11.763445+07:00"
        },
        "project_id": 4,
        "Project": {
            "id": 4,
            "name": "Project LiveChat Rizal(1)",
            "description": "Brown Theme Design(1)",
            "created_at": "2023-05-25T12:15:44.640946Z",
            "updated_at": "2023-05-25T12:15:44.640946Z"
        },
        "created_at": "2023-05-25T17:59:12.763307+07:00",
        "updated_at": "2023-05-25T17:59:12.763307+07:00"
    }
}{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 16,
        "name": "Etsy",
        "twitter_profile_url": "twitter/profile",
        "profile_pictModel": "www.poto.png",
        "live_chat_picture": "www.foto.png",
        "youtube_url": "youurl/1",
        "status": "pending",
        "price_list_id": 1,
        "PriceList": {
            "id": 1,
            "category": "Normaasdasd1l",
            "description": "Nor1mal designasdasd with normal price",
            "price_idr": "150adasK1",
            "price_usd": "10$asdasd1",
            "project_id": 3,
            "Project": {
                "id": 3,
                "name": "Project LiveChat Rizal",
                "description": "Brown Theme Design",
                "created_at": "2023-05-25T12:06:43.942178Z",
                "updated_at": "2023-05-25T12:06:43.942178Z"
            },
            "created_at": "2023-05-25T11:54:11.30405+07:00",
            "updated_at": "2023-05-25T13:10:11.763445+07:00"
        },
        "project_id": 4,
        "Project": {
            "id": 4,
            "name": "Project LiveChat Rizal(1)",
            "description": "Brown Theme Design(1)",
            "created_at": "2023-05-25T12:15:44.640946Z",
            "updated_at": "2023-05-25T12:15:44.640946Z"
        },
        "created_at": "2023-05-25T17:59:12.763307+07:00",
        "updated_at": "2023-05-25T17:59:12.763307+07:00"
    }
}
```

## Update Commission
<b>PUT</b>
```
localhost:4000/api/commission/{id}
```
Headers
```
Authorization : yourToken
```

Request body
```
{
    "name":"Rizal",
    "twitter_profile_url":"twitter/profile",
    "profile_pictModel" :"www.poto.png",
    "live_chat_picture" :"www.foto.png",
    "youtube_url" :"youurl/1",
    "status" : "pending",
    "project_id" :5,
    "price_list_id":4
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 16,
        "name": "Rizal",
        "twitter_profile_url": "twitter/profile",
        "profile_pictModel": "www.poto.png",
        "live_chat_picture": "www.foto.png",
        "youtube_url": "youurl/1",
        "status": "pending",
        "price_list_id": 4,
        "PriceList": {
            "id": 4,
            "category": "Norma1l",
            "description": "Nor1mal design with normal price",
            "price_idr": "150K1",
            "price_usd": "10$1",
            "project_id": 3,
            "Project": {
                "id": 3,
                "name": "Project LiveChat Rizal",
                "description": "Brown Theme Design",
                "created_at": "2023-05-25T12:06:43.942178Z",
                "updated_at": "2023-05-25T12:06:43.942178Z"
            },
            "created_at": "2023-05-25T12:17:22.214225+07:00",
            "updated_at": "2023-05-25T12:17:22.214225+07:00"
        },
        "project_id": 5,
        "Project": {
            "id": 5,
            "name": "Project LiveChat Rizal(11)",
            "description": "Brown Theme Design(11)",
            "created_at": "2023-05-25T15:25:07.167193Z",
            "updated_at": "2023-05-25T15:25:07.167193Z"
        },
        "created_at": "2023-05-25T17:59:12.763307+07:00",
        "updated_at": "2023-05-25T19:19:55.186257+07:00"
    }
}
```
## Delete Commission
<b>DELETE</b>
```
localhost:4000/api/commission
```
Headers
```
Authorization: yourToken
````
Response success (Status: 200)
```
{
    "status": true,
    "message": "Deleted",
    "errors": null,
    "data": {
        "id": 16,
        "name": "Rizal",
        "twitter_profile_url": "twitter/profile",
        "profile_pictModel": "www.poto.png",
        "live_chat_picture": "www.foto.png",
        "youtube_url": "youurl/1",
        "status": "pending",
        "price_list_id": 4,
        "PriceList": {
            "id": 4,
            "category": "Norma1l",
            "description": "Nor1mal design with normal price",
            "price_idr": "150K1",
            "price_usd": "10$1",
            "project_id": 3,
            "Project": {
                "id": 3,
                "name": "Project LiveChat Rizal",
                "description": "Brown Theme Design",
                "created_at": "2023-05-25T12:06:43.942178Z",
                "updated_at": "2023-05-25T12:06:43.942178Z"
            },
            "created_at": "2023-05-25T12:17:22.214225+07:00",
            "updated_at": "2023-05-25T12:17:22.214225+07:00"
        },
        "project_id": 5,
        "Project": {
            "id": 5,
            "name": "Project LiveChat Rizal(11)",
            "description": "Brown Theme Design(11)",
            "created_at": "2023-05-25T15:25:07.167193Z",
            "updated_at": "2023-05-25T15:25:07.167193Z"
        },
        "created_at": "2023-05-25T17:59:12.763307+07:00",
        "updated_at": "2023-05-25T19:19:55.186257+07:00"
    }
}
```
~~~
TEST
~~~