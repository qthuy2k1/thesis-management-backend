// Original file: src/proto/schedule.proto

import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';
import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from '../../google/protobuf/Timestamp';

export interface Notification {
  'id'?: (string);
  'senderUser'?: (_schedule_v1_UserScheduleResponse | null);
  'receiverAuthor'?: (_schedule_v1_UserScheduleResponse | null);
  'type'?: (string);
  'createdAt'?: (_google_protobuf_Timestamp | null);
}

export interface Notification__Output {
  'id'?: (string);
  'senderUser'?: (_schedule_v1_UserScheduleResponse__Output);
  'receiverAuthor'?: (_schedule_v1_UserScheduleResponse__Output);
  'type'?: (string);
  'createdAt'?: (_google_protobuf_Timestamp__Output);
}
