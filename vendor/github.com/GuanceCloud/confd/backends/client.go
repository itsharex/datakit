package backends

import (
	"errors"

	"github.com/GuanceCloud/confd/backends/aws"
	"github.com/GuanceCloud/confd/backends/consul"
	"github.com/GuanceCloud/confd/backends/dynamodb"
	"github.com/GuanceCloud/confd/backends/env"
	"github.com/GuanceCloud/confd/backends/etcdv3"
	"github.com/GuanceCloud/confd/backends/file"
	"github.com/GuanceCloud/confd/backends/nacos"
	"github.com/GuanceCloud/confd/backends/rancher"
	"github.com/GuanceCloud/confd/backends/redis"
	"github.com/GuanceCloud/confd/backends/ssm"
	"github.com/GuanceCloud/confd/backends/vault"
	"github.com/GuanceCloud/confd/backends/zookeeper"
)

// The StoreClient interface is implemented by objects that can retrieve
// key/value pairs from a backend store.
type StoreClient interface {
	GetValues(keys []string) (map[string]string, error)
	WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error)
	Close()
}

// New is used to create a storage client based on our configuration.
func New(config Config) (StoreClient, error) {

	if config.Backend == "" {
		config.Backend = "etcd"
	}
	backendNodes := config.BackendNodes

	// if config.Backend == "file" {
	// 	log.Info("Backend source(s) set to " + strings.Join(config.YAMLFile, ", "))
	// } else {
	// 	log.Info("Backend source(s) set to " + strings.Join(backendNodes, ", "))
	// }

	switch config.Backend {
	case "consul":
		return consul.New(config.BackendNodes, config.Scheme,
			config.ClientCert, config.ClientKey,
			config.ClientCaKeys,
			config.BasicAuth,
			config.Username,
			config.Password,
		)
	case "etcd":
		// etcd v2 has been deprecated and etcdv3 is now the client for both the etcd and etcdv3 backends.
		return etcdv3.NewEtcdClient(backendNodes, config.ClientCert, config.ClientKey, config.ClientCaKeys, config.BasicAuth, config.Username, config.Password)
	case "etcdv3":
		return etcdv3.NewEtcdClient(backendNodes, config.ClientCert, config.ClientKey, config.ClientCaKeys, config.BasicAuth, config.Username, config.Password)
	case "zookeeper":
		return zookeeper.NewZookeeperClient(backendNodes)
	case "rancher":
		return rancher.NewRancherClient(backendNodes)
	case "redis":
		return redis.NewRedisClient(backendNodes, config.ClientKey, config.Separator)
	case "env":
		return env.NewEnvClient()
	case "file":
		return file.NewFileClient(config.YAMLFile, config.Filter)
	case "vault":
		vaultConfig := map[string]string{
			"app-id":    config.AppID,
			"user-id":   config.UserID,
			"role-id":   config.RoleID,
			"secret-id": config.SecretID,
			"username":  config.Username,
			"password":  config.Password,
			"token":     config.AuthToken,
			"cert":      config.ClientCert,
			"key":       config.ClientKey,
			"caCert":    config.ClientCaKeys,
			"path":      config.Path,
		}
		return vault.New(backendNodes[0], config.AuthType, vaultConfig)
	case "dynamodb":
		table := config.Table
		// log.Info("DynamoDB table set to " + table)
		return dynamodb.NewDynamoDBClient(table)
	case "ssm":
		return ssm.New()
	case "nacos":
		return nacos.NewNacosClient(
			config.BackendNodes,
			config.Password,
			config.Username,
			config.Namespace,
			config.AccessKey,
			config.SecretKey,
			config.CircleInterval,
		)
	case "aws":
		return aws.NewAWSClient(
			config.AccessKey,
			config.SecretKey,
			config.Region,
			config.CircleInterval,
		)

	}

	return nil, errors.New("invalid backend")
}
