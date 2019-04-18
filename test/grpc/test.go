package main

import (
	"context"
	"log"
	"userService/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewUserClient(conn)
	tk := ""
	{
		rep, err := client.Login(context.Background(), &pb.LoginRequest{
			Username: "test2",
			Password: "111111",
		})
		if err != nil {
			//panic(err)
			log.Println(err)
		} else {
			log.Println(rep)
			tk = rep.Token
		}
	}

	{
		rep, err := client.Register(context.Background(), &pb.RegisterRequest{
			Username:  "test2",
			Password:  "111111",
			UserType:  "ABC",
			Email:     "liqingsanjin@163.com",
			LeaguerNo: "00007294",
		})
		if err != nil {
			//panic(err)
			log.Println(err)
		} else {
			log.Println(rep)
		}
	}

	{
		md := metadata.New(map[string]string{})
		md.Set("jwtToken", tk)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		rep, err := client.GetPermissions(ctx, &pb.GetPermissionsRequest{})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(rep)
		}
	}

	{
		md := metadata.New(map[string]string{})
		md.Set("jwtToken", tk)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		rep, err := client.AddPermission(ctx, &pb.AddPermissionRequest{
			Role:       "test2",
			Permission: "/trnlog/repay/query",
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(rep)
		}
	}

	{
		md := metadata.New(map[string]string{})
		md.Set("jwtToken", tk)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		rep, err := client.CheckPermission(ctx, &pb.CheckPermissionRequest{
			Route: "/trnlog/repay/query",
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(rep)
		}
	}

	{
		md := metadata.New(map[string]string{})
		md.Set("jwtToken", tk)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		rep, err := client.AddRole(ctx, &pb.AddRoleRequest{
			Role: "test2",
			On:   "admin",
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(rep)
		}
	}
}
