package xgen

import "common-toolkits-v1/XGen/GormXGens/config"

type CreateXGenCustom func(*ModelObject, config.XGenConfig) (any, any, error)
