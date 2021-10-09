package consul

import (
	"code_sim/config"
	"testing"
)

func TestRegister_Register(t *testing.T) {
	config.InitConfig(&config.cmdArgs{
		Env:        config.DevEnv,
		Port:       0,
		ConfigPath: "/Users/purchaser/go/src/code_sim/config.yml",
	})
	r := NewConsulRegister()
	err := r.Register()
	t.Logf("err=[%v]", err)
	// grpc_health_v1.RegisterHealthServer(svr.server, &microservice.HealthImpl{Status:grpc_health_v1.HealthCheckResponse_SERVING})

}
