module go_thinknetica/hw1/fib/cmd/fib

go 1.15

replace (
	go_thinknetica/hw1/fib/pkg/fib v0.0.0 => ../../pkg/fib
)

require (
	go_thinknetica/hw1/fib/pkg/fib v0.0.0
)