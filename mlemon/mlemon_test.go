package mlemon

import (
	"mlemon_ssh/model"
	"testing"
)

func TestMLemon_Exec(t *testing.T) {
	user := model.User{
		User:       "root",
		Password:   "abc-123",
		Host:       "192.168.0.200",
		Key:        "",
		Port:       22,
	}

	m := NewMLemon(&user)
	/*result,err := m.Exec("ls -l")
	if err != nil {
		t.Error(err,)
	}

	fmt.Println(result)*/
	err := m.Scp("host.json.list","/root",2048)
	if err != nil {
		t.Fatal(err)
	}
}