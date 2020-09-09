package main

import (
	"fmt"
)

type TestStruct struct {
	Number   int
	Boolen   bool
	String   string
	Byte     byte
	IntArray []int
	Pointer  uintptr
}

const (
	sensorsURLPrefix = `https://shence.ones.team/segmentation/?project=default#measures%5B0%5D%5Bevent_name%5D=TeamActiveData&measures%5B0%5D%5Baggregator%5D=SUM&measures%5B0%5D%5Bfield%5D=event.TeamActiveData.all_page_count&measures%5B1%5D%5Bevent_name%5D=TeamActiveData&measures%5B1%5D%5Baggregator%5D=SUM&measures%5B1%5D%5Bfield%5D=event.TeamActiveData.all_pipeline_count&measures%5B2%5D%5Bevent_name%5D=TeamActiveData&measures%5B2%5D%5Baggregator%5D=SUM&measures%5B2%5D%5Bfield%5D=event.TeamActiveData.all_project_count&measures%5B3%5D%5Bevent_name%5D=TeamActiveData&measures%5B3%5D%5Baggregator%5D=SUM&measures%5B3%5D%5Bfield%5D=event.TeamActiveData.all_testcase_count&measures%5B4%5D%5Bevent_name%5D=TeamActiveData&measures%5B4%5D%5Baggregator%5D=SUM&measures%5B4%5D%5Bfield%5D=event.TeamActiveData.member_count&measures%5B5%5D%5Bevent_name%5D=TeamActiveData&measures%5B5%5D%5Baggregator%5D=SUM&measures%5B5%5D%5Bfield%5D=event.TeamActiveData.member_access_count&unit=day&filter%5Bconditions%5D%5B0%5D%5Bfield%5D=event.TeamActiveData.team_uuid&filter%5Bconditions%5D%5B0%5D%5Bfunction%5D=equal&filter%5Bconditions%5D%5B0%5D%5Bparams%5D%5B%5D=`
	sensorsURLTail   = `&by_fields%5B%5D=event.TeamActiveData.team_name&sampling_factor=64&axis_config%5BisNormalize%5D=false&axis_config%5Bleft%5D%5B%5D=0&axis_config%5Bleft%5D%5B%5D=1&axis_config%5Bleft%5D%5B%5D=2&axis_config%5Bleft%5D%5B%5D=3&axis_config%5Bleft%5D%5B%5D=4&axis_config%5Bleft%5D%5B%5D=5&rangeText=%E8%BF%87%E5%8E%BB+7+%E5%A4%A9&from_date=2019-03-14&to_date=2019-03-20&tType=y&ratio=n&approx=false&chartsType=line&bookmarkid=326`
)

func (self TestStruct) valueTest() {
	fmt.Printf("Value: %p\n", &self)
}

func (self *TestStruct) pointerTest() {
	fmt.Printf("Pointer: %p\n", self)
}

func main() {

	fmt.Println(sensorsURLPrefix + "66gWjFij" + sensorsURLTail)
	d := TestStruct{}
	p := &d

	fmt.Printf("Data: %p\n", p)

	d.valueTest()
	d.pointerTest()

	p.valueTest()
	p.pointerTest()
}
