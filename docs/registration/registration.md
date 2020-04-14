# Registration Process
### User registration on back-end

User registration is done inside the [registration lambda][0].  

User's values are read from the handler's request (`http.Request.Body`) where the **Content-Type** header of the request can be:
- application/json
- application/x-www-form-urlencoded
- multipart/form-data

While extracting the user informations, a second check of the captcha is done using the key:value pair from the **X-Captcha header**.

Then, password validation (`password.Validate`), encryption using [bcrypt package][1] and user's data verification (`register.isValidUser`) are performed.  

Finally, while sending a verification mail (`register.sendMail`), the user is added to the database (`database.AddUser`).

### Registration flow chart

![A representation of the registration process][2]

[0]: https://github.com/komfy/api
[1]: https://godoc.org/golang.org/x/crypto/bcrypt
[2]: registration.png

