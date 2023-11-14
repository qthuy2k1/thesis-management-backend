// Original file: src/proto/schedule.proto

import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';
import type { Long } from '@grpc/proto-loader';

export interface AssessItem {
  'id'?: (string);
  'lecturer'?: (_schedule_v1_UserScheduleResponse | null);
  'point'?: (number | string | Long);
  'comment'?: (string);
}

export interface AssessItem__Output {
  'id'?: (string);
  'lecturer'?: (_schedule_v1_UserScheduleResponse__Output);
  'point'?: (Long);
  'comment'?: (string);
}
