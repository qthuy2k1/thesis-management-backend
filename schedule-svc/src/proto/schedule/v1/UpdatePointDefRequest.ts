// Original file: src/proto/schedule.proto

import type { Point as _schedule_v1_Point, Point__Output as _schedule_v1_Point__Output } from '../../schedule/v1/Point';

export interface UpdatePointDefRequest {
  'id'?: (string);
  'point'?: (_schedule_v1_Point | null);
}

export interface UpdatePointDefRequest__Output {
  'id'?: (string);
  'point'?: (_schedule_v1_Point__Output);
}
