package key

import uuid "github.com/nu7hatch/gouuid"

func GenerateKey(subject string) string {
	n, _ := uuid.NewV4()
	u, _ := uuid.NewV5(n, []byte(subject))
	return u.String()
}
