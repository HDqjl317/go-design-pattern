/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-04-18 13:57:23
 * @LastEditTime: 2023-04-18 14:05:42
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package composite

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e Employee) Count() int {
	return 1
}

type Department struct {
	Name             string
	SubOrgnaizations []IOrganization
}

func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrgnaizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSubs(org IOrganization) {
	d.SubOrgnaizations = append(d.SubOrgnaizations, org)
}

func NewOrgnization() IOrganization {
	root := &Department{
		Name: "root",
	}
	for i := 0; i < 10; i++ {
		root.AddSubs(&Employee{})
		root.AddSubs(&Department{
			Name:             "sub",
			SubOrgnaizations: []IOrganization{&Employee{}},
		})
	}
	return root
}
