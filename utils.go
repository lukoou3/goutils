package goutils

import (
	"fmt"
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
