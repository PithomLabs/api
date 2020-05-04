# Sign Process

1. [Registration process](#registration)
2. [Authentication process](#authentication)

## Registration
### How registration process works
User registration is done inside the [registration lambda][0].  

User's values are read from the handler request's body, which the **Content-Type** header can be:
- **application/json**
- **application/x-www-form-urlencoded**
- **multipart/form-data**

While extracting user's informations, a second check of the captcha is done using the key:value pair coming from the **X-Captcha header**. (The first check is done on frontend side)

Then, password validation ([password.Validate][1]), encryption ([bcrypt package][2]) and user's data verification ([register.isValidUser][3]) are performed.  

Finally, while sending a verification mail ([register.sendMail][4]), the user is added to the database ([database.AddUser][5]).

### Registration flow chart

![A representation of the registration process][6]

------

## Authentication
### How authentication process works
User authentication is done using the [authentication lambda][7].

As for registration, user's informations are read from the request's body. The request **Content-Type** header must be one of the following types:
- **application/json**
- **application/x-www-form-urlencoded**
- **multipart/form-data**

After extracting user's informations ([all the extract function following the handle... pattern][8]) authentication is just a matter of **creating the JWT (JSON Web Token)** and **comparing request informations with database ones**.

Verifying user informations  looking for the user in database ([database.GetUserByName][9]), and comparing the hashed password with the request one ([authenticate.checkUser][10]) using [bcrypt][11].

The **token-generation** ([jwt.Create][12]) is done using third-party package [jwt-go][13].

When authentication is finalized, user is **redirected to /set_cookie** on Frontend.

### Authentication flow-chart

![A representation of the authentication process][14]

[0]: https://github.com/komfy/api/blob/master/lambdas/register.go
[1]: https://github.com/komfy/api/blob/master/internal/password/validator.go#L33
[2]: https://godoc.org/golang.org/x/crypto/bcrypt
[3]: https://github.com/komfy/api/blob/master/internal/sign/register/utils.go#L139
[4]: https://github.com/komfy/api/blob/master/internal/sign/register/utils.go#L178
[5]: https://github.com/komfy/api/blob/master/internal/database/user.go#L17
[6]: registration.png

[7]: https://github.com/komfy/api/blob/master/lambdas/authentication.go
[8]: https://github.com/komfy/api/blob/master/internal/sign/authenticate/utils.go#L17
[9]: https://github.com/komfy/api/blob/master/internal/database/user.go#L45
[10]: https://github.com/komfy/api/blob/master/internal/sign/authenticate/utils.go#L88
[11]: https://golang.org/x/crypto/bcrypt
[12]: https://github.com/komfy/api/blob/master/internal/jwt/main.go#L12
[13]: https://github.com/dgrijalva/jwt-go
[14]: authentication.png
