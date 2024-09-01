# A Basic JWT Auth Implementation in Golang

This is a basic implementation of JWT authentication in Golang. It uses the `github.com/golang-jwt/jwt/v5` library for JWT token generation and validation.

`.env` setup:

```bash
PUBLIC_HOST=localhost
PORT=8080
DB_NAME=app.db
JWT_EXP=604800
JWT_SECRET=secret
```

To Run Tests:
```bash
make test
```

To Run the Server:
```bash
make run
```

