package main

import (
	"context"
	"log"
	"userService/pkg/model"
	"userService/pkg/pb"
	"userService/pkg/rbac"

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
		rep, err := client.CheckPermission(ctx, &pb.CheckPermissionRequest{
			Route: "/trnlog/repay/query",
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(rep)
		}
	}

	//
	//{
	//	md := metadata.New(map[string]string{})
	//	md.Set("jwtToken", tk)
	//	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//	rep, err := client.AddRoleForUser(ctx, &pb.AddRoleForUserRequest{
	//		Username: "test2",
	//		Role:     "test2",
	//	})
	//	if err != nil {
	//		log.Println(err)
	//	} else {
	//		log.Println(rep)
	//	}
	//}

	{
		md := metadata.New(map[string]string{})
		md.Set("jwtToken", tk)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		rep, err := client.AddRoutes(ctx, &pb.AddRoutesRequest{
			Routes: []string{"/prodbizfeemap/index"},
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
		rep, err := client.ListRoutes(ctx, &pb.ListRoutesRequest{})
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
		rep, err := client.CreatePermission(ctx, &pb.CreatePermissionRequest{
			Name: "T1补付权限",
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
		rep, err := client.UpdatePermission(ctx, &pb.UpdatePermissionRequest{
			Id:   1,
			Name: "T1补付权限1",
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
		rep, err := client.AddRouteForPermission(ctx, &pb.AddRouteForPermissionRequest{
			Permission: "T1补付权限",
			Route:      "/trnlog/repay/index",
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
		rep, err := client.RemoveRouteForPermission(ctx, &pb.RemoveRouteForPermissionRequest{
			Permission: "T1补付权限2",
			Route:      "/trnlog/repay/repay",
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
		rep, err := client.RemovePermission(ctx, &pb.RemovePermissionRequest{
			Permission: "T1补付权限2",
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
		rep, err := client.ListPermissions(ctx, &pb.ListPermissionsRequest{})
		if err != nil {
			//panic(err)
			log.Println(err)
		} else {
			log.Println(rep.Permissions)
		}
	}

	{
		md := metadata.New(map[string]string{})
		md.Set("jwtToken", tk)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		rep, err := client.AddPermissionForPermission(ctx, &pb.AddPermissionForPermissionRequest{
			From: "T1补付权限",
			To:   "T1补付权限1",
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
		rep, err := client.RemovePermissionForPermission(ctx, &pb.RemovePermissionForPermissionRequest{
			From:  "T1补付权限1",
			Child: "T1补付权限",
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
		rep, err := client.CreateRole(ctx, &pb.CreateRoleRequest{
			Role: "test",
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
		rep, err := client.ListRole(ctx, &pb.ListRoleRequest{})
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
		rep, err := client.UpdateRole(ctx, &pb.UpdateRoleRequest{
			Id:   1,
			Name: "test",
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
		rep, err := client.AddPermissionForRole(ctx, &pb.AddPermissionForRoleRequest{
			Role:       "test",
			Permission: "T1补付权限",
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
		rep, err := client.RemovePermissionForRole(ctx, &pb.RemovePermissionForRoleRequest{
			Role:       "test1",
			Permission: "T1补付权限",
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
		rep, err := client.AddRoleForRole(ctx, &pb.AddRoleForRoleRequest{
			From: "test",
			To:   "test1",
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
		rep, err := client.RemoveRoleForRole(ctx, &pb.RemoveRoleForRoleRequest{
			From:  "test1",
			Child: "test",
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
		rep, err := client.RemoveRole(ctx, &pb.RemoveRoleRequest{
			Role: "test",
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(rep)
		}
	}

	enforce := rbac.NewCasbin("configs/rbac.conf", &model.Options{
		User:     "root",
		Password: "root",
		Addr:     "127.0.0.1:3306",
	})
	log.Println(enforce.Enforce("test", "/trnlog/repay/index"))
}
