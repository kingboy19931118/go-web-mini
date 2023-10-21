package util

import "time"

func DoRetry(f func() error) error {
	var err error
	for i := 0; i < 3; i++ {
		err = f()
		if err == nil {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return err
}
