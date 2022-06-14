# POST COLLECTION
```bash
    #you can import post collection with this link
    https://www.getpostman.com/collections/3cbd38408b97b3796b8d

```
# Cara Menjalankan Project
```bash
    - Manual Tanpa Docker
        - Install MongoDb terlebih dahulu, lalu configurasikan ke env
        - masuk ke dalam project
        - lalu jalankan perintah go run main.go 
    - Dengan Docker
        - Install Docker terlebih dahulu
        - masuk ke dalam project
        - lalu jalankan perintah
        - docker-compose -f docker-compose.yml up -d

App akan berjalan di port 8000

```


# Api Spesification

## signup http://localhost:8000/users/signup

Request :
- Method : POST
- Endpoint : `users/signup`
- Body :
    raw
```json 
{
  "first_name": "Prayugo",
  "last_name": "Dwi",
  "email": "Prayugo@gmail.com",
  "password": "Prayugo",
  "phone": "+628975666"
}
```

## Login http://localhost:8000/users/login

Request :
- Method : POST
- Endpoint : `users/login`
- Body :
```json 
{
  "email": "Prayugo@gmail.com",
  "password": "Prayugo"
}
```


## Add Product http://localhost:8000/admin/addproduct

Request :
- Method : POST
- Endpoint : `/admin/addproduct`
- Body :
    
```json 
{
  "product_name": "Alienware x10",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg"
}
```

## List Product http://localhost:8000/users/productview

Request :
- Method : GET
- Endpoint : `/users/productview`
- Body :
    
## Search Product http://localhost:8000/users/search?name={keyword}

Request :
- Method : GET
- Endpoint : `/users/search`
- Body :
    QueryParam
        - name={keyword}


## Add Product To Cart http://localhost:8000/addtocart?id={id_product}&userID={id_user}

Request :
- Method : GET
- Endpoint : `/addtocart`
- Header:
    token: {token_user}
- Body :
    QueryParam
        - id={id_product}
        - userID={id_user}

## Remove Product Cart http://localhost:8000/removeitem?id={id_product}&userID={id_user}

Request :
- Method : GET
- Endpoint : `/removeitem`
- Header:
    token: {token_user}
- Body :
    QueryParam
        - id={id_product}
        - userID={id_user}


## List Product Cart http://localhost:8000/listcart?id={id_user}

Request :
- Method : GET
- Endpoint : `/listcart`
- Header:
    token: {token_user}
- Body :
    QueryParam
        - id={id_user}


## Add Home addaddress http://localhost:8000/addaddress?id={id_user}

Request :
- Method : POST
- Endpoint : `/addaddress`
- Header:
    token: {token_user}
- Body :
    QueryParam
        - id={id_user}
    raw
```json 
{
  "house_name": "Rumah Bahagia",
  "street": "Jalan Tirto Sari",
  "city_name": "Purwokerto",
  "pin_code": "20112"
}
```

## Edit Home addaddress http://localhost:8000/edithomeaddress?id={id_user}

Request :
- Method : PUT
- Endpoint : `/edithomeaddress`
- Header:
    token: {token_user}
- Body :
    QueryParam
        - id={id_user}
    raw
```json 
{
  "house_name": "Rumah Sentosa",
  "street": "Jalan Tirto Sari",
  "city_name": "Purwokerto",
  "pin_code": "20112"
}
```

## Delete Home addaddress http://localhost:8000/deleteaddresses?id={id_user}

Request :
- Method : DELETE
- Endpoint : `/deleteaddresses`
- Header:
    token: {token_user}
- Body :
    QueryParam
        - id={id_user}


## Checkout http://localhost:8000/cartcheckout?id={id_user}

Request :
- Method : DELETE
- Endpoint : `/deleteaddresses`
- Header:
    token: {token_user}
- Body :
    QueryParam
        - id={id_user}





