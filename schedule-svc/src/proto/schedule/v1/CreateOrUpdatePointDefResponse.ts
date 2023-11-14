// Original file: src/proto/schedule.proto

import type { Point as _schedule_v1_Point, Point__Output as _schedule_v1_Point__Output } from '../../schedule/v1/Point';

export interface CreateOrUpdatePointDefResponse {
  'point'?: (_schedule_v1_Point | null);
  'message'?: (string);
}

export interface CreateOrUpdatePointDefResponse__Output {
  'point'?: (_schedule_v1_Point__Output);
  'message'?: (string);
}
