package integration

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"testing"
	"time"
	"todolist/app"
	"todolist/config"
	"todolist/domain"
	"todolist/proto"
	"todolist/server"
)

const address = "localhost:8080"

var idOne, idTwo string

func getServer() (*grpc.Server, net.Listener) {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	apiServer := server.NewServer(app.SetUpDependencies())
	grpcServer := grpc.NewServer()
	proto.RegisterApiServer(grpcServer, apiServer)

	listener, err := net.Listen(config.GetServerConfig().Protocol(), address)
	if err != nil {
		panic(err)
	}

	return grpcServer, listener
}

func startServer(server *grpc.Server, listener net.Listener) {
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

func stopServer(server *grpc.Server) {
	server.GracefulStop()
}

func waitForServer() {
	time.Sleep(time.Second)
}

func TestAllAPIs(t *testing.T) {
	grpcServer, listener := getServer()
	go startServer(grpcServer, listener)
	waitForServer()
	defer stopServer(grpcServer)

	c, err := grpc.Dial(address, grpc.WithInsecure())

	defer func() {
		if err := c.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	client := proto.NewApiClient(c)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	testPing(ctx, client, t)
	testAdd(ctx, client, t)
	testUpdate(ctx, client, t)
	testGet(ctx, client, t)
	testRemove(ctx, client, t)
}

func testPing(ctx context.Context, client proto.ApiClient, t *testing.T) {
	t.Run("test ping", func(t *testing.T) {
		resp, err := client.Ping(ctx, &proto.PingRequest{})
		require.NoError(t, err)

		assert.Equal(t, "pong", string(resp.Data.Value))
	})
}

func testAdd(ctx context.Context, client proto.ApiClient, t *testing.T) {
	type testCases struct {
		name    string
		request *proto.AddRequest
	}

	addTests := make([]testCases, 0)

	addTests = append(addTests,
		testCases{
			name: "test add task buy groceries",
			request: &proto.AddRequest{
				Title:       "buy groceries",
				Description: "bread, biscuit",
				Status:      false,
				Tags:        []string{"weekly"},
			},
		},
		testCases{
			name: "test add task read",
			request: &proto.AddRequest{
				Title:       "read",
				Description: "book, xyz",
				Status:      false,
				Tags:        []string{"daily", "personal"},
			},
		},
	)

	for _, addTest := range addTests {
		t.Run(addTest.name, func(t *testing.T) {
			resp, err := client.Add(ctx, addTest.request)
			require.NoError(t, err)
			require.NotNil(t, resp)
		})
	}
}

func testUpdate(ctx context.Context, client proto.ApiClient, t *testing.T) {
	type testCases struct {
		name    string
		request func() *proto.UpdateRequest
	}

	updateTests := make([]testCases, 0)

	updateTests = append(updateTests,
		testCases{
			name: "change description and tag for buy groceries",
			request: func() *proto.UpdateRequest {
				res, err := client.Get(ctx, &proto.GetRequest{})
				require.NoError(t, err)

				var tasks []domain.DefaultTask
				err = json.Unmarshal(res.Data.Value, &tasks)
				require.NoError(t, err)

				var task domain.DefaultTask
				for _, t := range tasks {
					if t.GetTitle() == "buy groceries" {
						idOne = t.GetID()
						task = t
						break
					}
				}

				task.UpdateDescription("bread, biscuit, cake")
				task.UpdateTags("weekly", "grocery")
				task.UpdateStatus()

				return &proto.UpdateRequest{
					Id:          task.GetID(),
					Title:       task.GetTitle(),
					Description: task.GetDescription(),
					Status:      task.GetStatus(),
					Tags:        task.GetTags(),
				}
			},
		},
		testCases{
			name: "change tag for read",
			request: func() *proto.UpdateRequest {
				res, err := client.Get(ctx, &proto.GetRequest{})
				require.NoError(t, err)

				var tasks []domain.DefaultTask
				err = json.Unmarshal(res.Data.Value, &tasks)
				require.NoError(t, err)

				var task domain.DefaultTask
				for _, t := range tasks {
					if t.GetTitle() == "read" {
						idTwo = t.GetID()
						task = t
						break
					}
				}

				task.UpdateTags("daily")
				task.UpdateStatus()

				return &proto.UpdateRequest{
					Id:          task.GetID(),
					Title:       task.GetTitle(),
					Description: task.GetDescription(),
					Status:      task.GetStatus(),
					Tags:        task.GetTags(),
				}
			},
		},
	)

	for _, updateTest := range updateTests {
		t.Run(updateTest.name, func(t *testing.T) {
			resp, err := client.Update(ctx, updateTest.request())
			require.NoError(t, err)
			require.NotNil(t, resp)
		})
	}
}

func testGet(ctx context.Context, client proto.ApiClient, t *testing.T) {
	type testCases struct {
		name     string
		request  *proto.GetRequest
		response func() *proto.ApiResponse
	}

	getTests := make([]testCases, 0)

	getTests = append(getTests,
		testCases{
			name:    "test get all tasks",
			request: &proto.GetRequest{},
			response: func() *proto.ApiResponse {
				tasks := []domain.DefaultTask{
					{
						Title:       "buy groceries",
						Description: "bread, biscuit, cake",
						Status:      true,
						Tags:        pq.StringArray{"weekly", "grocery"},
					},
					{
						Title:       "read",
						Description: "book, xyz",
						Status:      true,
						Tags:        pq.StringArray{"daily"},
					},
				}

				b, err := json.Marshal(tasks)
				require.NoError(t, err)

				return &proto.ApiResponse{
					Data: &any.Any{
						Value: b,
					},
				}
			},
		},
		testCases{
			name:    "test get buy grocery task",
			request: &proto.GetRequest{Id: []string{idOne}},
			response: func() *proto.ApiResponse {
				tasks := []domain.DefaultTask{
					{
						Title:       "buy groceries",
						Description: "bread, biscuit, cake",
						Status:      true,
						Tags:        pq.StringArray{"weekly", "grocery"},
					},
				}

				b, err := json.Marshal(tasks)
				require.NoError(t, err)

				return &proto.ApiResponse{
					Data: &any.Any{
						Value: b,
					},
				}
			},
		},
		testCases{
			name:    "test get read task",
			request: &proto.GetRequest{Id: []string{idTwo}},
			response: func() *proto.ApiResponse {
				tasks := []domain.DefaultTask{
					{
						Title:       "read",
						Description: "book, xyz",
						Status:      true,
						Tags:        pq.StringArray{"daily"},
					},
				}

				b, err := json.Marshal(tasks)
				require.NoError(t, err)

				return &proto.ApiResponse{
					Data: &any.Any{
						Value: b,
					},
				}
			},
		},
	)

	for _, getTest := range getTests {
		t.Run(getTest.name, func(t *testing.T) {
			resp, err := client.Get(ctx, getTest.request)
			require.NoError(t, err)

			var actualTasks, expectedTasks []domain.DefaultTask

			err = json.Unmarshal(getTest.response().Data.Value, &expectedTasks)
			require.NoError(t, err)

			err = json.Unmarshal(resp.Data.Value, &actualTasks)
			require.NoError(t, err)

			assert.Equal(t, len(expectedTasks), len(actualTasks))

			l := len(expectedTasks)

			for i := 0; i < l; i++ {
				assert.Equal(t, expectedTasks[i].GetTitle(), actualTasks[i].GetTitle())
				assert.Equal(t, expectedTasks[i].GetDescription(), actualTasks[i].GetDescription())
				assert.Equal(t, expectedTasks[i].GetStatus(), actualTasks[i].GetStatus())
				assert.Equal(t, expectedTasks[i].GetTags(), actualTasks[i].GetTags())
			}

		})
	}

}

func testRemove(ctx context.Context, client proto.ApiClient, t *testing.T) {
	type testCases struct {
		name    string
		request *proto.RemoveRequest
	}

	removeTests := make([]testCases, 0)

	removeTests = append(removeTests,
		testCases{
			name:    "test remove buy groceries",
			request: &proto.RemoveRequest{Id: idOne},
		},
		testCases{
			name:    "test remove read",
			request: &proto.RemoveRequest{Id: idTwo},
		},
	)

	for _, removeTest := range removeTests {
		t.Run(removeTest.name, func(t *testing.T) {
			resp, err := client.Remove(ctx, removeTest.request)
			require.NoError(t, err)
			require.NotNil(t, resp)
		})
	}

}
