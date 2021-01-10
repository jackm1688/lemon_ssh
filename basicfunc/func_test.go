package basicfunc

//功能用例测试
import "testing"

func TestBasic(t *testing.T)  {
	grade := GetGrade(40)
	if grade != "D"{
		t.Error("Test case failed")
	}
}

func TestAddFunc(t *testing.T)  {
	sum := Add(1,2)
	if sum == 3{
		t.Log("Passed: 1+1==2")
	}else{
		t.Log("Failed: 1+1==2")
	}
}
