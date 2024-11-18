package etcd

import (
	"encoding/json"
	"github.com/svc0a/tools2/tools"
	"testing"
	"time"
)

func Test_gss_Export(t *testing.T) {
	data, err := Define([]string{
		"etcd-headless.gss-test-middleware.svc.cluster.local:2379",
	}, time.Second*30).Export()
	if err != nil {
		return
	}
	marshal, _ := json.Marshal(data)
	tools.RewriteFile("gss.json", marshal)
}

func Test_ghs_Export(t *testing.T) {
	data, err := Define([]string{
		"etcd-headless.ghs-test-middleware.svc.cluster.local:2379",
	}, time.Second*30).Export()
	if err != nil {
		return
	}
	marshal, _ := json.Marshal(data)
	tools.RewriteFile("ghs.json", marshal)
}

func Test_phs_Export(t *testing.T) {
	data, err := Define([]string{
		"etcd-headless.phs-test-middleware.svc.cluster.local:2379",
	}, time.Second*30).Export()
	if err != nil {
		return
	}
	marshal, _ := json.Marshal(data)
	tools.RewriteFile("phs.json", marshal)
}
