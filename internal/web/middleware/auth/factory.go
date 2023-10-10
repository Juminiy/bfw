package auth

import (
	"bfw/internal/logger"
	"errors"
)

var (
	_method        AuthMethod    = AuthNone
	_inst          AuthInterface = &BaseAuth{}
	WebAuthEnabled bool          = false
)

func Get() AuthInterface {
	return _inst
}

func SetAuthMethod(method AuthMethod) error {
	_method = method
	if method > AuthNone {
		WebAuthEnabled = true
	}
	var err error
	if _inst, err = ProduceAuth(); err != nil {
		logger.Errorf("error inner producing auth: %v", err)
		return err
	}

	return nil
}

func ProduceAuth() (AuthInterface, error) {
	return InstantiateAuth(_method)
}

func InstantiateAuth(method AuthMethod) (AuthInterface, error) {
	switch method {
	case AuthJwt:
		return &JwtAuth{}, nil
	default:
		logger.Errorf("unimplemented auth: %d", method)
		break
	}
	return nil, errors.New("has not implemented")
}
