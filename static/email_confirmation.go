package static

func EmailConfirmation(code string) string {
	return `<a href="http://localhost:3005/api/email-confirmation/` + code + `">Confirmar Email</a>`
}
