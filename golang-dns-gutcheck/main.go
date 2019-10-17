package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"reflect"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type globalConfig struct {
	LogLevel string  `yaml:"log_level"`
	Interval string  `yaml:"interval"`
	Checks   []check `yaml:"checks"`
}

type check struct {
	Alias               string   `yaml:"alias"`
	Host                string   `yaml:"host"`
	ExpectedARecords    []string `yaml:"expected_a_records"`
	ExpectedShouldMatch string   `yaml:"expected_should_match"`
	HaltOnFailure       bool     `yaml:"halt_on_failure,omitempty"`
	LogLevel            string   `yaml:"log_level"`
}

func (c *check) evaluate(log *log.Entry) error {

	ips, err := net.LookupIP(c.Host)
	if err != nil {
		return err
	}

	var sortedIps []string = nil
	for _, ip := range ips {
		log.Debugf("%s IN A %s\n", c.Host, ip.String())
		sortedIps = append(sortedIps, ip.String())
	}

	sort.Strings(c.ExpectedARecords)
	log.Tracef("Sorted expected_a_records %s", c.ExpectedARecords)

	sort.Strings(sortedIps)
	log.Tracef("Sorted foundIps %s", sortedIps)

	switch c.ExpectedShouldMatch {
	case "exact":
		if reflect.DeepEqual(c.ExpectedARecords, sortedIps) {
			log.Tracef("Exact match between expected %s and found IPs %s", c.ExpectedARecords, sortedIps)
			return nil
		}
		err := fmt.Errorf("Expected records '%s' do not exact-match found IPs %s", c.ExpectedARecords, sortedIps)
		return err
	case "subsetOfFound":
		for _, expected := range c.ExpectedARecords {
			i := sort.Search(len(sortedIps), func(i int) bool { return sortedIps[i] >= expected })
			if i < len(sortedIps) && sortedIps[i] == expected {
				log.Tracef("Expected '%s' is present at index '%d' within %s", expected, i, sortedIps)
			} else {
				err := fmt.Errorf("Could not find expected record '%s' within %s", expected, sortedIps)
				return err
			}
		}
		return nil
	case "any":
		for _, expected := range c.ExpectedARecords {
			i := sort.Search(len(sortedIps), func(i int) bool { return sortedIps[i] >= expected })
			if i < len(sortedIps) && sortedIps[i] == expected {
				log.Tracef("Expected '%s' is present at index '%d' within %s", expected, i, sortedIps)
				return nil
			}
			log.Tracef("Could not find expected record '%s' within %s", expected, sortedIps)
		}
		err := fmt.Errorf("Could not find a single expected record from %s within %s", c.ExpectedARecords, sortedIps)
		return err
	default:
		err := fmt.Errorf("Unknown value for 'expected_should_match': '%s'", c.ExpectedShouldMatch)
		return err
	}
}

func main() {
	configFile, err := ioutil.ReadFile("config.example.yml")
	if err != nil {
		log.Fatalln(err)
	}

	var globalConfig globalConfig

	err = yaml.UnmarshalStrict(configFile, &globalConfig)
	if err != nil {
		log.Fatalln(err)
	}

	globalLogLevel, err := log.ParseLevel(globalConfig.LogLevel)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(globalLogLevel)

	interval, err := time.ParseDuration(globalConfig.Interval)
	if err != nil {
		log.Fatalln(err)
	}

	log.Trace("Beginning main application loop")
	for true {
		for _, check := range globalConfig.Checks {
			log.Infof("Evaluating %s", check.Alias)

			if check.LogLevel != "" {
				localLogLevel, err := log.ParseLevel(check.LogLevel)
				if err != nil {
					log.Fatalln(err)
				}
				log.SetLevel(localLogLevel)
				log.Tracef("Setting log level to %s", localLogLevel)
			} else {
				log.SetLevel(globalLogLevel)
				log.Tracef("Setting log level to %s", globalLogLevel)
			}

			ctxLog := log.WithFields(log.Fields{
				"alias": check.Alias,
			})

			err := check.evaluate(ctxLog)
			if err != nil {
				if check.HaltOnFailure {
					log.Fatalln(err)
				}
				log.Error(err)
			} else {
				log.Debugf("Success on %s", check.Alias)
			}
		}

		log.Debugf("Sleeping for %s", interval)
		time.Sleep(interval)
	}
}
