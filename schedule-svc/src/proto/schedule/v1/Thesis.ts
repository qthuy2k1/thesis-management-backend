// Original file: src/proto/schedule.proto

import type { Schedule as _schedule_v1_Schedule, Schedule__Output as _schedule_v1_Schedule__Output } from '../../schedule/v1/Schedule';
import type { UserScheduleResponse as _schedule_v1_UserScheduleResponse, UserScheduleResponse__Output as _schedule_v1_UserScheduleResponse__Output } from '../../schedule/v1/UserScheduleResponse';

export interface Thesis {
  'schedule'?: (_schedule_v1_Schedule | null);
  'council'?: (_schedule_v1_UserScheduleResponse)[];
  'id'?: (string);
}

export interface Thesis__Output {
  'schedule'?: (_schedule_v1_Schedule__Output);
  'council'?: (_schedule_v1_UserScheduleResponse__Output)[];
  'id'?: (string);
}
