package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"sort"
	"strings"
)

func ListRegions(code string) []string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)
	input := &ec2.DescribeRegionsInput{}

	regions, err := svc.DescribeRegions(input)
	if err != nil {
		log.Fatalf("ERROR: describing regions. \n%s", err.Error())
	}

	var regionList []string

	for _, region := range regions.Regions {
		if strings.HasPrefix(*region.RegionName, code) {
			regionList = append(regionList, *region.RegionName)
		}
	}

	sort.Strings(regionList)

	return regionList
}
