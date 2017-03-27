// Author houguofa
// Copyright @2017 houguofa. All Rights Reserved.

package goconfiger

import (
	"sync"
)

var (
	defaultConfiger *configer     = nil
	singletonLock   *sync.RWMutex = new(sync.RWMutex)
)

func installConfiger(cfg *configer) {

	singletonLock.Lock()
	defer singletonLock.Unlock()
	defaultConfiger = cfg
}

func GetValueByString(section, key string) (string, error) {

	if nil == defaultConfiger {
		return "", defaultError
	}

	singletonLock.RLock()
	defer singletonLock.RUnlock()
	return defaultConfiger.GetValueByString(section, key)
}

func GetValueByBool(section, key string) (bool, error) {

	if nil == defaultConfiger {
		return false, defaultError
	}

	singletonLock.RLock()
	defer singletonLock.RUnlock()
	return defaultConfiger.GetValueByBool(section, key)
}

func GetValueByFloat64(section, key string) (float64, error) {

	if nil == defaultConfiger {
		return 0, defaultError
	}

	singletonLock.RLock()
	defer singletonLock.RUnlock()
	return defaultConfiger.GetValueByFloat64(section, key)
}

func GetValueByInt(section, key string) (int, error) {
	if nil == defaultConfiger {
		return 0, defaultError
	}

	singletonLock.RLock()
	defer singletonLock.RUnlock()
	return defaultConfiger.GetValueByInt(section, key)
}

func GetValueByInt64(section, key string) (int64, error) {
	if nil == defaultConfiger {
		return 0, defaultError
	}

	singletonLock.RLock()
	defer singletonLock.RUnlock()
	return defaultConfiger.GetValueByInt64(section, key)
}

func GetValueByStringList(section, key string) ([]string, error) {
	if nil == defaultConfiger {
		return nil, defaultError
	}

	singletonLock.RLock()
	defer singletonLock.RUnlock()
	return defaultConfiger.GetValueByStringList(section, key)
}
