package main
import(
	"fmt"
)

func BubbleSort(arr []int){
	for i:=0; i < len(arr)-1; i++{
		for j:=0; j <len(arr) -i-1;j++{
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}

func BinaryFind(arr []int,l int, r int, val int)(int,int){
	if l > r{
		fmt.Println("not found:",val)
		return -1,-1
	}

	m := (l+r)/2
	if arr[m] > val{
		return  BinaryFind(arr,l,m-1,val)
	} else if arr[m] < val{
		return BinaryFind(arr,m+1,r,val)
	}else{
		fmt.Println("found:",val," index:",m)
		return val,m
	}
}

func main(){
	arr := []int{8,9,10,11,5,6,7,4,2,3,1}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
	val,index := BinaryFind(arr,0,len(arr)-1,10)
	fmt.Println("index:",index,"value",val)
}

