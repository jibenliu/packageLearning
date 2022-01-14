package wpool

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type ExecutionFn func(ctx context.Context, args interface{}) (interface{}, error)

type JobDescriptor struct {
	ID       int
	JType    string
	Metadata map[string]interface{}
}

type Result struct {
	Value      interface{}
	Err        error
	Descriptor JobDescriptor
}

type Job struct {
	Descriptor JobDescriptor
	ExecFn     ExecutionFn
	Args       interface{}
}

// 处理 job 逻辑,处理结果包装成 Result 结果
func (j Job) execute(ctx context.Context) Result {
	fmt.Printf("job is %v\n", j)
	if j.ExecFn != nil {
		value, err := j.ExecFn(ctx, j.Args)
		if err != nil {
			return Result{
				Err:        err,
				Descriptor: j.Descriptor,
			}
		}

		return Result{
			Value:      value,
			Descriptor: j.Descriptor,
		}
	}
	return Result{
		Err:        errors.New("ExecFn is not defined"),
		Descriptor: j.Descriptor,
	}
}

func GenJobs() []Job {
	return []Job{
		{
			Descriptor: JobDescriptor{
				ID:       1,
				JType:    "jType1",
				Metadata: map[string]interface{}{"job1": "sss"},
			},
			ExecFn: func(ctx context.Context, args interface{}) (interface{}, error) {
				time.Sleep(1 * time.Second)
				fmt.Println("job1 done")
				return 123456, nil
			},
			Args: "job1",
		},
		{
			Descriptor: JobDescriptor{
				ID:       2,
				JType:    "jType2",
				Metadata: map[string]interface{}{"job2": "ddd"},
			},
			ExecFn: func(ctx context.Context, args interface{}) (interface{}, error) {
				time.Sleep(3 * time.Second)
				fmt.Println("job2 done")
				return 234567, nil
			},
			Args: "job2",
		},
		{
			Descriptor: JobDescriptor{
				ID:       3,
				JType:    "jType3",
				Metadata: map[string]interface{}{"job3": "eee"},
			},
			Args: "job3",
		},
		{
			Descriptor: JobDescriptor{
				ID:       4,
				JType:    "jType4",
				Metadata: map[string]interface{}{"job4": "fff"},
			},
			ExecFn: func(ctx context.Context, args interface{}) (interface{}, error) {
				time.Sleep(5 * time.Second)
				fmt.Println("job4 done")
				return 456789, errors.New("something wrong")
			},
			Args: "job4",
		},
		{
			Descriptor: JobDescriptor{
				ID:       5,
				JType:    "jType5",
				Metadata: map[string]interface{}{"job5": "kkk"},
			},
			Args: "job5",
		},
	}
}
