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
      Attachment: MessageTypeDefinition
      AttachmentExerciseInput: MessageTypeDefinition
      AttachmentExerciseResponse: MessageTypeDefinition
      AuthorExerciseResponse: MessageTypeDefinition
      ClassroomResponse: MessageTypeDefinition
      CreateAttachmentRequest: MessageTypeDefinition
      CreateAttachmentResponse: MessageTypeDefinition
      CreateExerciseRequest: MessageTypeDefinition
      CreateExerciseResponse: MessageTypeDefinition
      CreateNotificationRequest: MessageTypeDefinition
      CreateNotificationResponse: MessageTypeDefinition
      CreateOrUpdatePointDefRequest: MessageTypeDefinition
      CreateOrUpdatePointDefResponse: MessageTypeDefinition
      CreateScheduleRequest: MessageTypeDefinition
      CreateScheduleResponse: MessageTypeDefinition
      ExerciseInput: MessageTypeDefinition
      ExerciseResponse: MessageTypeDefinition
      GetAllPointDefsRequest: MessageTypeDefinition
      GetAllPointDefsResponse: MessageTypeDefinition
      GetAttachmentRequest: MessageTypeDefinition
      GetAttachmentResponse: MessageTypeDefinition
      GetNotificationsRequest: MessageTypeDefinition
      GetNotificationsResponse: MessageTypeDefinition
      GetSchedulesRequest: MessageTypeDefinition
      GetSchedulesResponse: MessageTypeDefinition
      Notification: MessageTypeDefinition
      Point: MessageTypeDefinition
      ReportingStageExerciseResponse: MessageTypeDefinition
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

