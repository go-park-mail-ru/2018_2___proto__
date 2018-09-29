package api

type ApiService struct {
	Users    IUserStorage
	Sessions ISessionStorage
}

func NewApiService() *ApiService {
	return &ApiService {
		Users: NewUserStorage(),
		Sessions: NewSessionStorage(),
	}
}