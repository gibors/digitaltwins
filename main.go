package main

import (
	"caidc_auto_devicetwins/config"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/service"
	"caidc_auto_devicetwins/domain/utils"
	pb "caidc_auto_devicetwins/simulator"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const REQUIRED_PARAMS = "Required parameters Device Createtion simulator: caidc_auto_device_twins {tenantId} {model} {numberOfDevices} /n" +
	"Device events simulation: caidc_auto_device_twins {tenantId} {serialNumber} {queue_event_name} {numberOfEvents}"
const (
	port = ":50051"
)

type deviceTwinServer struct {
	pb.UnimplementedDeviceTwinServiceServer
}

func (s *deviceTwinServer) CreateDeviceSimulated(ctx context.Context, in *pb.DevicesCreationRequest) (*pb.DeviceSimulationResponse, error) {

	succed := &wrappers.BoolValue{Value: true}
	message := &wrappers.StringValue{Value: "Successfully created device(s)"}
	log.Printf("Received tenantId: %v", in.GetTenantId())
	return &pb.DeviceSimulationResponse{SimulationSucced: succed, DetailedMessage: message}, nil
}

func (s *deviceTwinServer) SendDataToDevice(ctx context.Context, in *pb.DevicesDataSimulationRequest) (*pb.DeviceSimulationResponse, error) {
	succed := &wrappers.BoolValue{Value: true}
	message := &wrappers.StringValue{Value: "Successfully created device(s)"}
	log.Printf("Received tenantId: %v", in.GetTenantId())
	return &pb.DeviceSimulationResponse{SimulationSucced: succed, DetailedMessage: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	deviceSimulatorServer := deviceTwinServer{}
	pb.RegisterDeviceTwinServiceServer(s, &deviceSimulatorServer)
	err = s.Serve(lis)
	utils.FailOnError(err, "failed to serve: %v")
}

func simulation() {
	var tenantID string

	if len(os.Args) < 2 {
		log.Fatal(REQUIRED_PARAMS)
	}
	log.Println(os.Args)

	tenantID = os.Args[1]

	_, err := primitive.ObjectIDFromHex(tenantID)
	utils.FailOnError(err, "Tenant id format is incorrect")

	simulate := service.SimulationConfig{}
	values := config.GetConfigValues()

	simulate.InitSimulation(values, tenantID) // configuration to start simulation

	log.Println("Configuration file read sucessfully >>> ")

	if len(os.Args) == 4 {

		log.Println(">>> Device creation simulation started.. ")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

		model := os.Args[2]
		numOfDevices, err := strconv.ParseInt(os.Args[3], 10, 64)
		utils.FailOnError(err, " num of devices should be integer")

		for k := 0; k < int(numOfDevices); k++ {

			completed, err := simulate.OnboardDeviceOnPremise(strings.ToUpper(model), device.MOBILECOMPUTER)

			utils.FailOnError(err, "Failed when creating simulated divice")
			log.Printf("Device created Succesfully:  %t", completed)
		}
		log.Println(" ")
		log.Println(">>> Device creation simulation completed sucessfully >> ")

	} else if len(os.Args) == 5 {
		serialNumber := os.Args[2]
		eventName := os.Args[3]
		numOfEvents, err := strconv.ParseInt(os.Args[4], 10, 64)
		utils.FailOnError(err, "Number of events should be integer ")

		log.Println("Calling function to send events.. >>")

		for i := 0; i < int(numOfEvents); i++ {
			simulate.SendEvents(serialNumber, eventName)
		}
	} else {
		log.Fatal(REQUIRED_PARAMS)
	}
}
