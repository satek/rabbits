# Rabbits
Practice using RabbitMQ with Go language and managing worker queues

Based on a very nice tutorial found [here](http://nesv.github.io/golang/2014/02/25/worker-queues-in-go.html).
The tutorial implements a web interface for creating work that is then consumed by workers. This is extended here so that
work can be scheduled also by a RabbitMQ queue.

There is also one change in how work is being scheduled in `worker.go`. In the tutorial [dispather](https://gist.github.com/nesv/9233300#file-dispatcher-go-L20)
a new goroutine is started in the background every time a new job is found and that can lead to too many worker routines
being created if workers are not processing fast enough and the process can run out of memory. Instead, [worker](https://github.com/satek/rabbits/blob/master/worker_queues/worker.go#L26)
spawns only one goroutine in which it fetches work from the WorkQueue and only gets a new message after the first one was processed.
That ensures that work stays in Rabbit queues if it is not processed fast enough and we never have more routines than the number
we set when starting the process (4 by default). We could also scale number of workers up and down based on how much work
there is in queues.


## Usage

Only one package needs to be installed, RabbitMQ (AMQP) client. After setting up the GOPATH environment variable run:

`go get github.com/streadway/amqp`

Compile two binaries:

`go build rabbits/worker_queues` and `go build rabbits/produce`

and run each in a separate terminal.
