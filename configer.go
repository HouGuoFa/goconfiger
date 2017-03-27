// Author houguofa
// Copyright @2017 houguofa. All Rights Reserved.

package goconfiger

import (
	//"errors"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
*
* # 注释
* [section]
* key1 = val
* key2 = val
*
 */

type sectionUnit struct {
	comment     string
	sectionName string
	keyNameList []string
	confData    map[string][]string
}

func newSection(name, comments string) *sectionUnit {
	return &sectionUnit{
		comment:     comments,
		sectionName: name,
		keyNameList: make([]string, 0, 10),
		confData:    make(map[string][]string)}
}

type configer struct {
	confFileName    string
	sectionNameList []string
	sectionData     map[string]*sectionUnit
}

func newConfiger(file string) (*configer, error) {
	return constructerConfiger(file), nil
}

//func (cfg *configer) Show() {

//	if cfg == nil {
//		return
//	}

//	fmt.Println(cfg.confFileName)
//	fmt.Println("----------------------------------------------------------------")
//	for sec, value := range cfg.sectionData {
//		fmt.Println(value.comment)
//		fmt.Println("[", sec, "]")

//		for k, values := range value.confData {
//			fmt.Println(values[0])
//			for i := 1; i < len(values); i++ {
//				fmt.Println(k, " = ", values[i])
//			}
//		}
//	}
//	fmt.Println("----------------------------------------------------------------")
//}

func constructerConfiger(file string) *configer {

	cfg := &configer{confFileName: file,
		sectionNameList: make([]string, 0, 10),
		sectionData:     make(map[string]*sectionUnit)}
	return cfg
}

func (cfg *configer) readFile() error {

	fd, err := os.Open(cfg.confFileName)
	if err != nil {
		return err
	}

	defer fd.Close()

	return cfg.read(fd)
}

func (cfg *configer) read(fd io.Reader) error {

	reader := bufio.NewReader(fd)

	var commentsTmp string
	var sectionTmp string

	for {

		line, err := reader.ReadString('\n')
		if nil != err {
			if io.EOF == err {
				return nil
			}
			return err
		}

		line = strings.TrimSpace(line)

		if 0 == len(line) {
			continue
		}

		switch line[0] {
		case '[':

			if line[len(line)-1] != ']' {
				return errors.New("section format lost")
			}

			sectionTmp = line[1 : len(line)-1]
			if err := cfg.setSection(sectionTmp, commentsTmp); nil != err {
				return err
			}
			commentsTmp = ""

		case ';':
			fallthrough
		case '#':
			commentsTmp = fmt.Sprintf("%s%s%s", commentsTmp, lineDelimiter, line)

		default:

			key, value, err := parseSingleLine(line)
			if nil != err {
				return err
			}
			if err := cfg.setValue(sectionTmp, key, value, commentsTmp); nil != err {
				return err
			}
			commentsTmp = ""
		}

	}
}

func (cfg *configer) setSection(section, comment string) error {

	sec, exist := cfg.sectionData[section]
	if exist {
		return errors.New("can't load multi same secion " + section)
	}

	if nil == sec {
		sec = newSection(section, comment)
	}

	cfg.sectionNameList = append(cfg.sectionNameList, section)
	cfg.sectionData[section] = sec
	return nil
}

func (cfg *configer) setValue(section, key, value, comment string) error {

	if 0 == len(section) || 0 == len(key) || 0 == len(value) {
		return errors.New("set value failed")
	}

	sec, exist := cfg.sectionData[section]
	if false == exist {
		return errors.New("not found section[" + section + "]for set value")
	}

	//sec.comment = comment
	if len(sec.confData[key]) == 0 {
		sec.confData[key] = append(sec.confData[key], comment)
		sec.keyNameList = append(sec.keyNameList, key)
	}

	sec.confData[key] = append(sec.confData[key], value)
	cfg.sectionData[section] = sec
	return nil
}

func (cfg *configer) getValue(section, key string) ([]string, error) {

	sec := cfg.sectionData[section]
	if nil == sec {
		return nil, errors.New(p_SECTION_NOT_FOUND)
	}

	if values := sec.confData[key]; len(values) > 0 {
		return values, nil
	}

	return nil, errors.New(p_KEY_NOT_FOUND)
}

func (cfg *configer) GetValueByString(section, key string) (string, error) {

	values, err := cfg.getValue(section, key)
	if err != nil {
		return "", err
	}

	return values[1], nil
}

func (cfg *configer) GetValueByBool(section, key string) (bool, error) {

	values, err := cfg.getValue(section, key)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(values[1])
}

func (cfg *configer) GetValueByFloat64(section, key string) (float64, error) {
	values, err := cfg.getValue(section, key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(values[1], 64)
}

func (cfg *configer) GetValueByInt(section, key string) (int, error) {
	values, err := cfg.getValue(section, key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(values[1])
}

func (cfg *configer) GetValueByInt64(section, key string) (int64, error) {
	values, err := cfg.getValue(section, key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(values[1], 10, 64)
}

func (cfg *configer) GetValueByStringList(section, key string) ([]string, error) {
	values, err := cfg.getValue(section, key)
	if err != nil {
		return nil, err
	}
	return values[1:], nil
}
