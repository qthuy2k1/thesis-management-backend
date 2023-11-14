// Original file: src/proto/schedule.proto

import type { RoomSchedule as _schedule_v1_RoomSchedule, RoomSchedule__Output as _schedule_v1_RoomSchedule__Output } from '../../schedule/v1/RoomSchedule';
import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';
import type { StudentDefScheduleResponse as _schedule_v1_StudentDefScheduleResponse, StudentDefScheduleResponse__Output as _schedule_v1_StudentDefScheduleResponse__Output } from '../../schedule/v1/StudentDefScheduleResponse';
import type { Long } from '@grpc/proto-loader';

export interface CreateScheduleRequest {
  'startDate'?: (string);
  'quantityWeek'?: (number | string | Long);
  'rooms'?: (_schedule_v1_RoomSchedule)[];
  'councils'?: (_schedule_v1_UserScheduleResponse)[];
  'studentDefs'?: (_schedule_v1_StudentDefScheduleResponse)[];
}

export interface CreateScheduleRequest__Output {
  'startDate'?: (string);
  'quantityWeek'?: (Long);
  'rooms'?: (_schedule_v1_RoomSchedule__Output)[];
  'councils'?: (_schedule_v1_UserScheduleResponse__Output)[];
  'studentDefs'?: (_schedule_v1_StudentDefScheduleResponse__Output)[];
}
