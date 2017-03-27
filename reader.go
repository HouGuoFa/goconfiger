// Author houguofa
// Copyright @2017 houguofa. All Rights Reserved.

package goconfiger

import (
	"errors"
	"strings"
)

func LoadConfig(file string) error {

	cfg := constructerConfiger(file)
	err := cfg.readFile()
	if nil != err {
		return err
	}

	installConfiger(cfg)
	return err
}

func LoadAndGetConfig(file string) (*configer, error) {

	cfg := constructerConfiger(file)
	err := cfg.readFile()

	return cfg, err
}

func parseSingleLine(line string) (string, string, error) {

	strs := strings.Split(line, confDelimiter)
	if confLineCount != len(strs) {
		return "", "", errors.New("config " + line + " parse failed")
	}

	return strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1]), nil
}
