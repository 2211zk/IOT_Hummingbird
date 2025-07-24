package service

import (
	pb "IOT_Hummingbird_back_end/api/dashboard/v1"
	"IOT_Hummingbird_back_end/internal/biz"
	"context"
)

type DashboardService struct {
	pb.UnimplementedDashboardServer
	uc *biz.DashboardUsecase
}

func NewDashboardService(uc *biz.DashboardUsecase) *DashboardService {
	return &DashboardService{uc: uc}
}

func (s *DashboardService) Overview(ctx context.Context, req *pb.OverviewRequest) (*pb.OverviewReply, error) {
	data, err := s.uc.GetOverview(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.OverviewReply{
		ProductTotal:       int32(data["product_total"]),
		ProductPublished:   int32(data["product_published"]),
		ProductUnpublished: int32(data["product_unpublished"]),
		DeviceTotal:        int32(data["device_total"]),
		DeviceOnline:       int32(data["device_online"]),
		DeviceOffline:      int32(data["device_offline"]),
		DriverTotal:        int32(data["driver_total"]),
		DriverRunning:      int32(data["driver_running"]),
		DriverStopped:      int32(data["driver_stopped"]),
		AlarmTotal:         int32(data["alarm_total"]),
	}, nil
}
func (s *DashboardService) Resource(ctx context.Context, req *pb.ResourceRequest) (*pb.ResourceReply, error) {
	// 这里直接返回模拟数据，实际可用系统包获取
	return &pb.ResourceReply{
		Cpu:    23.5,
		Memory: 58.2,
		Load:   1.2,
		Disk:   70.1,
	}, nil
}
func (s *DashboardService) AlarmStat(ctx context.Context, req *pb.AlarmStatRequest) (*pb.AlarmStatReply, error) {
	db := s.uc.GetDB()
	var alarmTotal, alarmHandled, alarmUnhandled int64
	db.Table("wl_alarm").Count(&alarmTotal)
	db.Table("wl_alarm").Where("alarm_status=1").Count(&alarmHandled)
	db.Table("wl_alarm").Where("alarm_status=0").Count(&alarmUnhandled)
	return &pb.AlarmStatReply{
		AlarmTotal:     int32(alarmTotal),
		AlarmHandled:   int32(alarmHandled),
		AlarmUnhandled: int32(alarmUnhandled),
	}, nil
}
func (s *DashboardService) DeviceMessageCount(ctx context.Context, req *pb.DeviceMessageCountRequest) (*pb.DeviceMessageCountReply, error) {
	db := s.uc.GetDB()
	var count int64
	query := db.Table("wl_device_message")
	if req.StartTime != "" {
		query = query.Where("create_time >= ?", req.StartTime)
	}
	if req.EndTime != "" {
		query = query.Where("create_time <= ?", req.EndTime)
	}
	query.Count(&count)
	return &pb.DeviceMessageCountReply{
		Count: int32(count),
	}, nil
}
