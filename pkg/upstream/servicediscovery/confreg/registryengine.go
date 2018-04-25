package registry

import (
    "gitlab.alipay-inc.com/afe/mosn/pkg/upstream/servicediscovery/confreg/config"
    "gitlab.alipay-inc.com/afe/mosn/pkg/upstream/servicediscovery/confreg/servermanager"
    "sync"
)

var confregServerManager *servermanager.RegistryServerManager
var registryClient Client

var lock = new(sync.Mutex)

var ModuleStarted = false

func StartupRegistryModule(sysConfig *config.SystemConfig, registryConfig *config.RegistryConfig) Client {
    lock.Lock()

    defer func() {
        lock.Unlock()
    }()

    if ModuleStarted {
        return registryClient
    }
    confregServerManager = servermanager.NewRegistryServerManager(sysConfig, registryConfig)

    ModuleStarted = true

    return NewConfregClient(sysConfig, registryConfig, confregServerManager)
}