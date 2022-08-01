package task3

import (
	"fmt"
	"regexp"
	"strconv"
)

const MAX_TIME = 1_000_000_000
const MIN_TIME = 0

const (
	END    = "end"
	SET    = "set"
	GET    = "get"
	CANCEL = "cancel"
)

type CacheWithFutures interface {
	Set(t1 int, t2 int, k string, v string) bool
	Get(t int, k string) (string, bool)
	Cancel(t int) bool
}

type cacheWithFutures struct {
	timeLine map[int]map[string]string
	store    map[string]string
}

type operation struct {
	name          string
	callTime      int
	operationTime int
	args          []string
}

func CreateCacheWithFutures() *cacheWithFutures {
	return &cacheWithFutures{
		timeLine: make(map[int]map[string]string),
		store:    make(map[string]string),
	}
}

func CreateOperation(name string, callTime int, operationTime int, args []string) *operation {
	return &operation{
		name:          name,
		callTime:      callTime,
		operationTime: operationTime,
		args:          args,
	}
}

func (cache *cacheWithFutures) Set(callTime int, setTime int, k string, v string) (string, bool) {
	_, ok := cache.timeLine[setTime]
	if !ok {
		cache.timeLine[setTime] = make(map[string]string)
	}

	cache.timeLine[setTime][k] = v

	_, ok = cache.store[k]
	if !ok {
		return "", false
	}

	if callTime >= setTime {
		cache.store[k] = v
	}

	return k, true
}

func (cache *cacheWithFutures) Get(callTime int, k string) (string, bool) {
	m, ok := cache.timeLine[callTime]
	if !ok {
		return "", false
	}
	v, ok := m[k]
	if !ok {
		return "", false
	}
	return v, true
}

func (cache *cacheWithFutures) Cancel(callTime int, cancelTime int) bool {
	_, ok := cache.timeLine[cancelTime]
	if !ok {
		return false
	}
	if callTime < cancelTime {
		return false
	}
	delete((*cache).timeLine, cancelTime)
	return true
}

func parse(input []string) ([]operation, error) {
	var result []operation

	maxCallTime := -1

	for _, event := range input {
		re := regexp.MustCompile("\t")
		split := re.Split(event, -1)

		if split[0] == "-1" {
			result = append(result, *CreateOperation(END, maxCallTime, maxCallTime, []string{}))
			break
		}

		callTime, err := strconv.Atoi(split[0])
		if err != nil {
			return result, err
		}

		if callTime > maxCallTime {
			maxCallTime = callTime
		}

		operation := split[1]

		switch operation {
		case "set":
			operationTime, err := strconv.Atoi(split[2])
			if err != nil {
				return result, err
			}
			k := split[3]
			v := split[4]
			result = append(result, *CreateOperation(SET, callTime, operationTime, []string{k, v}))
		case "get":
			k := split[2]
			result = append(result, *CreateOperation(GET, callTime, callTime, []string{k}))
		case "cancel":
			operationTime, err := strconv.Atoi(split[2])
			if err != nil {
				return result, err
			}
			result = append(result, *CreateOperation(CANCEL, callTime, operationTime, []string{}))
		default:
			return result, fmt.Errorf("unkown operation: %v", operation)
		}
	}

	return result, nil
}

func invoke(operations []operation) (string, error) {
	var result string
	cache := CreateCacheWithFutures()

	t := MIN_TIME

	fmt.Println("--------------------------")
	fmt.Printf("operations:%v\n", operations)
	fmt.Println("--------------------------")

	for t <= MAX_TIME {
		op := operations[t]
		fmt.Println("--------------------------")
		fmt.Printf("invoke:%v\n", op)
		fmt.Println("--------------------------")
		switch op.name {
		case END:
			return result, nil
		case SET:
			k, ok := cache.Set(op.callTime, op.operationTime, op.args[0], op.args[1])
			result += strconv.FormatBool(ok)
			if ok {
				result += "\t" + k
			}
		case GET:
			k, ok := cache.Get(op.callTime, op.args[0])
			result += strconv.FormatBool(ok)
			if ok {
				result += "\t" + k
			}
		case CANCEL:
			ok := cache.Cancel(op.callTime, op.operationTime)
			result += strconv.FormatBool(ok)
		default:
			return result, fmt.Errorf("unknown operation: %v", op.name)
		}

		result += "\n"

		t += 1
	}
	return result, nil
}

func Solution(input []string) (string, error) {
	var result string

	operations, err := parse(input)
	if err != nil {
		return result, err
	}

	result, err = invoke(operations)
	if err != nil {
		return result, err
	}

	return result, nil
}
