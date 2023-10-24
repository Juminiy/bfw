package utils

import "bfw/wheel/adt"

// utils Stand for utilities

type ApplicationLevelVirtualStack adt.GenericStack[interface{}]

type ApplicationLevelVirtualQueue adt.GenericQueue[interface{}]

type CallbackFunctionPointerList adt.GenericList[interface{}]
