// Package plugin holds the implementation of a plugin.
package plugin

import "fmt"

func (p *plugin) validateConfig() (errs []error) {
	if p.config.SyncImage == "" {
		errs = append(errs, fmt.Errorf("syncImage cannot be empty"))
	}
	if len(p.config.GenevaConfig.ImagePullSecret) == 0 {
		errs = append(errs, fmt.Errorf("imagePullSecret cannot be empty"))
	}
	if p.config.GenevaConfig.LoggingCert == nil {
		errs = append(errs, fmt.Errorf("loggingCert cannot be nil"))
	}
	if p.config.GenevaConfig.LoggingKey == nil {
		errs = append(errs, fmt.Errorf("loggingKey cannot be nil"))
	}
	return errs
}
