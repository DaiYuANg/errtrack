package service

type AuthenticationService struct {
	JWTService *JWTService
}

func NewAuthenticationService(JwtService *JWTService) *AuthenticationService {
	return &AuthenticationService{
		JWTService: JwtService,
	}
}
