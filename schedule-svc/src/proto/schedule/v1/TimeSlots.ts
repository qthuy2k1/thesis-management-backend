// Original file: src/proto/schedule.proto

import type { StudentDefScheduleResponse as _schedule_v1_StudentDefScheduleResponse, StudentDefScheduleResponse__Output as _schedule_v1_StudentDefScheduleResponse__Output } from '../../schedule/v1/StudentDefScheduleResponse';
import type { TimeSlot as _schedule_v1_TimeSlot, TimeSlot__Output as _schedule_v1_TimeSlot__Output } from '../../schedule/v1/TimeSlot';

export interface TimeSlots {
  'student'?: (_schedule_v1_StudentDefScheduleResponse | null);
  'timeSlot'?: (_schedule_v1_TimeSlot | null);
}

export interface TimeSlots__Output {
  'student'?: (_schedule_v1_StudentDefScheduleResponse__Output);
  'timeSlot'?: (_schedule_v1_TimeSlot__Output);
}
