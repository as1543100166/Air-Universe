package V2RayAPI

import (
	"github.com/crossfw/Air-Universe/pkg/structures"
)

func (V2rayCtl *V2rayController) AddUsers(users *[]structures.UserInfo) error {

	for _, u := range *users {
		err := v2AddUser(*V2rayCtl.HsClient, &u)
		if err != nil {
			return err
		}
	}

	return nil
}

func (V2rayCtl *V2rayController) RemoveUsers(users *[]structures.UserInfo) error {
	for _, u := range *users {
		err := v2RemoveUser(*V2rayCtl.HsClient, &u)
		if err != nil {
			return err
		}
	}

	return nil
}

func (V2rayCtl *V2rayController) QueryUsersTraffic(users *[]structures.UserInfo) (usersTraffic *[]structures.UserTraffic, err error) {
	usersTraffic = new([]structures.UserTraffic)
	var ut structures.UserTraffic

	for _, u := range *users {
		ut.Id = u.Id
		ut.Up, err = v2QueryUserTraffic(*V2rayCtl.SsClient, u.Tag, "up")
		ut.Down, err = v2QueryUserTraffic(*V2rayCtl.SsClient, u.Tag, "down")
		// when a user used this node, post traffic data
		if ut.Up+ut.Down > 0 {
			*usersTraffic = append(*usersTraffic, ut)
		}
		if err != nil {
			return
		}
	}
	return
}
