package domain

import "github.com/GeneralFire/TillSummerBotGo/libs/timecalculator"

type Domain struct{}

func New() Domain {
	return Domain{}
}

func (d *Domain) GetTimeTillSummer() string {
	return timecalculator.GetTimeTillSummer().String()
}

func (d *Domain) GetHello() string {
	return "Hello"
}

func (d *Domain) GetPassedTime() string {
	return "3 days"
}
