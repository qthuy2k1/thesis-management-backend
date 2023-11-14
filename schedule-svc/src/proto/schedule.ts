import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { ScheduleServiceClient as _schedule_v1_ScheduleServiceClient, ScheduleServiceDefinition as _schedule_v1_ScheduleServiceDefinition } from './schedule/v1/ScheduleService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  google: {
    protobuf: {
      Timestamp: MessageTypeDefinition
    }
  }
  schedule: {
    v1: {
      AssessItem: MessageTypeDefinition
      CreateNotificationRequest: MessageTypeDefinition
      CreateNotificationResponse: MessageTypeDefinition
      CreateOrUpdatePointDefRequest: MessageTypeDefinition
      CreateOrUpdatePointDefResponse: MessageTypeDefinition
      CreateScheduleRequest: MessageTypeDefinition
      CreateScheduleResponse: MessageTypeDefinition
      GetAllPointDefRequest: MessageTypeDefinition
      GetAllPointDefResponse: MessageTypeDefinition
      GetNotificationsRequest: MessageTypeDefinition
      GetNotificationsResponse: MessageTypeDefinition
      GetSchedulesRequest: MessageTypeDefinition
      GetSchedulesResponse: MessageTypeDefinition
      Notification: MessageTypeDefinition
      Point: MessageTypeDefinition
      RoomSchedule: MessageTypeDefinition
      Schedule: MessageTypeDefinition
      ScheduleResponse: MessageTypeDefinition
      ScheduleService: SubtypeConstructor<typeof grpc.Client, _schedule_v1_ScheduleServiceClient> & { service: _schedule_v1_ScheduleServiceDefinition }
      StudentDefScheduleResponse: MessageTypeDefinition
      Thesis: MessageTypeDefinition
      TimeSlot: MessageTypeDefinition
      TimeSlots: MessageTypeDefinition
      UserScheduleResponse: MessageTypeDefinition
    }
  }
}

