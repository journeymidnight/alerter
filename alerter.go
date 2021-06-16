package alerter

type Alerter interface {
	Name() string
	Send(m Message) error
}

type Message struct {
	Type string
	Info string
}

type AlertType int

const (
	EmailType AlertType = iota
	// Add others
)

type Config interface {
	Type() AlertType
}

func InitAlerters(configs []Config) (alerters []Alerter) {
	for _, config := range configs {
		switch config.Type() {
		case EmailType:
			email := config.(EmailConfig)
			if email.Enable {
				alerters = append(alerters, InitEmailAlert(email.Host, email.Port,
					email.User, email.Password, email.UserFrom, email.UserTo))
			}
		}
	}
	return alerters
}
