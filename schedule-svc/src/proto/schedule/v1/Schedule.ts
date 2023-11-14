// Original file: src/proto/schedule.proto

import type { TimeSlots as _schedule_v1_TimeSlots, TimeSlots__Output as _schedule_v1_TimeSlots__Output } from '../../schedule/v1/TimeSlots';
import type { RoomSchedule as _schedule_v1_RoomSchedule, RoomSchedule__Output as _schedule_v1_RoomSchedule__Output } from '../../schedule/v1/RoomSchedule';

export interface Schedule {
  'timeSlots'?: (_schedule_v1_TimeSlots)[];
  'room'?: (_schedule_v1_RoomSchedule | null);
}

export interface Schedule__Output {
  'timeSlots'?: (_schedule_v1_TimeSlots__Output)[];
  'room'?: (_schedule_v1_RoomSchedule__Output);
}
