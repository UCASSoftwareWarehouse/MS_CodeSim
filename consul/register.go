package consul

import (
	"code_sim/config"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"time"
)

const (
	retryLimit = 30 * time.Second
)

// NewConsulRegister create a new consul register
func NewConsulRegister() *Register {
	return &Register{
		Address:                        config.Conf.ConsulAddr, //consul address
		Name:                           config.Conf.AppName,
		Tag:                            []string{},
		Port:                           config.Conf.Port,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
}

func MustRegisterGRPCServer(server *grpc.Server) {
	if config.Conf.GetEnv() == config.DevEnv {
		// dev do not use consul
		return
	}
	r := NewConsulRegister()
	for {
		start := time.Now()
		err := r.Register()
		if err == nil {
			break
		}
		if time.Now().Sub(start) >= retryLimit {
			panic("MustRegisterGRPCServer Failed to register")
		}
		time.Sleep(time.Second)
	}
	grpc_health_v1.RegisterHealthServer(server, &HealthImpl{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}

// Register consul service register
type Register struct {
	Address                        string
	Name                           string
	Tag                            []string
	Port                           int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

// Register register service
func (r *Register) Register() error {
	conf := api.DefaultConfig()
	conf.Address = r.Address
	client, err := api.NewClient(conf)
	if err != nil {
		return err
	}
	agent := client.Agent()

	IP := LocalIP()
	log.Printf("get local IP %s", IP)
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.Name, IP, r.Port), // 服务节点的名称
		Name:    r.Name,                                      // 服务名称
		Tags:    r.Tag,                                       // tag，可以为空
		Port:    r.Port,                                      // 服务端口
		Address: IP,                                          // 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       r.Interval.String(),                         // 健康检查间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Name), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),   // 注销时间，相当于过期时间
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}

func LocalIP() string {
	intf, err := net.InterfaceByName(config.Conf.NetworkInterface)
	if err != nil {
		panic(fmt.Sprintf("cannot find network interface with name %s", config.Conf.NetworkInterface))
	}
	addrs, err := intf.Addrs()
	if err != nil {
		panic(err)
	}
	log.Printf("find addrs %+v", addrs)
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("No valid ip addr")
}
