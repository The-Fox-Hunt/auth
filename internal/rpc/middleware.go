package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const jwtSecret = "carrot"

// JWTAuthInterceptor это интерсептор для проверки JWT-токенов в gRPC-запросах.
func JWTAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	log.Printf("Received gRPC method: %s", info.FullMethod) // Логируем вызываемый метод
	// Исключения для определенных методов.
	switch info.FullMethod {
	case "/AuthService/Signup":
		fallthrough
	case "/AuthService/Login":
		return handler(ctx, req)
	}

	// Получение токена из заголовков запроса.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get metadata header")
	}

	tokenStrings := md.Get("authorization")
	if len(tokenStrings) == 0 {
		return nil, fmt.Errorf("there are no strings in the authorization header")
	}

	//берем первый элемент, т.к. md.get("authorization) возвращает массив строк
	tokenString := tokenStrings[0]

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Парсинг и проверка токена.
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("incorrect token signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		log.Printf("Received gRPC method: %s", info.FullMethod) // Логируем вызываемый метод
		return nil, fmt.Errorf("cant parse token: %w", err)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		if username, ok := claims["username"].(string); ok {
			// Добавляем username в контекст
			ctx = context.WithValue(ctx, "username", username)
		}
	}
	return handler(ctx, req)
}
