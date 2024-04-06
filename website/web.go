package website

import "github.com/Tontonnow/ddddhmlist/website/Friday"

func Init() error {
	err := Friday.RefreshToken()
	if err != nil {
		return err
	}
	return nil
}
