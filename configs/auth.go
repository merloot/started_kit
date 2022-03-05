package configs

type AuthConfig struct {
	JwtAccessKey  string
	JwtAccessTTL  uint
	JwtRefreshKey string
	JwtRefreshTTL uint
}
