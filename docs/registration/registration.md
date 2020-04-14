# Registration Process
### User registration on back-end

User registration is done inside the [registration lambda][0].  

User's values are read from the handler's request for which the **Content-Type** header can be:
- application/json
- application/x-www-form-urlencoded
- multipart/form-data

While extracting user's informations, a second check of the captcha is done using the key:value pair coming from the **X-Captcha header**.

Then, password validation ([password.Validate][1]), encryption ([bcrypt package][2]) and user's data verification ([register.isValidUser][3]) are performed.  

Finally, while sending a verification mail ([register.sendMail][4]), the user is added to the database ([database.AddUser][5]).

### Registration flow chart

![A representation of the registration process][6]

[0]: https://github.com/komfy/api/blob/master/lambdas/register.go
[1]: https://github.com/komfy/api/blob/master/internal/password/validator.go#L33
[2]: https://godoc.org/golang.org/x/crypto/bcrypt
[3]: https://github.com/komfy/api/blob/master/internal/sign/register/utils.go#L139
[4]: https://github.com/komfy/api/blob/master/internal/sign/register/utils.go#L178
[5]: https://github.com/komfy/api/blob/master/internal/database/user.go#L17
[6]: registration.png

