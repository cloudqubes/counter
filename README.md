# Overview
`counter` is a simple Go lang application that counts up or down based on HTTP request.

We use `counter` for demonstrating and testing CI/CD pipelines.

# Get started

Clone the repo and run the server.
```shell
git clone git@github.com:cloudqubes/counter.git
cd counter
go run counter.go
```

# How counter works

`counter` application expose two URL endpoints.

## count-up

Path: `/count-up`
Increment the counter by 1.

## count-down

Path: `/count-down`
Reduce counter by 1.


# Testing

`counter v1`` does not include testing.

# Using Jenkins

You can build `counter 1.0.0` Docker container with Jenkins. Create a Jenkins Pipeline and run in a node that has Docker installed.