package lang

import "unsafe"

// wheel assertion

func ValidateStringArrayOrSlice(arrayOrSlice []string) bool {
	return ValidateInterfaceArrayOrSlice(arrayOrSlice)
}

func ValidateStructArrayOrSlice(arrayOrSlice []struct{}) bool {
	return ValidateInterfaceArrayOrSlice(arrayOrSlice)
}

func ValidateInterfaceArrayOrSlice(infArrayOrSlice interface{}) bool {
	switch infArrayOrSlice.(type) {
	case []int:
		{
			if len(infArrayOrSlice.([]int)) != 0 {
				return true
			}
			return false
		}
	case []string:
		{
			if len(infArrayOrSlice.([]string)) != 0 {
				return true
			}
			return false
		}
	case []*struct{}:
		{
			if len(infArrayOrSlice.([]*struct{})) != 0 {
				return true
			}
			return false
		}
	case []struct{}:
		{
			if len(infArrayOrSlice.([]struct{})) != 0 {
				return true
			}
			return false
		}
	case []*interface{}:
		{
			if len(infArrayOrSlice.([]*interface{})) != 0 {
				return true
			}
			return false
		}
	case []interface{}:
		{
			if len(infArrayOrSlice.([]interface{})) != 0 {
				return true
			}
			return false
		}
	case map[int][]int:
		{
			if len(infArrayOrSlice.(map[int][]int)) != 0 {
				return true
			}
			return false
		}
	case map[int]*struct{}:
		{
			if len(infArrayOrSlice.(map[int]*struct{})) != 0 {
				return true
			}
			return false
		}
	case map[int]struct{}:
		{
			if len(infArrayOrSlice.(map[int]struct{})) != 0 {
				return true
			}
			return false
		}
	case map[int]interface{}:
		{
			if len(infArrayOrSlice.(map[int]interface{})) != 0 {
				return true
			}
			return false
		}
	case map[int]*interface{}:
		{
			if len(infArrayOrSlice.(map[int]*interface{})) != 0 {
				return true
			}
			return false
		}
	default:
		{
			if infArrayOrSlice != nil {
				return true
			}
		}
	}
	return false
}

func ConvertInterfaceArrayOrSliceToString(infArrayOrSlice interface{}) []string {
	switch infArrayOrSlice.(type) {
	case []string:
		{
			return infArrayOrSlice.([]string)
		}
	case []interface{}:
		{
			if ValidateInterfaceArrayOrSlice(infArrayOrSlice) {
				infArrayOrSliceInf := infArrayOrSlice.([]interface{})
				infArrayOrSliceString := make([]string, len(infArrayOrSliceInf))
				for idx, ele := range infArrayOrSliceInf {
					infArrayOrSliceString[idx] = ele.(string)
				}
				return infArrayOrSliceString
			}
			return nil
		}
	default:
		{

		}
	}
	return nil
}

func GetGenericsZeroValue[T any]() T {
	var t T
	return t
}

func ConvertInterfaceElementToString(element interface{}) string {
	if element != nil {
		return element.(string)
	}
	return undefinedString
}

func GetUnSafePointer(value any) any {
	return unsafe.Pointer(&value)
}
