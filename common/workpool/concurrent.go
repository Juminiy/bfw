package workpool

import (
	"errors"
	"runtime"
	"sync"
)

func RunFunctions(fns ...func() error) (bool, []error) {
	fnCnt := len(fns)
	if fnCnt == 0 {
		return false, nil
	}

	errs := make([]error, fnCnt)
	hasErrors := false
	wg := sync.WaitGroup{}
	wg.Add(fnCnt)

	for i, fn := range fns {
		idx := i
		f := fn
		go func() {
			defer wg.Done()
			err := f()
			errs[idx] = err
			if err != nil {
				hasErrors = true
			}
		}()
	}

	wg.Wait()

	return hasErrors, errs
}

func RunConcurrently(conn int, fn func() error) (bool, []error) {
	if conn == 0 {
		return false, nil
	}

	if conn == -1 {
		conn = runtime.GOMAXPROCS(-1)
	}

	errs := make([]error, conn)
	hasErrors := false
	wg := sync.WaitGroup{}
	wg.Add(conn)

	for i := 0; i < conn; i++ {
		idx := i
		go func() {
			defer wg.Done()
			err := fn()
			errs[idx] = err
			if err != nil {
				hasErrors = true
			}
		}()
	}

	wg.Wait()
	return hasErrors, errs
}

func FirstError(errs []error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func MergeErrors(errs []error) error {
	var message string

	for _, err := range errs {
		if err != nil {
			message += err.Error() + "  "
		}
	}

	return errors.New(message)
}
