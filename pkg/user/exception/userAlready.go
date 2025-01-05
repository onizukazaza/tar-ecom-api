package exception

type EmailCheck struct {

}

func(e *EmailCheck) Error() string {
	return "failed to check email existence"
}