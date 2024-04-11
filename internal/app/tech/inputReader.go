package tech

import (
	"bufio"
	"strconv"
	"strings"
	"surlit/internal/app/tech/errors"
	"surlit/internal/app/tech/utils"
	"surlit/internal/logic/models"
)

type InputReader interface {
	ReadUUID() (models.UUID, error)
	ReadString() (string, error)
}

type InputReaderBufio struct {
	reader *bufio.Reader
}

var _ InputReader = (*InputReaderBufio)(nil)

func NewInputReaderBufio(reader *bufio.Reader) *InputReaderBufio {
	return &InputReaderBufio{reader: reader}
}

func (i InputReaderBufio) ReadUUID() (models.UUID, error) {
	str, err := i.reader.ReadString('\n')
	if err != nil {
		return 0, errors.ErrWrongInput
	}
	num, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		return 0, errors.ErrWrongInput
	}
	logicUUID, err := utils.UUIDUIToLogic(num)
	if err != nil {
		return 0, errors.ErrWrongInput
	}
	return logicUUID, nil
}

func (i InputReaderBufio) ReadString() (string, error) {
	str, err := i.reader.ReadString('\n')
	if err != nil {
		return "", errors.ErrWrongInput
	}
	return strings.TrimSpace(str), nil
}
