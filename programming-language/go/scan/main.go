package main

import (
	"fmt"

	// `context`는 Docker 클라이언트와 상호작용할 때 컨텍스트를 제공하는 데 사용된다.
	"context"
	// Docker 관련 패키지
	"github.com/moby/moby/client"
)

func main() {
	var userInput string

	// `Scanln`은 표준 입력으로부터 한 줄을 읽어와서 지정된 변수에 저장하지만,
	// 입력이 공백으로 구분된 여러 값일 경우 첫 번째 값만 읽어온다.
	// 단 여러 값을 입력 받고자 한다면 (&n, &n, ...)와 같이 여러 변수를 지정해야 한다.
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch userInput {
	case "list":
		// Docker Client 생성
		dockerClient, err := client.New(
			// Docker 클라이언트 설정을 환경 변수에서 가져옴
			client.FromEnv,
			// Ref `WithUserAgent` 옵션을 사용하여 사용자 에이전트를 설정함
			client.WithUserAgent("my-app/v1.0.1"),
		)

		if err != nil {
			fmt.Println("Error creating Docker client:", err)
			return
		}

		defer dockerClient.Close()

		// Docker Container의 운용 상태를 조회함 동시에 모든 컨테이너 정보를 가져옴
		result, err := dockerClient.ContainerList(context.Background(), client.ContainerListOptions{
			All: true,
		})

		if err != nil {
			fmt.Println("Error listing Docker containers:", err)
			return
		}

		// fmt.Println("Docker containers list result:", result)

		fmt.Printf("%s  %-22s  %s\n", "ID", "STATUS", "IMAGE")

		for _, ctr := range result.Items {
			fmt.Printf("%s  %-22s  %s\n", ctr.ID, ctr.Status, ctr.Image)
		}
		return
	case "images":
		dockerClient, err := client.New(
			client.FromEnv,
			client.WithUserAgent("my-app/v1.0.1"),
		)

		if err != nil {
			fmt.Println("Error creating Docker client:", err)
			return
		}

		defer dockerClient.Close()

		// Docker 이미지 정보를 조회함 동시에 모든 이미지 정보를 가져옴
		result, err := dockerClient.ImageList(context.Background(), client.ImageListOptions{
			All: true,
		})

		if err != nil {
			fmt.Println("Error listing Docker images:", err)
			return
		}

		fmt.Printf("%s  %s\n", "ID", "REPOSITORY")
		for _, img := range result.Items {
			fmt.Printf("%s  %s\n", img.ID, img.RepoTags)
		}
		return
	}
}
