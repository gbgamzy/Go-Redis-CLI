# GoLang Redis CLI Implementation

## Description
This project provides a lightweight Redis CLI implementation in GoLang, capable of handling core Redis commands. It serves as a simple and efficient tool for interacting with a Redis server using custom Go handlers.

## Features
- Custom handlers for key Redis commands:
  - `PING`: Test the connection to the server.
  - `SET`: Store a key-value pair.
  - `GET`: Retrieve the value of a key.
  - `HSET`: Set a field in a hash.
  - `HGET`: Retrieve a field from a hash.
  - `HGETALL`: Retrieve all fields and values from a hash.
  - `HDEL`: Delete a field from a hash.
- Designed for simplicity and extensibility.
- Works seamlessly with `redis-cli`.

<img width="445" alt="Screenshot 2024-12-07 at 4 58 42 PM" src="https://github.com/user-attachments/assets/1c28940a-0ca1-4ebd-a9a8-74a87937fae9">

<img width="437" alt="Screenshot 2024-12-07 at 4 53 27 PM" src="https://github.com/user-attachments/assets/fcb30018-2078-4ed9-b3b3-fa61bdc5a167">

<img width="445" alt="Screenshot 2024-12-07 at 4 58 42 PM" src="https://github.com/user-attachments/assets/3594dc6f-be43-47ef-9dac-a3c3891dc53b">

<img width="561" alt="Screenshot 2024-12-07 at 5 07 00 PM" src="https://github.com/user-attachments/assets/a6ae5132-139e-45c2-8987-9a2dadb96805">



## Installation

Clone the repository:

```git clone https://github.com/yourusername/redis-cli-go.git
Navigate to the project directory:
cd redis-cli-go
Install dependencies:
go mod tidy
Start the Go server:
go run *.go
```
Usage

Open a separate terminal and launch the official Redis CLI:
```
redis-cli
```

Use the following commands, which are supported by the Go server:
```
PING
```

Test the connection to the Redis server:

```

PING
Output: PONG
```

SET
Store a key-value pair:

```

SET mykey myvalue
Output: OK
```

GET
Retrieve the value for a given key:

```

GET mykey
Output: "myvalue"
```

HSET
Set a field in a hash:

```

HSET myhash field1 value1
Output: (integer) 1
```

HGET
Retrieve the value of a specific field in a hash:

```

HGET myhash field1
Output: "value1"
```

HGETALL
Retrieve all fields and values in a hash:

```

HGETALL myhash
Output:

1) "field1"
2) "value1"
```


HDEL
Delete a specific field in a hash:

```

HDEL myhash field1
Output: (integer) 1
```


Demo


Here’s a visual demonstration of the commands in action:

Prerequisites

GoLang (v1.18 or higher)
Redis installed locally or access to a Redis server
License

This project is licensed under the MIT License.

