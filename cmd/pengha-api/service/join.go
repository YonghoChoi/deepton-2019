package service

import "yonghochoi.com/depthon-2019/model/user"

func Join(u user.User) (user.User, error) {
	u.Join()
	if err := user.Insert(u); err != nil {
		return u, err
	}

	return u, nil
}
