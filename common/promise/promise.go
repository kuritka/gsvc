//Provides promise. All you need to do is find non-blocking function
//and wrap it to another function returning *Promise
//see the details https://github.com/kuritka/threading/blob/master/c_promises/main.go
package promise

import (
	"errors"
	"time"
)

type Promise struct {
	SuccessChannel chan interface {}
	ErrorChannel   chan error
}


func (p *Promise) Then(success func(interface{}) error, failure func(error)) *Promise{
	result := new(Promise)

	//buffer must be set at least to one because it could take sometime until someone drain the value
	result.SuccessChannel = make(chan interface{},1)
	result.ErrorChannel = make(chan error,1)

	timeout := time.After(5*time.Second)
	go func(){
		select {
		case obj := <- p.SuccessChannel:
			newErr := success(obj)
			if newErr == nil {
				result.SuccessChannel <- obj
			} else {
				result.ErrorChannel <- newErr
			}
		case err := <- p.ErrorChannel:
			failure(err)
			result.ErrorChannel <-err
		case <- timeout :
			panic(errors.New("timeout occurred"))
		}
	}()

	return result
}