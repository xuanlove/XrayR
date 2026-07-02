package controller

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/xtls/xray-core/infra/conf"

	"github.com/XrayR-project/XrayR/api"
)

// addTUNInbound registers a TUN inbound handler when TUNConfig is enabled.
// Introduced in Xray-core v26.1.23. The TUN inbound allows transparent
// proxying on the node host itself.
func (c *Controller) addTUNInbound() error {
	tc := c.config.TUNConfig
	if tc == nil || !tc.Enable {
		return nil
	}

	tag := tc.Tag
	if tag == "" {
		tag = fmt.Sprintf("tun_%s", c.Tag)
	}

	// TUN inbound settings are forwarded verbatim from config; if empty,
	// use a minimal default that lets xray-core pick an interface address.
	settings := tc.SettingsRaw
	if len(settings) == 0 {
		settings = json.RawMessage(`{}`)
	}

	inboundDetourConfig := &conf.InboundDetourConfig{
		Protocol: "tun",
		Tag:      tag,
		Settings: &settings,
		PortList: &conf.PortList{}, // TUN does not use a real port, but the builder requires a non-nil value
	}

	inboundConfig, err := inboundDetourConfig.Build()
	if err != nil {
		return fmt.Errorf("build TUN inbound failed: %w", err)
	}
	return c.addInbound(inboundConfig)
}

// buildProcessRules precompiles the process-name/path reject list into
// api.DetectRule entries. Xray-core v26.1.23+ supports a "process" routing
// rule; XrayR reuses the rule manager to reject matched processes at the
// dispatcher layer for audit consistency.
func (c *Controller) buildProcessRules() []api.DetectRule {
	pc := c.config.ProcessRouteConfig
	if pc == nil || !pc.Enable {
		return nil
	}
	rules := make([]api.DetectRule, 0, len(pc.RejectNames)+len(pc.RejectPaths))
	for i, name := range pc.RejectNames {
		rules = append(rules, api.DetectRule{
			ID:      -(i + 1),
			Pattern: regexp.MustCompile(regexp.QuoteMeta(name)),
		})
	}
	for i, p := range pc.RejectPaths {
		rules = append(rules, api.DetectRule{
			ID:      -(len(pc.RejectNames) + i + 1),
			Pattern: regexp.MustCompile(regexp.QuoteMeta(p)),
		})
	}
	return rules
}

// startMetricsServer logs the intended metrics endpoint. xray-core's app/metrics
// is registered via cmd/distro/all and configured through the core Config; the
// actual HTTP listener is owned by xray-core's metrics app. XrayR exposes the
// Listen/Path through config so operators know where to scrape, and the metrics
// app picks up its tag from the core instance's policy/stats managers.
func (c *Controller) startMetricsServer() error {
	mc := c.config.MetricsConfig
	if mc == nil || !mc.Enable {
		return nil
	}
	if mc.Listen == "" {
		return fmt.Errorf("MetricsConfig enabled but Listen is empty")
	}
	// The metrics app is started by xray-core itself; we only validate config
	// here and rely on the core instance to serve /metrics on mc.Listen.
	path := mc.Path
	if path == "" {
		path = "/metrics"
	}
	c.logger.Printf("Metrics endpoint configured: http://%s%s", mc.Listen, path)
	return nil
}

// applyEnhancements wires up all v26.x enhancements after the main inbound
// is registered. Errors are logged but non-fatal to preserve node availability.
func (c *Controller) applyEnhancements() {
	if c.config.TUNConfig != nil && c.config.TUNConfig.Enable {
		if err := c.addTUNInbound(); err != nil {
			c.logger.Printf("TUN inbound failed: %s", err)
		} else {
			c.logger.Printf("TUN inbound registered")
		}
	}

	if c.config.MetricsConfig != nil && c.config.MetricsConfig.Enable {
		if err := c.startMetricsServer(); err != nil {
			c.logger.Printf("Metrics server failed: %s", err)
		}
	}

	// Process rules augment the existing audit rule set.
	if c.config.ProcessRouteConfig != nil && c.config.ProcessRouteConfig.Enable {
		if rules := c.buildProcessRules(); len(rules) > 0 {
			if err := c.UpdateRule(c.Tag, rules); err != nil {
				c.logger.Printf("Process route rules failed: %s", err)
			} else {
				c.logger.Printf("Process route rules registered (%d)", len(rules))
			}
		}
	}
}
