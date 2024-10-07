package functionswithneeds

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// Пакет нужен для удобства организации параллельного и последовательного выполнения функций.
//
// Function - функция, которая должна быть выполнена.
// Needs - функции, которые должны быть выполнены до того, как Function начнет выполняться.
//
// Все функции, которые должны быть выполнены и есть в Needs, должны быть указаны.

type FunctionWithNeeds struct {
	Function func(ctx context.Context) error
	Needs    []func(ctx context.Context) error
}

type FunctionsWithNeeds []FunctionWithNeeds

func Start(ctx context.Context, functions FunctionsWithNeeds) error {
	if err := validate(functions); err != nil {
		return err
	}

	return start(ctx, functions)
}

func start(ctx context.Context, functions FunctionsWithNeeds) error {
	if len(functions) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	newCtx, cancelCause := context.WithCancelCause(ctx)
	notifyCompleteFunctionChs := make([]chan string, len(functions))

	for i := range notifyCompleteFunctionChs {
		notifyCompleteFunctionChs[i] = make(chan string, len(functions))
	}

	for i, function := range functions {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			executeFunction(
				newCtx,
				cancelCause,
				function,
				notifyCompleteFunctionChs[i],
				notifyCompleteFunctionChs,
			)
		}()
	}

	wg.Wait()

	return context.Cause(newCtx)
}

func executeFunction(
	ctx context.Context,
	cancelCause context.CancelCauseFunc,
	function FunctionWithNeeds,
	receiveCompletedOtherFnCh chan string,
	allChs []chan string,
) {
	currentFunctionPointer := fmt.Sprintf("%p", function.Function)
	defer broadcastMessage(allChs, currentFunctionPointer)

	needsPointers := map[string]struct{}{}

	for _, need := range function.Needs {
		needsPointers[fmt.Sprintf("%p", need)] = struct{}{}
	}

	for {
		if len(needsPointers) == 0 {
			err := function.Function(ctx)

			if err != nil {
				cancelCause(err)
			}

			break
		}

		select {
		case <-ctx.Done():
			return
		case completedFnPtr := <-receiveCompletedOtherFnCh:
			delete(needsPointers, completedFnPtr)
		}
	}
}

func broadcastMessage(chs []chan string, message string) {
	for _, ch := range chs {
		ch <- message
	}
}

func validate(functions FunctionsWithNeeds) error {
	if len(functions) == 0 {
		return nil
	}

	functionCanExecutePointers := map[string]bool{}

	for _, function := range functions {
		if len(function.Needs) == 0 {
			functionCanExecutePointers[fmt.Sprintf("%p", function.Function)] = true
		}
	}

	if len(functionCanExecutePointers) == 0 {
		return errors.New("validation failed, check functions needs")
	}

	filteredFunctions := FunctionsWithNeeds{}

	for _, function := range functions {
		if functionCanExecutePointers[fmt.Sprintf("%p", function.Function)] {
			continue
		}

		functionToAdd := FunctionWithNeeds{}
		functionToAdd.Function = function.Function

		for _, need := range function.Needs {
			if functionCanExecutePointers[fmt.Sprintf("%p", need)] {
				continue
			}

			functionToAdd.Needs = append(functionToAdd.Needs, need)
		}

		filteredFunctions = append(filteredFunctions, functionToAdd)
	}

	return validate(filteredFunctions)
}
