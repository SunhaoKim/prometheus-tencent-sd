package main

import (
	"fmt"
	"io"
	"net"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"gopkg.in/yaml.v2"
)

const (
	cvmLabelInstanceID   = "instance_id"
	cvmLabelPrivateIP    = "private_ip"
	cvmLabelInstanceName = "instance_name"
	cvmLabelInstanceType = "instance_type"
)

func (s *InstanceSet) To() (tgs []*targetgroup.Group) {
	for _, instance := range s.instances {
		var tg targetgroup.Group

		if instance.PrivateIpAddress == "" {
			continue
		}

		tg.Labels = make(model.LabelSet)
		tg.Labels[cvmLabelInstanceID] = model.LabelValue(instance.InstanceId)
		tg.Labels[cvmLabelInstanceName] = model.LabelValue(instance.InstanceName)
		tg.Labels[cvmLabelInstanceType] = model.LabelValue(instance.InstanceType)

		addr := net.JoinHostPort(instance.PrivateIpAddress, fmt.Sprintf("%d", s.port))
		tg.Targets = append(tg.Targets, model.LabelSet{
			model.AddressLabel: model.LabelValue(addr),
		})

		tgs = append(tgs, &tg)
	}

	return tgs
}

func (s *InstanceSet) Write(w io.Writer) error {
	tgs := s.To()
	return yaml.NewEncoder(w).Encode(tgs)
}
