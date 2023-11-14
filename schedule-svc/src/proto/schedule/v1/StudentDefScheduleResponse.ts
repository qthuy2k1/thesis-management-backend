// Original file: src/proto/schedule.proto

import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';
import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from '../../google/protobuf/Timestamp';

export interface StudentDefScheduleResponse {
  'id'?: (string);
  'infor'?: (_schedule_v1_UserScheduleResponse | null);
  'instructor'?: (_schedule_v1_UserScheduleResponse | null);
  'createdAt'?: (_google_protobuf_Timestamp | null);
}

export interface StudentDefScheduleResponse__Output {
  'id'?: (string);
  'infor'?: (_schedule_v1_UserScheduleResponse__Output);
  'instructor'?: (_schedule_v1_UserScheduleResponse__Output);
  'createdAt'?: (_google_protobuf_Timestamp__Output);
}
