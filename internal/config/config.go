package config

type (
	App struct {
		Server   Server
		Database Database
	}
)

func NewApp() App {
	return App{}
}
