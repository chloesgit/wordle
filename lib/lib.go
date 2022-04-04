package lib

import (
	"errors"
	"fmt"
	"strings"
)

type InfoLettre string

const (
	infoBienPlace InfoLettre = "green"
	infoMalPlace  InfoLettre = "orange"
	infoExistePas InfoLettre = "red"
)

func Isword(motcible string, motsaisi string) ([]InfoLettre, error) {
	motcible = strings.ToLower(motcible)
	motsaisi = strings.ToLower(motsaisi)
	lettrescorrectes := make([]InfoLettre, len(motsaisi))

	if len(motcible) != len(motsaisi) {
		return lettrescorrectes, errors.New("le mot entré n'a pas la bonne longueur")
	}
	for i, _ := range motsaisi {
		lettrescorrectes[i] = infoExistePas
		if motsaisi[i] == motcible[i] {
			lettrescorrectes[i] = infoBienPlace
			continue
		}
		for j, _ := range motcible {
			if motsaisi[i] == motcible[j] {
				lettrescorrectes[i] = infoMalPlace
				break
			}
		}
		fmt.Println(lettrescorrectes)
	}

	return lettrescorrectes, nil
}
