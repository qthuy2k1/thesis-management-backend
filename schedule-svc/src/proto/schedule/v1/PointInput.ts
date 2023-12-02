// Original file: src/proto/schedule.proto

import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';
import type { AssessItemInput as _schedule_v1_AssessItemInput, AssessItemInput__Output as _schedule_v1_AssessItemInput__Output } from '../../schedule/v1/AssessItemInput';

export interface PointInput {
  'student'?: (_schedule_v1_UserScheduleResponse | null);
  'assesses'?: (_schedule_v1_AssessItemInput)[];
}

export interface PointInput__Output {
  'student'?: (_schedule_v1_UserScheduleResponse__Output);
  'assesses'?: (_schedule_v1_AssessItemInput__Output)[];
}
