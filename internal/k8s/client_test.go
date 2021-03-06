package k8s

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

//go:embed _kubeconfig_test.yaml
var kubeconfig string

func TestBuildK8sConfig(t *testing.T) {
	c := prepareKubeconfig()
	defer func() {
		_ = os.Remove(c)
	}()

	config, ns, err := buildK8sConfig(c, "")
	assert.NotNil(t, config)
	assert.Equal(t, "namespace1", ns)
	assert.Nil(t, err)
	config, ns, err = buildK8sConfig(c, "context-2")
	assert.Nil(t, err)
	assert.Equal(t, "namespace2", ns)
	assert.NotNil(t, config)
	config, ns, err = buildK8sConfig(c, "context-nonexistent")
	assert.Nil(t, config)
	assert.Equal(t, "", ns)
	assert.NotNil(t, err)
}

func prepareKubeconfig() string {
	testConfig, _ := ioutil.TempFile("", "pv-migrate-testconfig-*.yaml")
	_, _ = testConfig.WriteString(kubeconfig)
	return testConfig.Name()
}
