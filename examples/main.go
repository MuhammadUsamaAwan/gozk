package main

import (
	"fmt"

	"github.com/MuhammadUsamaAwan/gozk"
)

func main() {
	zk := gozk.NewZK("192.168.100.201", 4370, 0, gozk.DefaultTimezone)
	if err := zk.Connect(); err != nil {
		panic(err)
	}

	defer zk.Disconnect()

	properties, err := zk.GetProperties()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total Users: %d\n", properties.TotalUsers)
	fmt.Printf("Total Fingers: %d\n", properties.TotalFingers)
	fmt.Printf("Total Records: %d\n", properties.TotalRecords)
	fmt.Printf("Finger Capacity: %d\n", properties.FingerCap)
	fmt.Printf("User Capacity: %d\n", properties.UserCap)
	fmt.Printf("Record Capacity: %d\n", properties.RecordCap)

	attendances, err := zk.GetAttendances()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Attendances: %d\n", len(attendances))
}
