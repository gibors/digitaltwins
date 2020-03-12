package main

import (
	"caidc_auto_devicetwins/config"
	device "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/service"
	"caidc_auto_devicetwins/domain/utils"
	pb "caidc_auto_devicetwins/simulator"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

const REQUIRED_PARAMS = "Required parameters Device Createtion simulator: caidc_auto_device_twins {tenantId} {model} {numberOfDevices} /n" +
	"Device events simulation: caidc_auto_device_twins {tenantId} {serialNumber} {queue_event_name} {numberOfEvents}"

const ENV = "local"

type Server struct {
}

var port = flag.Int("port", 50059, "The server port")

func (s *Server) CreateDeviceSimulated(ctx context.Context, request *pb.DevicesCreationRequest) (*pb.DeviceSimulationResponse, error) {

	if request.GetTenantId() == nil || request.GetModel() == nil || request.GetNumberDevices() <= 0 {
		utils.FailOnError(errors.New(" Bad Request: "), REQUIRED_PARAMS)
	}

	tenantID := request.GetTenantId().Value
	model := request.GetModel().Value
	numOfDevices := request.GetNumberDevices()

	simulate := service.SimulationConfig{}
	values := config.GetConfigValues(ENV)

	simulate.InitSimulation(values, tenantID)
	for k := 0; k < int(numOfDevices); k++ {

		completed, err := simulate.OnboardDeviceOnPremise(strings.ToUpper(model), device.MOBILECOMPUTER)
		utils.FailOnError(err, "Failed when creating simulated divice")
		log.Printf("Device created Succesfully:  %t", completed)
	}
	log.Println(" ")
	log.Println(">>> Device creation simulation completed sucessfully >> ")

	succed := &wrappers.BoolValue{Value: true}
	message := &wrappers.StringValue{Value: "Successfully created device(s)"}
	log.Printf("Received tenantId: %v", request.GetTenantId())
	return &pb.DeviceSimulationResponse{SimulationSucced: succed, DetailedMessage: message}, nil
}

func (s *Server) SendDataToDevice(ctx context.Context, request *pb.DevicesDataSimulationRequest) (*pb.DeviceSimulationResponse, error) {

	if request.GetTenantId() == nil || request.GetSerialNumber() == nil || request.GetNumberEvents() <= 0 {
		utils.FailOnError(errors.New(" Bad Request: "), REQUIRED_PARAMS)
	}

	tenantID := request.GetTenantId().Value
	serial := request.GetSerialNumber().Value
	numOfEvents := request.GetNumberEvents()

	var event string
	if request.GetEventType() == nil {
		event = "telemetry"
	} else {
		event = request.GetEventType().Value
	}

	simulate := service.SimulationConfig{}
	values := config.GetConfigValues(ENV)

	simulate.InitSimulation(values, tenantID)

	for i := 0; i < int(numOfEvents); i++ {
		simulate.SendEvents(serial, event)
	}

	succed := &wrappers.BoolValue{Value: true}
	message := &wrappers.StringValue{Value: "Successfully created device(s)"}
	log.Printf("Received tenantId: %v", request.GetTenantId())
	return &pb.DeviceSimulationResponse{SimulationSucced: succed, DetailedMessage: message}, nil
}

func main() {
	if len(os.Args) == 4 { // Onboarding for devices DeviceToRegister (Automation Flow)
		model := os.Args[1]
		numberOfDevices, err := strconv.Atoi(os.Args[2])
		utils.FailOnError(err, "Number of devices should be integer")
		env := os.Args[3]
		simulate := service.SimulationConfig{}
		values := config.GetConfigValues(env)
		simulate.InitSimulation(values, "")
		var deviceType string
		if strings.ToLower(model) == "hcc" {
			deviceType = device.GATEWAY
		} else {
			deviceType = device.MOBILECOMPUTER
		}
		for i := 0; i < numberOfDevices; i++ {
			device := simulate.OnboarDeviceToForge(strings.ToUpper(model), deviceType)
			if device.SystemGUID == "" {
				log.Print("Failed to onboard , please try again ")
			}
			log.Printf("sucessfully created %s", device.SystemGUID)
			log.Println("")
		}

	} else { // Start GRPC Server
		lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		pb.RegisterDeviceTwinServer(s, &Server{})

		log.Println(fmt.Sprintf(" Server started port:%d", *port))
		err = s.Serve(lis)
		utils.FailOnError(err, "failed to serve: %v")
	}
}

func simulateSentienceTenant() {

}
