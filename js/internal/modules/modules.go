/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2020 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package modules

import (
	"fmt"
	"sync"

	"github.com/loadimpact/k6/js/modules/k6"
	"github.com/loadimpact/k6/js/modules/k6/crypto"
	"github.com/loadimpact/k6/js/modules/k6/crypto/x509"
	"github.com/loadimpact/k6/js/modules/k6/data"
	"github.com/loadimpact/k6/js/modules/k6/encoding"
	"github.com/loadimpact/k6/js/modules/k6/grpc"
	"github.com/loadimpact/k6/js/modules/k6/html"
	"github.com/loadimpact/k6/js/modules/k6/http"
	"github.com/loadimpact/k6/js/modules/k6/metrics"
	"github.com/loadimpact/k6/js/modules/k6/ws"
)

//nolint:gochecknoglobals
var (
	modules = make(map[string]interface{})
	mx      sync.RWMutex
)

// HasModuleInstancePerVU should be implemented by all native Golang modules that
// would require per-VU state. k6 will call their NewModuleInstancePerVU() methods
// every time a VU imports the module and use its result as the returned object.
type HasModuleInstancePerVU interface {
	NewModuleInstancePerVU() interface{}
}

// Register the given mod as a JavaScript module, available
// for import from JS scripts by name.
// This function panics if a module with the same name is already registered.
func Register(name string, mod interface{}) {
	mx.Lock()
	defer mx.Unlock()

	if _, ok := modules[name]; ok {
		panic(fmt.Sprintf("module already registered: %s", name))
	}
	modules[name] = mod
}

// GetJSModules returns a map of all js modules
func GetJSModules() map[string]HasModuleInstancePerVU {
	result := map[string]HasModuleInstancePerVU{
		// TODO add others
		"k6":             HasModuleInstancePerVUDummyWrapper{Module: k6.New()},
		"k6/crypto":      HasModuleInstancePerVUDummyWrapper{Module: crypto.New()},
		"k6/crypto/x509": HasModuleInstancePerVUDummyWrapper{Module: x509.New()},
		"k6/data":        HasModuleInstancePerVUDummyWrapper{Module: data.New()},
		"k6/encoding":    HasModuleInstancePerVUDummyWrapper{Module: encoding.New()},
		"k6/net/grpc":    HasModuleInstancePerVUDummyWrapper{Module: grpc.New()},
		"k6/html":        HasModuleInstancePerVUDummyWrapper{Module: html.New()},
		"k6/http":        http.New(),
		"k6/metrics":     HasModuleInstancePerVUDummyWrapper{Module: metrics.New()},
		"k6/ws":          HasModuleInstancePerVUDummyWrapper{Module: ws.New()},
	}

	for name, module := range modules {
		if perInstance, ok := module.(HasModuleInstancePerVU); ok {
			result[name] = perInstance
		} else {
			result[name] = HasModuleInstancePerVUDummyWrapper{Module: module}
		}
	}
	return result
}

// HasModuleInstancePerVUDummyWrapper is a wrapper to be used around an module that doesn't support
// HasModuleIntancePerVU, but needs to be used in places where it is required.
type HasModuleInstancePerVUDummyWrapper struct {
	Module interface{}
}

// NewModuleInstancePerVU implements HasModuleInstancePerVU
func (h HasModuleInstancePerVUDummyWrapper) NewModuleInstancePerVU() interface{} {
	return h.Module
}

var _ HasModuleInstancePerVU = HasModuleInstancePerVUDummyWrapper{}
