package main

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type Instance struct {
	InstanceId       string `json:"InstanceId"`
	PrivateIpAddress string `json:"PrivateIpAddress"`
	InstanceName     string `json:"InstanceName"`
	InstanceType     string `json:"InstanceType"`
}

type InstanceSet struct {
	port      int
	source    string
	instances []Instance
}

// 转换数据类型
func from(j *cvm.Instance) Instance {
	var i Instance
	i.InstanceName = *j.InstanceName
	//*string to string **p = p
	for _, p := range j.PrivateIpAddresses {
		i.PrivateIpAddress = *p
	}
	i.InstanceId = *j.InstanceId
	i.InstanceType = *j.InstanceType
	return i
}

//获取instances
func Getinstances(c *cvm.Client) (*InstanceSet, error) {
	request := cvm.NewDescribeInstancesRequest()
	request.Limit = common.Int64Ptr(50)
	for _, filter := range config.Filters {
		request.Filters = append(request.Filters, &cvm.Filter{
			Name:   common.StringPtr(filter.Name),
			Values: common.StringPtrs(filter.Values),
		})
	}
	response, err := c.DescribeInstances(request)
	fmt.Println(response.ToJsonString())
	if err != nil {
		return nil, err
	}
	data := &InstanceSet{
		source: config.Region,
		port:   config.Port,
	}
	//从获取的结果里append 到 instanceset里面，将ins 转换为自己的数据类型
	for _, ins := range response.Response.InstanceSet {
		data.instances = append(data.instances, from(ins))
	}
	return data, nil
}
