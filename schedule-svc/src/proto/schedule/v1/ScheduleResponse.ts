// Original file: src/proto/schedule.proto

import type { Thesis as _schedule_v1_Thesis, Thesis__Output as _schedule_v1_Thesis__Output } from '../../schedule/v1/Thesis';
import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from '../../google/protobuf/Timestamp';

export interface ScheduleResponse {
  'id'?: (string);
  'thesis'?: (_schedule_v1_Thesis | null);
  'createdAt'?: (_google_protobuf_Timestamp | null);
}

export interface ScheduleResponse__Output {
  'id'?: (string);
  'thesis'?: (_schedule_v1_Thesis__Output);
  'createdAt'?: (_google_protobuf_Timestamp__Output);
}
