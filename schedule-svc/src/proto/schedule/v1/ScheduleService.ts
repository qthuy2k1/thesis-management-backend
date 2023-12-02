// Original file: src/proto/schedule.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { CreateNotificationRequest as _schedule_v1_CreateNotificationRequest, CreateNotificationRequest__Output as _schedule_v1_CreateNotificationRequest__Output } from '../../schedule/v1/CreateNotificationRequest';
import type { CreateNotificationResponse as _schedule_v1_CreateNotificationResponse, CreateNotificationResponse__Output as _schedule_v1_CreateNotificationResponse__Output } from '../../schedule/v1/CreateNotificationResponse';
import type { CreateOrUpdatePointDefRequest as _schedule_v1_CreateOrUpdatePointDefRequest, CreateOrUpdatePointDefRequest__Output as _schedule_v1_CreateOrUpdatePointDefRequest__Output } from '../../schedule/v1/CreateOrUpdatePointDefRequest';
import type { CreateOrUpdatePointDefResponse as _schedule_v1_CreateOrUpdatePointDefResponse, CreateOrUpdatePointDefResponse__Output as _schedule_v1_CreateOrUpdatePointDefResponse__Output } from '../../schedule/v1/CreateOrUpdatePointDefResponse';
import type { CreateScheduleRequest as _schedule_v1_CreateScheduleRequest, CreateScheduleRequest__Output as _schedule_v1_CreateScheduleRequest__Output } from '../../schedule/v1/CreateScheduleRequest';
import type { CreateScheduleResponse as _schedule_v1_CreateScheduleResponse, CreateScheduleResponse__Output as _schedule_v1_CreateScheduleResponse__Output } from '../../schedule/v1/CreateScheduleResponse';
import type { GetAllPointDefsRequest as _schedule_v1_GetAllPointDefsRequest, GetAllPointDefsRequest__Output as _schedule_v1_GetAllPointDefsRequest__Output } from '../../schedule/v1/GetAllPointDefsRequest';
import type { GetAllPointDefsResponse as _schedule_v1_GetAllPointDefsResponse, GetAllPointDefsResponse__Output as _schedule_v1_GetAllPointDefsResponse__Output } from '../../schedule/v1/GetAllPointDefsResponse';
import type { GetNotificationsRequest as _schedule_v1_GetNotificationsRequest, GetNotificationsRequest__Output as _schedule_v1_GetNotificationsRequest__Output } from '../../schedule/v1/GetNotificationsRequest';
import type { GetNotificationsResponse as _schedule_v1_GetNotificationsResponse, GetNotificationsResponse__Output as _schedule_v1_GetNotificationsResponse__Output } from '../../schedule/v1/GetNotificationsResponse';
import type { GetSchedulesRequest as _schedule_v1_GetSchedulesRequest, GetSchedulesRequest__Output as _schedule_v1_GetSchedulesRequest__Output } from '../../schedule/v1/GetSchedulesRequest';
import type { GetSchedulesResponse as _schedule_v1_GetSchedulesResponse, GetSchedulesResponse__Output as _schedule_v1_GetSchedulesResponse__Output } from '../../schedule/v1/GetSchedulesResponse';

export interface ScheduleServiceClient extends grpc.Client {
  CreateNotification(argument: _schedule_v1_CreateNotificationRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  CreateNotification(argument: _schedule_v1_CreateNotificationRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  CreateNotification(argument: _schedule_v1_CreateNotificationRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  CreateNotification(argument: _schedule_v1_CreateNotificationRequest, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  createNotification(argument: _schedule_v1_CreateNotificationRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  createNotification(argument: _schedule_v1_CreateNotificationRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  createNotification(argument: _schedule_v1_CreateNotificationRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  createNotification(argument: _schedule_v1_CreateNotificationRequest, callback: grpc.requestCallback<_schedule_v1_CreateNotificationResponse__Output>): grpc.ClientUnaryCall;
  
  CreateOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  CreateOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  CreateOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  CreateOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  createOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  createOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  createOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  createOrUpdatePointDef(argument: _schedule_v1_CreateOrUpdatePointDefRequest, callback: grpc.requestCallback<_schedule_v1_CreateOrUpdatePointDefResponse__Output>): grpc.ClientUnaryCall;
  
  CreateSchedule(argument: _schedule_v1_CreateScheduleRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  CreateSchedule(argument: _schedule_v1_CreateScheduleRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  CreateSchedule(argument: _schedule_v1_CreateScheduleRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  CreateSchedule(argument: _schedule_v1_CreateScheduleRequest, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  createSchedule(argument: _schedule_v1_CreateScheduleRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  createSchedule(argument: _schedule_v1_CreateScheduleRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  createSchedule(argument: _schedule_v1_CreateScheduleRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  createSchedule(argument: _schedule_v1_CreateScheduleRequest, callback: grpc.requestCallback<_schedule_v1_CreateScheduleResponse__Output>): grpc.ClientUnaryCall;
  
  GetAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  GetAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  GetAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  GetAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  getAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  getAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  getAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  getAllPointDefs(argument: _schedule_v1_GetAllPointDefsRequest, callback: grpc.requestCallback<_schedule_v1_GetAllPointDefsResponse__Output>): grpc.ClientUnaryCall;
  
  GetNotifications(argument: _schedule_v1_GetNotificationsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  GetNotifications(argument: _schedule_v1_GetNotificationsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  GetNotifications(argument: _schedule_v1_GetNotificationsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  GetNotifications(argument: _schedule_v1_GetNotificationsRequest, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  getNotifications(argument: _schedule_v1_GetNotificationsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  getNotifications(argument: _schedule_v1_GetNotificationsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  getNotifications(argument: _schedule_v1_GetNotificationsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  getNotifications(argument: _schedule_v1_GetNotificationsRequest, callback: grpc.requestCallback<_schedule_v1_GetNotificationsResponse__Output>): grpc.ClientUnaryCall;
  
  GetSchedules(argument: _schedule_v1_GetSchedulesRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  GetSchedules(argument: _schedule_v1_GetSchedulesRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  GetSchedules(argument: _schedule_v1_GetSchedulesRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  GetSchedules(argument: _schedule_v1_GetSchedulesRequest, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  getSchedules(argument: _schedule_v1_GetSchedulesRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  getSchedules(argument: _schedule_v1_GetSchedulesRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  getSchedules(argument: _schedule_v1_GetSchedulesRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  getSchedules(argument: _schedule_v1_GetSchedulesRequest, callback: grpc.requestCallback<_schedule_v1_GetSchedulesResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface ScheduleServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateNotification: grpc.handleUnaryCall<_schedule_v1_CreateNotificationRequest__Output, _schedule_v1_CreateNotificationResponse>;
  
  CreateOrUpdatePointDef: grpc.handleUnaryCall<_schedule_v1_CreateOrUpdatePointDefRequest__Output, _schedule_v1_CreateOrUpdatePointDefResponse>;
  
  CreateSchedule: grpc.handleUnaryCall<_schedule_v1_CreateScheduleRequest__Output, _schedule_v1_CreateScheduleResponse>;
  
  GetAllPointDefs: grpc.handleUnaryCall<_schedule_v1_GetAllPointDefsRequest__Output, _schedule_v1_GetAllPointDefsResponse>;
  
  GetNotifications: grpc.handleUnaryCall<_schedule_v1_GetNotificationsRequest__Output, _schedule_v1_GetNotificationsResponse>;
  
  GetSchedules: grpc.handleUnaryCall<_schedule_v1_GetSchedulesRequest__Output, _schedule_v1_GetSchedulesResponse>;
  
}

export interface ScheduleServiceDefinition extends grpc.ServiceDefinition {
  CreateNotification: MethodDefinition<_schedule_v1_CreateNotificationRequest, _schedule_v1_CreateNotificationResponse, _schedule_v1_CreateNotificationRequest__Output, _schedule_v1_CreateNotificationResponse__Output>
  CreateOrUpdatePointDef: MethodDefinition<_schedule_v1_CreateOrUpdatePointDefRequest, _schedule_v1_CreateOrUpdatePointDefResponse, _schedule_v1_CreateOrUpdatePointDefRequest__Output, _schedule_v1_CreateOrUpdatePointDefResponse__Output>
  CreateSchedule: MethodDefinition<_schedule_v1_CreateScheduleRequest, _schedule_v1_CreateScheduleResponse, _schedule_v1_CreateScheduleRequest__Output, _schedule_v1_CreateScheduleResponse__Output>
  GetAllPointDefs: MethodDefinition<_schedule_v1_GetAllPointDefsRequest, _schedule_v1_GetAllPointDefsResponse, _schedule_v1_GetAllPointDefsRequest__Output, _schedule_v1_GetAllPointDefsResponse__Output>
  GetNotifications: MethodDefinition<_schedule_v1_GetNotificationsRequest, _schedule_v1_GetNotificationsResponse, _schedule_v1_GetNotificationsRequest__Output, _schedule_v1_GetNotificationsResponse__Output>
  GetSchedules: MethodDefinition<_schedule_v1_GetSchedulesRequest, _schedule_v1_GetSchedulesResponse, _schedule_v1_GetSchedulesRequest__Output, _schedule_v1_GetSchedulesResponse__Output>
}
