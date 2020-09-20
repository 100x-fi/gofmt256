package gofmt256

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagName      = "gofmt256"
	tagSep       = ","
	subTagAssign = "="
)

type Builder interface {
	Build() (string, error)
}

type file struct {
	header interface{}
	body   interface{}
	footer interface{}
}

func New(header, body, footer interface{}) Builder {
	return &file{
		header: header,
		body:   body,
		footer: footer,
	}
}

func (f *file) Build() (string, error) {
	fmt256 := ""

	headerValue := reflect.ValueOf(f.header)
	if headerValue.Kind() != reflect.Struct {
		return "", errors.New("header must be struct")
	}

	bodyValue := reflect.ValueOf(f.body)
	if bodyValue.Kind() != reflect.Slice {
		return "", errors.New("input must be a slice")
	}

	footerValue := reflect.ValueOf(f.footer)
	if footerValue.Kind() != reflect.Struct {
		return "", errors.New("footer must be struct")
	}

	headerLine, err := makeLine(headerValue)
	if err != nil {
		return "", err
	}
	fmt256 = fmt256 + headerLine

	sliceLen := bodyValue.Len()
	lines := ""
	for i := 0; i < sliceLen; i++ {
		line, err := makeLine(bodyValue.Index(i))
		if err != nil {
			return "", err
		}
		lines = lines + line
	}
	fmt256 = fmt256 + lines

	footerLine, err := makeLine(footerValue)
	if err != nil {
		return "", err
	}
	fmt256 = fmt256 + footerLine

	return fmt256, nil
}

type FieldStruct struct {
	Name    string
	Data    string
	from    int
	to      int
	align   string
	padding string
}

func makeLine(input reflect.Value) (line string, err error) {
	line = ""

	fieldStructs := make(map[string]FieldStruct)
	for i := 0; i < input.Type().NumField(); i++ {
		field := input.Type().Field(i)
		errLocation := "[" + field.Name + "] %s"
		tag := field.Tag.Get(tagName)
		subtags := strings.Split(tag, tagSep)

		from, to, align, padding, err := extractSubTag(subtags)
		if err != nil {
			return "", errors.Wrapf(err, errLocation, "unable to extract subtags")
		}
		if from < 0 || to < 0 {
			return "", errors.New(fmt.Sprintf(errLocation, "from or to is missing from subtag or the provided value is minus"))
		}
		if from > to {
			return "", errors.New(fmt.Sprintf(errLocation, "from must less than to"))
		}
		if from > 256 || to > 256 {
			return "", errors.New(fmt.Sprintf(errLocation, "from and to must less than 256"))
		}

		fieldStructs[field.Name] = FieldStruct{
			Name:    field.Name,
			Data:    fmt.Sprint(input.Field(i).Interface()),
			from:    from,
			to:      to,
			align:   align,
			padding: padding,
		}
	}

	sortedFieldStructs, err := sort(fieldStructs)
	if err != nil {
		return "", errors.Wrap(err, "failed to validate slot in 256 length")
	}
	for _, fs := range sortedFieldStructs {
		subline, err := pad(fs)
		if err != nil {
			return "", errors.Wrap(err, "failed to pad data")
		}
		line = line + subline
	}

	line = line + "\n"
	return line, nil
}

func pad(fs FieldStruct) (string, error) {
	padData := fs.Data
	length := (fs.to - fs.from) + 1
	if len(fs.Data) > length {
		return "", errors.New("data is longer than length")
	}
	toPad := length - len(fs.Data)
	for i := 0; i < toPad; i++ {
		if fs.align == "L" {
			padData = padData + fs.padding
		}
		if fs.align == "R" {
			padData = fs.padding + padData
		}
	}
	return padData, nil
}

func sort(mapFs map[string]FieldStruct) ([]FieldStruct, error) {
	slots := [256]string{}
	var sortedFieldStructs []FieldStruct
	for _, fs := range mapFs {
		for i := fs.from; i <= fs.to; i++ {
			if slots[i-1] != "" {
				return nil, errors.New("slot conflict between fields")
			}
			slots[i-1] = fs.Name
		}
	}
	isAllocAll := true
	for i := 1; i < len(slots); i++ {
		if slots[i] == "" {
			isAllocAll = false
		}
		if slots[i] != slots[i-1] {
			sortedFieldStructs = append(sortedFieldStructs, mapFs[slots[i-1]])
		}
	}
	if !isAllocAll {
		return nil, errors.New("from to not fill 256 bytes")
	}
	sortedFieldStructs = append(sortedFieldStructs, mapFs[slots[len(slots)-1]])

	return sortedFieldStructs, nil
}

func extractSubTag(subtags []string) (from int, to int, align string, padding string, err error) {
	from = -1
	to = -1
	align = "L"
	padding = " "
	for _, subtag := range subtags {
		splitedSubTag := strings.Split(subtag, subTagAssign)
		if len(splitedSubTag) != 2 {
			return -1, -1, "", "", errors.New("malformat for value within a gofmt256 tag")
		}
		if splitedSubTag[1] == "" {
			return -1, -1, "", "", errors.New("given sub tag doesn't has an right hand value")
		}
		switch splitedSubTag[0] {
		case "from":
			from, err = strconv.Atoi(splitedSubTag[1])
			if err != nil {
				return -1, -1, "", "", errors.Wrap(err, "unable to convert `from` to `int`")
			}
		case "to":
			to, err = strconv.Atoi(splitedSubTag[1])
			if err != nil {
				return -1, -1, "", "", errors.Wrap(err, "unable to covert `to` to int")
			}
		case "align":
			align = splitedSubTag[1]
		case "padding":
			padding = strings.ReplaceAll(splitedSubTag[1], "'", "")
		}
	}
	return from, to, align, padding, nil
}
