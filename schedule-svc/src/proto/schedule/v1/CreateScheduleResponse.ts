// Original file: src/proto/schedule.proto

import type { Thesis as _schedule_v1_Thesis, Thesis__Output as _schedule_v1_Thesis__Output } from '../../schedule/v1/Thesis';

export interface CreateScheduleResponse {
  'id'?: (string);
  'thesis'?: (_schedule_v1_Thesis)[];
}

export interface CreateScheduleResponse__Output {
  'id'?: (string);
  'thesis'?: (_schedule_v1_Thesis__Output)[];
}
