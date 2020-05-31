package shortener

type RedirectServiceInterface interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
