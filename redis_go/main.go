package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// RedisClient 结构体，封装 TCP 连接
type RedisClient struct {
	conn net.Conn
}

// NewRedisClient 创建一个新的 Redis 客户端
func NewRedisClient(address string) (*RedisClient, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &RedisClient{conn: conn}, nil
}

// SendCommand 发送命令到 Redis 服务器
func (c *RedisClient) SendCommand(command ...string) (string, error) {
	// 按照 RESP 协议格式化命令
	cmd := fmt.Sprintf("*%d\r\n", len(command))
	for _, arg := range command {
		cmd += fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	}

	// 发送命令
	_, err := c.conn.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	// 读取服务器响应
	return c.readResponse()
}

// readResponse 读取并解析 Redis 服务器的响应
func (c *RedisClient) readResponse() (string, error) {
	reader := bufio.NewReader(c.conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// 根据 Redis RESP 协议解析响应
	switch response[0] {
	case '+': // 简单字符串
		return strings.TrimSpace(response[1:]), nil
	case '-': // 错误信息
		return "", fmt.Errorf("Redis error: %s", strings.TrimSpace(response[1:]))
	case ':': // 整数
		return strings.TrimSpace(response[1:]), nil
	case '$': // 批量字符串
		length := 0
		fmt.Sscanf(response, "$%d\r\n", &length)
		if length == -1 {
			return "", nil // Redis 返回空值（nil）
		}
		buf := make([]byte, length+2)
		_, err := reader.Read(buf)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(buf[:length])), nil
	default:
		return "", fmt.Errorf("unknown response: %s", response)
	}
}

// Auth 发送 AUTH 命令
func (c *RedisClient) Auth(password string) error {
	response, err := c.SendCommand("AUTH", password)
	if err != nil {
		return err
	}
	fmt.Println("AUTH Response:", response)
	return nil
}

// Set 发送 SET 命令
func (c *RedisClient) Set(key, value string) error {
	response, err := c.SendCommand("SET", key, value)
	if err != nil {
		return err
	}
	fmt.Println("SET Response:", response)
	return nil
}

// Get 发送 GET 命令
func (c *RedisClient) Get(key string) (string, error) {
	response, err := c.SendCommand("GET", key)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Close 关闭连接
func (c *RedisClient) Close() {
	c.conn.Close()
}

func main() {
	// 创建 Redis 客户端
	client, err := NewRedisClient("127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接 Redis 失败:", err)
		return
	}
	defer client.Close()

	// 如果 Redis 设置了密码，需要先 AUTH
	// err = client.Auth("yourpassword")
	// if err != nil {
	// 	fmt.Println("AUTH 失败:", err)
	// 	return
	// }

	// 发送 SET 命令
	err = client.Set("name", "Alice")
	if err != nil {
		fmt.Println("SET 失败:", err)
		return
	}

	// 发送 GET 命令
	value, err := client.Get("name")
	if err != nil {
		fmt.Println("GET 失败:", err)
		return
	}
	fmt.Println("GET name:", value)
}
