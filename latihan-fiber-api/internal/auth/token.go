package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Klaim struct {
	Peran string `json:"peran"`
	jwt.RegisteredClaims
}

func TerbitkanToken(
	idPengguna uint,
	peran string,
	rahasia string,
	ttlMenit int,
) (string, error) {

	acak := make([]byte, 16)
	if _, err := rand.Read(acak); err != nil {
		return "", err
	}

	sekarang := time.Now()

	klaim := Klaim{
		Peran: peran,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       hex.EncodeToString(acak),
			Subject:  strconv.FormatUint(uint64(idPengguna), 10),
			IssuedAt: jwt.NewNumericDate(sekarang),
			ExpiresAt: jwt.NewNumericDate(
				sekarang.Add(time.Duration(ttlMenit) * time.Minute),
			),
			Issuer: "pemweb2-fiber-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, klaim)

	return token.SignedString([]byte(rahasia))
}

func PeriksaToken(
	nilaiToken string,
	rahasia string,
) (*Klaim, error) {

	token, err := jwt.ParseWithClaims(
		nilaiToken,
		&Klaim{},
		func(t *jwt.Token) (any, error) {
			return []byte(rahasia), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuer("pemweb2-fiber-api"),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return nil, err
	}

	klaim, cocok := token.Claims.(*Klaim)

	if !cocok || !token.Valid || klaim.ID == "" {
		return nil, errors.New("token tidak valid")
	}

	return klaim, nil
}
