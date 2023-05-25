package goutils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Retry 执行某个函数, 最多重试n次
func Retry[T any](n int, waitSleep time.Duration, f func() (T, error)) (T, error) {
	var rst T
	var err error
	for times := 1; times <= n; times++ {
		rst, err = f()
		if err == nil {
			break
		} else {
			if times < n {
				fmt.Printf("retry func failed for %d times\n", times)
				fmt.Println(err)
				if waitSleep > 0 {
					time.Sleep(waitSleep)
				}
			} else {
				fmt.Printf("retry func %d times finally failed \n", times)
			}
		}
	}
	return rst, err
}

// RetryVoid 执行某个函数, 最多重试n次
func RetryVoid(n int, waitSleep time.Duration, f func() error) error {
	var err error
	for times := 1; times <= n; times++ {
		err = f()
		if err == nil {
			break
		} else {
			if times < n {
				fmt.Printf("retry func failed for %d times", times)
				if waitSleep > 0 {
					time.Sleep(waitSleep)
				}
			} else {
				fmt.Printf("retry func %d times finally failed ", times)
			}
		}
	}
	return err
}

var (
	pyFmtKvRe  *regexp.Regexp
	pyFmtArgRe *regexp.Regexp
)

func init() {
	pyFmtKvRe = regexp.MustCompile(`\{(\w+)\}`)
	pyFmtArgRe = regexp.MustCompile(`\{(\d*)\}`)
}

func PyFormat(str string, args ...string) string {
	i := -1
	return pyFmtArgRe.ReplaceAllStringFunc(str, func(s string) string {
		// s = strings.Trim(s, "{}")
		s = s[1 : len(s)-1]
		i++
		if s == "" {
			return args[i]
		} else {
			index, _ := strconv.ParseInt(s, 10, 32)
			return args[index]
		}
	})
}

func PyFormat2(str string, args ...interface{}) string {
	i := -1
	return pyFmtArgRe.ReplaceAllStringFunc(str, func(s string) string {
		// s = strings.Trim(s, "{}")
		s = s[1 : len(s)-1]
		i++
		if s == "" {
			return InterfaceToString(args[i])
		} else {
			index, _ := strconv.ParseInt(s, 10, 32)
			return InterfaceToString(args[index])
		}
	})
}

func PyFormatKv(str string, kvs map[string]string) string {
	return pyFmtKvRe.ReplaceAllStringFunc(str, func(s string) string {
		// s = strings.Trim(s, "{}")
		k := s[1 : len(s)-1]
		if v, ok := kvs[k]; ok {
			return v
		} else {
			return k
		}
	})
}

func PyFormatKv2(str string, kvs map[string]interface{}) string {
	return pyFmtKvRe.ReplaceAllStringFunc(str, func(s string) string {
		// s = strings.Trim(s, "{}")
		k := s[1 : len(s)-1]
		if v, ok := kvs[k]; ok {
			return InterfaceToString(v)
		} else {
			return k
		}
	})
}

func InterfaceMapToString[T comparable](a map[T]interface{}) map[T]string {
	b := make(map[T]string)
	for key, value := range a {
		if value != nil {
			b[key] = InterfaceToString(value)
		}
	}
	return b
}

func InterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		//newValue, _ := json.Marshal(value)
		//key = string(newValue)
		key = fmt.Sprintf("%v", value)
	}

	return key
}
