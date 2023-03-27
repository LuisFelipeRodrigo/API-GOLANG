package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `jason:"id,omitempty"`
	Nome     string    `jason:"nome,omitempty"`
	Nick     string    `jason:"nick,omitempty"`
	Email    string    `jason:"email,omitempty"`
	Senha    string    `jason:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar usuario
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o email não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("a senha não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}
