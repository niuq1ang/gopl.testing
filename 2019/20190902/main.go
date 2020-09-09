package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const value = `+----------+
| owner    |
+----------+
| QxjnENdy |
+----------+
1 row in set (0.00 sec)
`

func main() {
	req := map[string]interface{}{
		"phone":               "15814757323",
		"new_owner_user_uuid": "BvhVUprQ",
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(string(bytes))

	// fmt.Println(GetSQLSingleValue(value))
	//fmt.Println(UpgradeBackupVolumeName("sabo", "sabo-backup-v1"))
}

func GetSQLSingleValue(value string) string {
	if value == "" {
		return value
	}

	fmt.Println(value)
	arr := strings.Split(value, "\n")
	fmt.Printf("%d,%+v", len(arr), arr)
	if len(arr) <= 4 {
		return value
	}
	fmt.Println(len(arr))
	values := strings.Split(arr[3], "|")
	if len(values) == 3 {
		fmt.Println(strings.Trim(values[1], " "))
		return strings.Trim(values[1], " ")
	}
	return value
}

// func FindAllAliveContainers() ([]types.Container, error) {
// 	dockerCli, err := client.NewEnvClient()
// 	containers, err := dockerCli.ContainerList(context.TODO(), types.ContainerListOptions{})
// 	if err != nil {
// 		return nil, fmt.Errorf("list containers failed: %v", err)
// 	}

// 	return containers, nil
// }

// func ContainerListToString(containers []types.Container) string {
// 	var result string
// 	for index, container := range containers {
// 		containerID := container.ID
// 		containerName := container.Names[0]
// 		imageID := container.ImageID
// 		imageName := container.Image
// 		result = fmt.Sprintf("%d,%s,%s,%s,%s", index, containerID, containerName, imageID, imageName)
// 	}
// 	return result
// }

func UpgradeBackupVolumeName(volumeName, oldBackVolumeName string) string {
	if oldBackVolumeName == "" {
		return fmt.Sprintf("%s-backup-v1", volumeName)
	}

	index := strings.LastIndexAny(oldBackVolumeName, "v")
	if index >= 1 {
		version := oldBackVolumeName[index+1:]
		versionNum, err := strconv.ParseInt(version, 10, 64)
		if err != nil {
			return fmt.Sprintf("%s-backup-v1", volumeName)
		}
		versionNum++
		return fmt.Sprintf("%s-backup-v%d", volumeName, versionNum)
	}
	return fmt.Sprintf("%s-backup-v1", volumeName)
}
