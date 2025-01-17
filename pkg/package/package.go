package _package

import (
	"path/filepath"

	"github.com/rancherlabs/corral/pkg/version"
)

const (
	corralUserAgent = "Corral/" + version.Version

	TerraformVersionAnnotation = "corral.cattle.io/terraform-version"
	PublisherAnnotation        = "corral.cattle.io/published-by"
	CorralVersionAnnotation    = "corral.cattle.io/corral-version"
	PublishTimestampAnnotation = "corral.cattle.io/published-at"
)

type Package struct {
	RootPath string

	Manifest
}

func (b Package) TerraformVersion() string {
	v := b.Manifest.GetAnnotation(TerraformVersionAnnotation)

	if v == "" {
		v = version.TerraformVersion
	}

	return v
}

func (b Package) ManifestPath() string {
	return filepath.Join(b.RootPath, "manifest.yaml")
}

func (b Package) TerraformModulePath() string {
	return filepath.Join(b.RootPath, "terraform", "module")
}

func (b *Package) ScriptPath() string {
	return filepath.Join(b.RootPath, "scripts")
}
