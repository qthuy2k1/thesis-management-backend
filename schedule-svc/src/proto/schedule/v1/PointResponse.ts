// Original file: src/proto/schedule.proto

import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';
import type { AssessItem as _schedule_v1_AssessItem, AssessItem__Output as _schedule_v1_AssessItem__Output } from '../../schedule/v1/AssessItem';

export interface PointResponse {
  'id'?: (string);
  'student'?: (_schedule_v1_UserScheduleResponse | null);
  'assesses'?: (_schedule_v1_AssessItem)[];
}

export interface PointResponse__Output {
  'id'?: (string);
  'student'?: (_schedule_v1_UserScheduleResponse__Output);
  'assesses'?: (_schedule_v1_AssessItem__Output)[];
}
