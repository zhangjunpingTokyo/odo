package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/openshift/odo/pkg/util"
)

func TestSetLocalConfiguration(t *testing.T) {

	tempConfigFile, err := ioutil.TempFile("", "odoconfig")
	if err != nil {
		t.Fatal(err)
	}
	defer tempConfigFile.Close()
	os.Setenv(localConfigEnvName, tempConfigFile.Name())
	minCPUValue := "0.5"
	maxCPUValue := "2"
	minMemValue := "500M"
	maxMemValue := "1000M"
	testValue := "test"
	portsValue := "8080/TCP,45/UDP"
	typeValue := "nodejs"
	applicationValue := "odotestapp"
	projectValue := "odotestproject"
	sourceTypeValue := "git"
	sourceLocationValue := "https://github.com/sclorg/nodejs-ex"
	refValue := "develop"

	tests := []struct {
		name           string
		parameter      string
		value          string
		existingConfig LocalConfig
	}{
		// update notification
		{
			name:      fmt.Sprintf("Case 1: %s set nil to true", Ignore),
			parameter: Ignore,
			value:     "true",
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 2: %s set true to false", Ignore),
			parameter: Ignore,
			value:     "false",
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 3: %s to test", Name),
			parameter: Name,
			value:     testValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 5: %s set to %s from 0", MaxCPU, maxCPUValue),
			parameter: MaxCPU,
			value:     maxCPUValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 6: %s set to %s", MinCPU, minCPUValue),
			parameter: MinCPU,
			value:     minCPUValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 6: %s set to %s", MinMemory, minMemValue),
			parameter: MinMemory,
			value:     minMemValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 7: %s set to %s", MaxMemory, maxCPUValue),
			parameter: MaxMemory,
			value:     maxMemValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 8: %s set to %s", Ports, portsValue),
			parameter: Ports,
			value:     portsValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 9: %s set to %s", Type, typeValue),
			parameter: Type,
			value:     typeValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 10: %s set to %s", Application, applicationValue),
			parameter: Application,
			value:     applicationValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 11: %s set to %s", Project, projectValue),
			parameter: Project,
			value:     projectValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 12: %s set to %s", SourceType, sourceTypeValue),
			parameter: SourceType,
			value:     sourceTypeValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 12: %s set to %s", SourceLocation, sourceLocationValue),
			parameter: SourceLocation,
			value:     sourceLocationValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
		{
			name:      fmt.Sprintf("Case 13: %s set to %s", Ref, refValue),
			parameter: Ref,
			value:     refValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := NewLocalConfigInfo("", false)
			if err != nil {
				t.Error(err)
			}
			cfg.LocalConfig = tt.existingConfig

			err = cfg.SetConfiguration(tt.parameter, tt.value)
			if err != nil {
				t.Error(err)
			}

			isSet := cfg.IsSet(tt.parameter)

			if !isSet {
				t.Errorf("the '%v' is not set", tt.parameter)
			}

		})
	}
}

func TestLocalUnsetConfiguration(t *testing.T) {
	tempConfigFile, err := ioutil.TempFile("", "odoconfig")
	if err != nil {
		t.Fatal(err)
	}
	defer tempConfigFile.Close()
	os.Setenv(localConfigEnvName, tempConfigFile.Name())
	trueValue := true
	minCPUValue := "0.5"
	maxCPUValue := "2"
	minMemValue := "500M"
	testValue := "test"

	tests := []struct {
		name           string
		parameter      string
		value          string
		existingConfig LocalConfig
	}{
		// update notification
		{
			name:      fmt.Sprintf("Case 1: unset %s", Ignore),
			parameter: Ignore,
			value:     "true",
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{
					Ignore: &trueValue,
				},
			},
		},
		{
			name:      fmt.Sprintf("Case 3: unset %s", Name),
			parameter: Name,
			value:     testValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{
					Name: &testValue,
				},
			},
		},
		{
			name:      fmt.Sprintf("Case 5: unset %s", MaxCPU),
			parameter: MaxCPU,
			value:     maxCPUValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{
					MaxCPU: &maxCPUValue,
				},
			},
		},
		{
			name:      fmt.Sprintf("Case 6: unset %s", MinCPU),
			parameter: MinCPU,
			value:     minCPUValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{
					MinCPU: &minCPUValue,
				},
			},
		},
		{
			name:      fmt.Sprintf("Case 6: unset %s", MinMemory),
			parameter: MinMemory,
			value:     minMemValue,
			existingConfig: LocalConfig{
				componentSettings: ComponentSettings{
					MinMemory: &minMemValue,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := NewLocalConfigInfo("", false)
			if err != nil {
				t.Error(err)
			}
			cfg.LocalConfig = tt.existingConfig

			err = cfg.SetConfiguration(tt.parameter, tt.value)
			if err != nil {
				t.Error(err)
			}
			isSet := cfg.IsSet(tt.parameter)
			if !isSet {
				t.Errorf("the '%v' was not set", tt.parameter)
			}

			err = cfg.DeleteConfiguration(tt.parameter)

			if err != nil {
				t.Error(err)
			}
			isSet = cfg.IsSet(tt.parameter)
			if isSet {
				t.Errorf("the '%v' is not set to nil", tt.parameter)
			}

		})
	}
}

func TestLowerCaseParameterForLocalParameters(t *testing.T) {
	expected := map[string]bool{"name": true, "minmemory": true, "ignore": true, "project": true,
		"application": true, "type": true, "ref": true, "mincpu": true, "cpu": true, "ports": true, "maxmemory": true,
		"maxcpu": true, "sourcetype": true, "sourcelocation": true, "memory": true}
	actual := util.GetLowerCaseParameters(GetLocallySupportedParameters())
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected '%v', got '%v'", expected, actual)
	}
}

func TestLocalConfigInitDoesntCreateLocalOdoFolder(t *testing.T) {
	// cleaning up old odo files if any
	filename, err := getLocalConfigFile("")
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll(filename)

	conf, err := NewLocalConfigInfo("", false)
	if err != nil {
		t.Errorf("error while creating local config %v", err)
	}
	if _, err = os.Stat(conf.Filename); !os.IsNotExist(err) {
		t.Errorf("local config.yaml shouldn't exist yet")
	}
}

func TestMetaTypePopulatedInLocalConfig(t *testing.T) {
	ci, err := NewLocalConfigInfo("", false)

	if err != nil {
		t.Error(err)
	}
	if ci.typeMeta.APIVersion != localConfigAPIVersion || ci.typeMeta.Kind != localConfigKind {
		t.Error("the api version and kind in local config are incorrect")
	}
}
