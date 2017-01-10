package main

import (
	"bytes"
	"encoding/json"
	"fmt"
    "github.com/docker/docker/api/client"
    "github.com/docker/engine-api/types" 
    "log"
	"os"
	"runtime"
	"strconv"
	"time"
	//"golang.org/x/net/context"
	"context"
	"github.com/spf13/cobra"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	var interval int
	rootCmd := &cobra.Command{
		Use:   "Goroutine leak debug",
		Short: "debug goroutine leak",
		Run: func(cmd *cobra.Command, args []string) {
			if runtime.GOOS != "linux" {
				fmt.Println("From linux for now, application exit.")
				return
			}
			debugHttp()
			for {
				containerLogsInfo := GetContainerLogsInfo(int64(interval))
				fmt.Println(len(containerLogsInfo))

				if interval == 0 {
					return
				}

				time.Sleep(time.Millisecond * time.Duration(interval))
			}
		},
	}

	rootCmd.Flags().IntVarP(&interval, "interval", "i", 0, "Interval to retrieval data(millisecond), default 0 is not repeat.")

	rootCmd.Execute()
}

type ContainerLogs struct {
	ID    string   `json:"id"`
	Names []string `json:"name"`
	Image string   `json:"image"`
	Log   string   `json:"log"`
}

type ContainerLogsInfo struct {
	Timestamp int64           `json:"timestamp"`
	MachineID string          `json:"machine_id"`
	LogList   []ContainerLogs `json:"log_list"`
}

func GetContainers(cli *client.Client, ctx context.Context) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	return containers, err
}

func GetContainerLogByID(cli *client.Client, ctx context.Context, ID string, opts types.ContainerLogsOptions) string {
	reader, err := cli.ContainerLogs(ctx, ID, opts)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	reader.Close()
	return buf.String()
}

func GetContainerLogsInfo(interval int64) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	
	containers, err := GetContainers(cli, ctx)
	if err != nil {
		log.Fatal("Get containers error:", err)
		return ""
	}

	var containerLogsInfo ContainerLogsInfo
	containerLogsInfo.Timestamp = getUnixTimestamp()
	containerLogsInfo.MachineID = getMachineID()

	opts := types.ContainerLogsOptions{
		Since:      timestampSubInterval(containerLogsInfo.Timestamp, interval),
		ShowStdout: true,
		ShowStderr: true,
	}

    var containerLogs ContainerLogs
	for _, container := range containers {
		containerLogs.ID = container.ID
		containerLogs.Names = container.Names
		containerLogs.Image = container.Image
		containerLogs.Log = GetContainerLogByID(cli, ctx, container.ID, opts)
		containerLogsInfo.LogList = append(containerLogsInfo.LogList, containerLogs)
	}

	fmt.Println("get containers done! num of containers is:", len(containers))

	retJson, err := json.Marshal(containerLogsInfo)
	if err != nil {
		log.Fatal("marshal json error:", err)
	}

	return string(retJson)
}

func getUnixTimestamp() int64 {
	return time.Now().Unix()
}

func getMachineID() string {
	//stdout, err := exec.Command("hostname").CombinedOutput()
	h, err := os.Hostname()
	if err != nil {
		log.Println("hostname cannot retrieval.")
		log.Println(err)
		return ""
	}

	return h
}

func timestampSubInterval(now, interval int64) string {
	//change interval from millsecond to second
	delta := now - interval/1000
	return strconv.FormatInt(delta, 10)
}

func debugHttp() {
	fmt.Println("start pprof")
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:34888", nil))
	}()
	fmt.Println("started pprof")
}
