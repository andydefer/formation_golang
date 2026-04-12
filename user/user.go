// Package user pour les infos user user/user.go
package user

import "errors"

// Erreurs exportées du package
var (
	ErrNotFound      = errors.New("user: utilisateur non trouvé")
	ErrInvalidEmail  = errors.New("user: email invalide")
	ErrDuplicateUser = errors.New("user: utilisateur déjà existant")
)
