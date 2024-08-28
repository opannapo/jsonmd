App : go run ./main.go --file=sample/response-small.json

Input Json 
```json
{
  "success": true,
  "code": 200,
  "message": "Create user successfully",
  "data": {
    "id": 15634,
    "name": "opan napo",
    "birthdate": "17 Aug 1945",
    "nik": {
      "number": "1001212121212112121",
      "name": "Opan Napo ST. MsI",
      "expired_at": "17 Aug 2075"
    },
    "meta": {
      "is_return_config": false,
      "return_config_limit": 0,
      "return_config_reason_ids": "",
      "insurance_type": "",
      "status": {
        "is_enable": true,
        "child_status": {
          "status": false,
          "name": "child1"
        }
      }
    },
    "image": {
      "url": "images/0001.jpg",
      "width": 200,
      "height": 200
    },
    "address": [
      {
        "name": "home 1",
        "address": "JL.Selamat Dunia Akhirat Rt/Rw 001/002"
      },
      {
        "name": "home 2",
        "address": "JL.Selamat Dunia Akhirat Rt/Rw 001/002"
      }
    ],
    "hobbies": [
      "Ibadah",
      "Menabung",
      "Boxing",
      "Music"
    ]
  }
}
```

Output result.md file:
```md



#### result

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| success  | bool  | true | |  | 
| code  | float64  | 200 | |  | 
| message  | string  | Create user successfully | |  | 
| data  | object  | Table : .data | |  | 

#### .data

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| id  | float64  | 15634 | |  | 
| name  | string  | opan napo | |  | 
| birthdate  | string  | 17 Aug 1945 | |  | 
| nik  | object  | Table : .data.nik | |  | 
| meta  | object  | Table : .data.meta | |  | 
| image  | object  | Table : .data.image | |  | 
| address  | array  | Table : .data.address | |  | 
| hobbies  | array  | Table : .data.hobbies | |  | 

#### .data.nik

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| number  | string  | 1001212121212112121 | |  | 
| name  | string  | Opan Napo ST. MsI | |  | 
| expired_at  | string  | 17 Aug 2075 | |  | 

#### .data.meta

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_return_config  | bool  | false | |  | 
| return_config_limit  | float64  | 0 | |  | 
| return_config_reason_ids  | string  |  | |  | 
| insurance_type  | string  |  | |  | 
| status  | object  | Table : .data.meta.status | |  | 

#### .data.meta.status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_enable  | bool  | true | |  | 
| child_status  | object  | Table : .data.meta.status.child_status | |  | 

#### .data.meta.status.child_status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| status  | bool  | false | |  | 
| name  | string  | child1 | |  | 

#### .data.image

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| width  | float64  | 200 | |  | 
| height  | float64  | 200 | |  | 
| url  | string  | images/0001.jpg | |  | 

#### .data.address

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| name  | string  | home 1 | |  | 
| address  | string  | JL.Selamat Dunia Akhirat Rt/Rw 001/002 | |  | 

#### .data.hobbies

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
|   | string  | Ibadah | |  | 
```

Output Preview:


#### result

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| success  | bool  | true | |  | 
| code  | float64  | 200 | |  | 
| message  | string  | Create user successfully | |  | 
| data  | object  | Table : .data | |  | 

#### .data

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| id  | float64  | 15634 | |  | 
| name  | string  | opan napo | |  | 
| birthdate  | string  | 17 Aug 1945 | |  | 
| nik  | object  | Table : .data.nik | |  | 
| meta  | object  | Table : .data.meta | |  | 
| image  | object  | Table : .data.image | |  | 
| address  | array  | Table : .data.address | |  | 
| hobbies  | array  | Table : .data.hobbies | |  | 

#### .data.nik

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| number  | string  | 1001212121212112121 | |  | 
| name  | string  | Opan Napo ST. MsI | |  | 
| expired_at  | string  | 17 Aug 2075 | |  | 

#### .data.meta

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_return_config  | bool  | false | |  | 
| return_config_limit  | float64  | 0 | |  | 
| return_config_reason_ids  | string  |  | |  | 
| insurance_type  | string  |  | |  | 
| status  | object  | Table : .data.meta.status | |  | 

#### .data.meta.status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| is_enable  | bool  | true | |  | 
| child_status  | object  | Table : .data.meta.status.child_status | |  | 

#### .data.meta.status.child_status

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| status  | bool  | false | |  | 
| name  | string  | child1 | |  | 

#### .data.image

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| width  | float64  | 200 | |  | 
| height  | float64  | 200 | |  | 
| url  | string  | images/0001.jpg | |  | 

#### .data.address

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
| name  | string  | home 1 | |  | 
| address  | string  | JL.Selamat Dunia Akhirat Rt/Rw 001/002 | |  | 

#### .data.hobbies

| Name | Data Type | Sample Value | Is Nullable | Description |
|------|-----------|-------------|--------------|-------------|
|   | string  | Ibadah | |  | 