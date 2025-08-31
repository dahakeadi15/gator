package config

type Config struct {
	DatabaseURL     string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(name string) error {
	c.CurrentUserName = name

	err := write(c)
	if err != nil {
		return err
	}

	return nil
}
