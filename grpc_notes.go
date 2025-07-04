package main

// Make a proto file.
// Generate the go and grpc files from the protoc command.
// Things to underdtand in the grpc.pb.go file (taking the example of this project):
// Protoc first creates a CoffeeShopNimarServer interface that has all the methods defined in the proto file.
// Then Protoc also creates a UnimplementedCoffeeShopNimarServer struct that implements the CoffeeShopNimarServer interface.
// This struct has all the methods defined in the CoffeeShopNimarServer interface, but they are all empty.
// Then we ceate a MyServer struct that embeds the UnimplementedCoffeeShopNimarServer struct.
// This allows us to implement the CoffeeShopNimarServer interface without having to implement all the methods.
// We implement only the methods we need in the MyServer struct.
// Refer to this for more details on this example: https://chatgpt.com/share/68683c74-04f4-8007-ab60-487915d60ed0
// Youtube video of this example: https://www.youtube.com/watch?v=mPESsBfUKkc&t=599s
