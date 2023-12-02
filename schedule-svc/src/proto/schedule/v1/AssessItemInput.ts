// Original file: src/proto/schedule.proto

import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';

export interface AssessItemInput {
  'lecturer'?: (_schedule_v1_UserScheduleResponse | null);
  'point'?: (number | string);
  'comment'?: (string);
}

export interface AssessItemInput__Output {
  'lecturer'?: (_schedule_v1_UserScheduleResponse__Output);
  'point'?: (number);
  'comment'?: (string);
}
